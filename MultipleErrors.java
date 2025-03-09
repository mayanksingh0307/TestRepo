import java.io.*;
import java.util.*;

public class MajorIssues {
    public static void main(String[] args) {
        System.out.println("Hello, World")

        int result = a + 5;

        printMessage("This will cause an error");

        String username = "admin";
        String password = "password123";

        String longString = "";
        for (int i = 0; i < 10000; i++) {
            longString += "a";
        }

        for (int i = 0; i < 10000; i++) {
            Random random = new Random();
            int num = random.nextInt();
            System.out.println(num);
        }

        if (result > 42) {
            System.out.println("Result is greater than the magic number.");
        }

        try (FileReader fileReader = new FileReader("largefile.txt")) {
            int ch;
            while ((ch = fileReader.read()) != -1) {
                System.out.print((char) ch);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }

        for (int i = 0; i < 10; i++) {
            for (int j = 0; j < 10; j++) {
                for (int k = 0; k < 10; k++) {
                    System.out.println("Nested loop level 3");
                }
            }
        }

        int insecureRandom = (int) (Math.random() * 100);
        System.out.println("Insecure random number: " + insecureRandom);
    }

    // public static void printMessage(String message) {
    //     System.out.println(message);
    // }
}
