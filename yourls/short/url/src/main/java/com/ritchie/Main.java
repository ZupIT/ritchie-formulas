package com.ritchie;

import com.ritchie.formula.Formula;

public class Main {

  public static void main(String[] args) throws Exception {
    String yourlsApiEndpoint = System.getenv("YOURLS_API_ENDPOINT");
    String yourlsApiSecret = System.getenv("YOURLS_API_SECRET");
    String urlToShorten = System.getenv("URL_TO_SHORTEN");
    String keyword = System.getenv("KEYWORD");
    Formula formula = new Formula(yourlsApiEndpoint, yourlsApiSecret, urlToShorten, keyword);
    formula.Run();
    System.exit(0);
  }
}
