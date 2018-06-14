package linked_list;

public class LinkedListTest {

  public static void main(String[] args) {
    LinkedList ll = new LinkedList();

    ll.append(1);
    ll.show();

    ll.prepend(9);
    ll.show();

    ll.append(2);
    ll.append(3);
    ll.append(4);
    ll.append(5);
    ll.show();

    System.out.println(ll.pop());
    System.out.println("------------------");

    ll.show();

    System.out.println(ll.shift());
    System.out.println("------------------");

    ll.show();

    System.out.println(ll.count());
    System.out.println("------------------");

    ll.insert(2, 10);
    ll.show();
    ll.show();

  }

}
