import coffee.Coffee;

public class Main {

    public static void main(String[] args) throws Exception {
        String name = System.getenv("NAME");
        String type = System.getenv("COFFEE_TYPE");
        boolean delivery = Boolean.parseBoolean(System.getenv("DELIVERY"));
        Coffee coffee = new Coffee(name, type, delivery);
        coffee.Prepare();
    }
}
