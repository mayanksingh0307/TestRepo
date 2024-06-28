import java.util.ArrayList;
import java.util.List;
import java.util.Random;

public class PerformanceIssues {
    public static void main(String[] args) {
        String result = "";
        for (int i = 0; i < 10000; i++) {
            result += "a";
        }
        System.out.println(result);

        List<Integer> list = new ArrayList<>();
        for (int i = 0; i < 10000; i++) {
            list.add(i);
        }

        for (int i = 0; i < 10000; i++) {
            Random random = new Random();
            int num = random.nextInt();
            System.out.println(num);
        }

        int count = 0;
        for (int i = 2; i < 10000; i++) {
            if (isPrime(i)) {
                count++;
            }
        }
        System.out.println("Number of primes: " + count);

        try (FileReader fileReader = new FileReader("largefile.txt")) {
            int ch;
            while ((ch = fileReader.read()) != -1) {
                System.out.print((char) ch);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private static boolean isPrime(int n) {
        if (n <= 1) return false;
        for (int i = 2; i < n; i++) {
            if (n % i == 0) return false;
        }
        return true;
    }
}
