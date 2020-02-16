package github.com.kondroid00.book_tech8_golang.book.code.example03;

class A {
  public int i;

  public A(int i){
      this.i = i;
  }
}

public class Main {
  public static void main(String[] args){
      A a = new A(1);
      A b = a;   //‥‥‥③
      b.i = 100; //‥‥‥④
      System.out.printf("a.i: %d\n", a.i);
      System.out.printf("b.i: %d\n", b.i);
  }
}