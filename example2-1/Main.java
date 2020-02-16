package github.com.kondroid00.book_tech8_golang.book.code.example02;

public class Main {
    public static void main(String[] args){
        int a = 1;
        int b = a; //‥‥‥①
        b = 100;   //‥‥‥②
        System.out.printf("a: %d\n", a);
        System.out.printf("b: %d\n", b);
    }
}