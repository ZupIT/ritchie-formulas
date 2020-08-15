import shorturl.Shorturl;

public class TestMain {

    public static void main(String[] args) throws Exception {
        Shorturl shorturl = new Shorturl("http://uol.com.br", "uol");
        shorturl.Run();
    }
}