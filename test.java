package test;

import java.util.Scanner;

public static void main(String[] args) {
    Scanner in = new Scanner(System.in);
    String input = in.nextLine();
    input = "test";
    System.out.println(input);
    if (5 > 6) {
        System.out.println("This isn't printed");
    }

    int i = 0;
    do {
        System.out.println(i);
        i++;
    } while (i < 3);
    for (i = 0; i < 3; i++) {
        System.out.println(i);
    }
    int[] arrayOfInts = {1, 2, 3};
    for (int i : arrayOfInts) {
        System.out.println(i);
    }
}
