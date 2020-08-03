package shorturl;

import org.json.simple.JSONObject;
import org.json.simple.parser.JSONParser;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.HashMap;
import java.util.Map;
import java.util.Properties;
import java.util.StringJoiner;

public class Shorturl {

    private final String API_ENDPOINT = "yourls_api_endpoint";
    private final String API_SECRET = "yourls_api_secret";

    private String urlToShorten;
    private String keyword;

    public void Run() throws Exception {

        //Load API config
        Properties prop = null;
        try {
            prop = readYOURLSProperties();
            if (!prop.containsKey(API_ENDPOINT) || !prop.containsKey(API_SECRET)) {
                System.out.println("Invalid configuration file. Please run again: \"rit yourls init\"");
                System.exit(-1);
            }
        } catch (FileNotFoundException fileNotFoundException) {
            System.out.println("Configuration file not found. Please run \"rit yourls init\"");
            System.exit(-1);
        } catch (IOException ioException) {
            System.out.println("Error loading configuration file. Please run again: \"rit yourls init\"");
            System.exit(-1);
        }

        HttpClient httpClient = HttpClient.newBuilder()
                .version(HttpClient.Version.HTTP_2)
                .build();

        HashMap<String, String> params = createRequestParams(prop.getProperty(API_SECRET));
        HttpRequest request = createApiRequest(prop.getProperty(API_ENDPOINT), params);

        if(request == null) {
            System.out.println("Error generating HTTP request.");
            System.exit(-1);
        }
        HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());

        if (response.statusCode() == 200) {
            JSONObject jsonObject = (JSONObject) new JSONParser().parse(response.body());
            System.out.println("Success! " + urlToShorten + " -> " + jsonObject.get("shorturl"));
            System.exit(0);
        } else {
            System.out.println("Something went wrong: Status code " + response.statusCode());
            System.out.println(response.body());
            System.exit(0);
        }
    }

    private HttpRequest createApiRequest(String apiEndpoint, HashMap<String, String> params) {
        try {
            HttpRequest request = HttpRequest.newBuilder()
                    .POST(HttpRequest.BodyPublishers.ofString(createPostBody(params)))
                    .uri(new URI(apiEndpoint))
                    .setHeader("Content-Type", "application/x-www-form-urlencoded")
                    .build();
            return request;
        } catch (Exception e) {
            return null;
        }
    }

    private HashMap<String, String> createRequestParams(String apiSecret) {
        HashMap<String, String> params = new HashMap<>();
        params.put("signature", apiSecret);
        params.put("format", "json");
        params.put("action", "shorturl");
        params.put("url", this.urlToShorten);
        if (this.keyword != null && !(this.keyword.isEmpty() || this.keyword.isEmpty())) {
            params.put("keyword", this.keyword);
        }

        return params;
    }

    private String createPostBody(HashMap<String, String> params) {
        StringJoiner sj = new StringJoiner("&");
        for (Map.Entry<String, String> param : params.entrySet()) {
            sj.add(param.getKey() + "=" + param.getValue());
        }

        return sj.toString();
    }

    private Properties readYOURLSProperties() throws IOException {
        String home = System.getProperty("user.home");
        FileInputStream fis = null;

        try {
            fis = new FileInputStream(home + "/.yourls/yourls.properties");
            Properties properties = new Properties();
            properties.load(fis);
            return properties;
        } catch (IOException ioException) {
            throw ioException;
        } finally {
            if (fis != null) {
                fis.close();
            }
        }
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