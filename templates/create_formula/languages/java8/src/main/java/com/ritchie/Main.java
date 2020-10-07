package com.ritchie;

import com.ritchie.formula.Formula;

public class Main {

  public static void main(String[] args) {

    String inputText = System.getenv("INPUT_TEXT");
    boolean inputBoolean = Boolean.parseBoolean(System.getenv("INPUT_BOOLEAN"));
    String inputList = System.getenv("INPUT_LIST");
    String inputPassword = System.getenv("INPUT_PASSWORD");

    Formula formula = new Formula(inputText, inputBoolean, inputList, inputPassword);
    formula.Run();
  }
}
