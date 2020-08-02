package shorturl;

import org.json.simple.JSONObject;
import org.json.simple.parser.JSONParser;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.HashMap;
import java.util.Map;
import java.util.StringJoiner;

public class Shorturl {

    private static final String YOURLS_SECRET = "bef5385e47";

    private String urlToShorten;
    private String keyword;

    public void Run() throws Exception {

        HttpClient httpClient = HttpClient.newBuilder()
                .version(HttpClient.Version.HTTP_2)
                .build();

        long timestamp = System.currentTimeMillis();
        HashMap<String, String> params = new HashMap<>();
        params.put("signature", YOURLS_SECRET);
        params.put("format", "json");
        params.put("action", "shorturl");
        params.put("url", this.urlToShorten);
        if (this.keyword != null && !(this.keyword.isEmpty() || this.keyword.isEmpty())) {
            params.put("keyword", this.keyword);
        }
        HttpRequest request = HttpRequest.newBuilder()
                .POST(HttpRequest.BodyPublishers.ofString(createPostBody(params)))
                .uri(new URI("http://netomar.in/yourls-api.php"))
                .setHeader("Content-Type", "application/x-www-form-urlencoded")
                .build();

        HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());

        if (response.statusCode() == 200) {
            JSONObject jsonObject = (JSONObject) new JSONParser().parse(response.body());
            System.out.println("Success! " + urlToShorten + " -> " + jsonObject.get("shorturl"));
        } else {
            System.out.println("Something went wrong: Status code " + response.statusCode());
            System.out.println(response.body());
        }
    }

    private String createPostBody(HashMap<String, String> params) {
        StringJoiner sj = new StringJoiner("&");
        for (Map.Entry<String, String> param : params.entrySet()) {
            sj.add(param.getKey() + "=" + param.getValue());
        }

        return sj.toString();
    }

    public Shorturl(String urlToShorten, String keyword) {
        this.urlToShorten = urlToShorten;
        this.keyword = keyword;
    }

    public String getUrlToShorten() {
        return urlToShorten;
    }

    public void setUrlToShorten(String urlToShorten) {
        this.urlToShorten = urlToShorten;
    }
}