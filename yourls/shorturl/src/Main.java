import shorturl.Shorturl;

public class Main {

    public static void main(String[] args) throws Exception {
        String input1 = System.getenv("URL_TO_SHORTEN");
        String input2 = System.getenv("KEYWORD");
        Shorturl shorturl = new Shorturl(input1, input2);
        shorturl.Run();
    }
}