package com.ritchie;

import com.ritchie.formula.Formula;

public class Main {
  public static void main(String[] args) {
    String input_text = System.getenv("INPUT_TEXT");
    String input_list = System.getenv("INPUT_LIST");
    String password = System.getenv("PASSWORD");
    boolean input_boolean = Boolean.parseBoolean(System.getenv("INPUT_BOOLEAN"));
    Formula formula = new Formula(input_text, input_list, password, input_boolean);
    formula.Run();
  }
}
