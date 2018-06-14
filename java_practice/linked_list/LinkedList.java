package linked_list;

public class LinkedList {

  public Node head;
  public int count;

  public LinkedList() {
    count = 0;
  }

  public void append(int data) {
    if (head == null) {
      head = new Node(data);
    } else {
      Node currentNode = head;
      while (currentNode.next != null) {
        currentNode = currentNode.next;
      }
      currentNode.next = new Node(data);
    }
    count++;
  }

  public void prepend(int data) {
    Node placeholder = head;
    head = new Node(data);
    head.next = placeholder;
    count++;
  }

  public void insert(int index, int data) {
    Node currentNode = head;
    for (int i = 1; i <= index; i++) {
      currentNode = currentNode.next;
    }
    Node placeholder = currentNode.next;
    currentNode.next = new Node(data);
    currentNode.next.setNext(placeholder);
    count++;
  }

  public int pop() {
    Node currentNode = head;
    if (head.next != null) {
      while (currentNode.next.next != null) {
        currentNode = currentNode.next;
      }
      int rn = currentNode.next.data;
      currentNode.next = null;
      count--;
      return rn;
    } else {
      int returnedNode = head.next.data;
      head.next = null;
      count--;
      return returnedNode;
    }
  }

  public int shift() {
    int returnedData = head.data;
    head = head.next;
    count--;
    return returnedData;
  }

  public int count() {
    return count;
  }

  public void show() {
    Node currentNode = head;
    if (head != null) {
      for (int i = 0; i < count; i++) {
        System.out.println(currentNode.data);
        currentNode = currentNode.next;
      }
    } else {
      System.out.println("Linked List Empty");
    }
    System.out.println("------------------");
  }

}
