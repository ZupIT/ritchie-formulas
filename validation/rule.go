package validation

type Rule struct {
	Language string
	Extensions string
	Files    []string
	Folders  []string
}

func createRules() map[string]Rule {
	return map[string]Rule{
		"go" : {
			Language:   "Golang",
			Extensions: "go",
			Files:      []string{"go.mod", "main.go", "README.md", "config.json", "help.txt"},
			Folders:    []string{"src", "pkg"},
		},
		"java" : {
			Language: "Java",
			Extensions: "java",
			Files:      []string{"main.java", "README.md", "config.json", "help.txt"},
			Folders:    []string{"src", "pkg"},
		},
		"sh" : {
			Language: "shell script",
			Extensions: "sh",
			Files:      []string{"main.sh", "README.md", "config.json", "Makefile", "help.txt"},
			Folders:    []string{"src", "pkg"},
		}}
}