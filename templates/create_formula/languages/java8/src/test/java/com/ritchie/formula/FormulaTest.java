package com.ritchie.formula;

import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import java.io.ByteArrayOutputStream;
import java.io.PrintStream;

import static org.junit.Assert.assertEquals;

public class FormulaTest {

  private final ByteArrayOutputStream outContent = new ByteArrayOutputStream();
  private final PrintStream originalOut = System.out;

  @Before
  public void setUpStreams() {
    System.setOut(new PrintStream(outContent));
  }

  @After
  public void restoreStreams() {
    System.setOut(originalOut);
  }

  @Test
  public void runTrueInput() {
    new Formula("Hello", true, "world", "pass").Run();

    assertEquals("Hello World!\n" +
            "\u001B[32mMy name is Hello.\u001B[39m\n" +
            "\u001B[34mI’ve already created formulas using Ritchie.\u001B[39m\n" +
            "\u001B[33mToday, I want to automate world.\u001B[39m\n" +
            "\u001B[36mMy secret is pass.\u001B[39m\n", outContent.toString());
  }

  @Test
  public void runFalseInput() {
    new Formula("Hello", false, "world", "pass").Run();

    assertEquals("Hello World!\n" +
            "\u001B[32mMy name is Hello.\u001B[39m\n" +
            "\u001B[31mI’m excited in creating new formulas using Ritchie.\u001B[39m\n" +
            "\u001B[33mToday, I want to automate world.\u001B[39m\n" +
            "\u001B[36mMy secret is pass.\u001B[39m\n", outContent.toString());
  }

  @Test
  public void runNoSecresInput() {
    new Formula("Hello", false, "world", "").Run();

    assertEquals("Hello World!\n" +
            "\u001B[32mMy name is Hello.\u001B[39m\n" +
            "\u001B[31mI’m excited in creating new formulas using Ritchie.\u001B[39m\n" +
            "\u001B[33mToday, I want to automate world.\u001B[39m\n" +
            "\u001B[36mMy secret is .\u001B[39m\n", outContent.toString());
  }
}
