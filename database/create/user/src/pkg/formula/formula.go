// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"fmt"
	"io"
	"net/url"
	"strconv"

	"database/sql"

	"github.com/gookit/color"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Formula struct {
	Login        string
	Credential   string
	DatabaseType string
	Hostname     string
	Port         string
	Username     string
	Password     string
}

func (f Formula) Run(writer io.Writer) {

	sqlStatement, err := CreateStatement(f)
	if err != nil {
		fmt.Fprintf(writer, color.FgRed.Render(fmt.Sprintln(err.Error())))
		return
	}

	switch f.DatabaseType {
	case "mysql":
		db := MySQLConn(f)
		defer db.Close()

		_, err = db.Exec(sqlStatement)

		if err != nil {
			fmt.Fprintf(writer, color.FgRed.Render(fmt.Sprintln(err.Error())))
			return
		}
	case "pgsql":
		db := PgSQLConn(f)
		defer db.Close()

		_, err = db.Exec(sqlStatement)

		if err != nil {
			fmt.Fprintf(writer, color.FgRed.Render(fmt.Sprintln(err.Error())))
			return
		}
	case "mssql":
		db := MssqlConn(f)
		defer db.Close()

		_, err = db.Exec(sqlStatement)

		if err != nil {
			fmt.Fprintf(writer, color.FgRed.Render(fmt.Sprintln(err.Error())))
			return
		}
	}

	fmt.Fprintf(writer, color.FgGreen.Render(fmt.Sprintln("User "+f.Username+" created with success")))

}

func MySQLConn(f Formula) *sql.DB {
	dbDriver := "mysql"
	host := f.Hostname
	dbUser := f.Login
	dbPass := f.Credential
	dbName := "mysql"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+host+":"+f.Port+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func PgSQLConn(f Formula) *sql.DB {
	port, _ := strconv.Atoi(f.Port)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		f.Hostname, port, f.Login, f.Credential, "postgres")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func MssqlConn(f Formula) *sql.DB {
	query := url.Values{}
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(f.Login, f.Credential),
		Host:   fmt.Sprintf("%s:%s", f.Hostname, f.Port),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CreateStatement(f Formula) (string, error) {

	switch f.DatabaseType {
	case "mysql":
		return fmt.Sprintf(`create user '%s' identified by '%s'`, f.Username, f.Password), nil
	case "pgsql":
		return fmt.Sprintf(`create user "%s" with login password '%s'`, f.Username, f.Password), nil
	case "mssql":
		return fmt.Sprintf(`create login [%s] with password='%s'`, f.Username, f.Password), nil
	}

	return "", fmt.Errorf("Erro: wrong database type")
}
