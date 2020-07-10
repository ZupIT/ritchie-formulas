package coffee;

public class Coffee {

  private String owner;
  private String type;
  private boolean delivery;

  public void Prepare() throws Exception {
    System.out.printf("Preparing your coffee %s .....\n", owner);
    Thread.sleep(1000);
    System.out.println("......");
    Thread.sleep(1000);
    System.out.println("......");
    Thread.sleep(1000);
    System.out.println("......");
    Thread.sleep(1000);
    if (delivery) {
      System.out.printf("Your %s coffee is ready, enjoy your trip\n", type);
    } else {
      System.out.printf("Your %s coffee is ready, have a seat and enjoy your drink\n", type);
    }
  }

  public Coffee(String owner, String type, boolean delivery) {
    this.owner = owner;
    this.type = type;
    this.delivery = delivery;
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
}
