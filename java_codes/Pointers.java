package java_codes;

public class Pointers {

    void changeNum(int num){
        num = 5;
        System.out.println(num);
    }
    public static void main(String[] args) {
        int num = 1;
        Pointers p = new Pointers();
        p.changeNum(num);
        System.out.println(num);
    }
}
