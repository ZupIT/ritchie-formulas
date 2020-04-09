package coffee;

public class Coffee {

    private String owner;
    private String type;
    private boolean delivery;
    private boolean noDelay;

    public void Prepare() throws Exception {
        System.out.printf("Preparing your coffee %s .....\n", owner);
        System.out.println("......");
        delay(noDelay);
        System.out.println("......");
        delay(noDelay);
        System.out.println("......");
        delay(noDelay);
        if (delivery) {
            System.out.printf("Your %s coffee is ready, enjoy your trip\n", type);
        } else {
            System.out.printf("Your %s coffee is ready, have a seat and enjoy your drink\n", type);
        }
    }

    public void delay(boolean noDelay) {
        if (!noDelay)
           Thread.sleep(1000);
    }

    public Coffee(String owner, String type, boolean delivery, boolean noDelay) {
        this.owner = owner;
        this.type = type;
        this.delivery = delivery;
        this.noDelay = noDelay == null ? false : true;
    }

    public String getOwner() {
        return owner;
    }

    public void setOwner(String owner) {
        this.owner = owner;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public boolean isDelivery() {
        return delivery;
    }

    public void setDelivery(boolean delivery) {
        this.delivery = delivery;
    }

    public boolean noDelay() {
            return noDelay;
        }

    public void setDelay(boolean noDelay) {
        this.noDelay = noDelay;
    }

}
