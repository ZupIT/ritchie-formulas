package coffee;

import org.junit.Test;

public class CoffeeTest {

    @Test
    public void prepareEspresso() throws Exception {
        Coffee coffee = new Coffee(
            "Dennis Ritchie",
            "espresso",
            true,
            true
        );

        coffee.Prepare();
    }

    @Test
    public void prepareCappuccino() throws Exception {
        Coffee coffee = new Coffee(
                "Dennis Ritchie",
                "cappuccino",
                true,
                true
        );

        coffee.Prepare();
    }

    @Test
    public void prepareMacchiato() throws Exception {
        Coffee coffee = new Coffee(
                "Dennis Ritchie",
                "macchiato",
                true,
                true
        );

        coffee.Prepare();
    }

    @Test
    public void prepareLatte() throws Exception {
        Coffee coffee = new Coffee(
                "Dennis Ritchie",
                "latte",
                true,
                true
        );

        coffee.Prepare();
    }
}