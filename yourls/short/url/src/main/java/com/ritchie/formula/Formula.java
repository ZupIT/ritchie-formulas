package com.ritchie.formula;

import org.json.simple.JSONObject;
import org.json.simple.parser.JSONParser;
import org.json.simple.parser.ParseException;

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.HashMap;
import java.util.Map;
import java.util.StringJoiner;


public class Formula {

    private final String yourlsApiEndpoint;
    private final String yourlsApiSecret;
    private final String urlToShorten;
    private final String keyword;

    public Formula(String yourlsApiEndpoint, String yourlsApiSecret, String urlToShorten, String keyword) {
        this.yourlsApiEndpoint = yourlsApiEndpoint;
        this.yourlsApiSecret = yourlsApiSecret;
        this.urlToShorten = urlToShorten.trim();
        this.keyword = keyword.trim();
    }

    public String Run() throws IOException, InterruptedException, ParseException {

        if (this.yourlsApiEndpoint == null || this.yourlsApiEndpoint == null
                || this.yourlsApiEndpoint.isEmpty() || this.yourlsApiSecret.isEmpty()) {
            System.out.println("Invalid credentials. Please run: \"rit set credential\"");
            System.exit(-1);
        }

        HttpClient httpClient = HttpClient.newBuilder()
                .version(HttpClient.Version.HTTP_2)
                .build();

        HashMap<String, String> params = createRequestParams(this.yourlsApiSecret);
        HttpRequest request = createApiRequest(this.yourlsApiEndpoint, params);

        if (request == null) {
            System.out.println("Error generating HTTP request.");
            System.exit(-1);
        }

        HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());

        if (response.statusCode() == 200) {
            JSONObject jsonObject = (JSONObject) new JSONParser().parse(response.body());
            String msg = "Success! " + urlToShorten + " -> " + jsonObject.get("shorturl");
            System.out.println(msg);
            return msg;
        } else {
            String msgError = "Something went wrong: Status code " + response.statusCode();
            System.out.println();
            System.out.println(response.body());
            return msgError;
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
        if (this.keyword != null && !this.keyword.isEmpty()) {
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
}
