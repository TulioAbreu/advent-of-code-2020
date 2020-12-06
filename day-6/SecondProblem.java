import java.io.File;
import java.io.FileNotFoundException;
import java.security.PermissionCollection;
import java.util.LinkedList;
import java.util.Scanner;
import java.util.List;

public class SecondProblem {
    public static List<String> readFile(String filepath) {
        List<String> result = new LinkedList<String>();
        try {
            File file = new File(filepath);
            Scanner scanner = new Scanner(file);
            while (scanner.hasNextLine()) {
                String line = scanner.nextLine();
                result.add(line);
            }
            scanner.close();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }
        return result;
    }

    public static int[] getEmptyGroupStatus() {
        int[] groupStatus = new int[26];
        for (int i = 0; i < 26; ++i) {
            groupStatus[i] = 0;
        }
        return groupStatus;
    }

    public static int getResponsesCount(int[] groupStatus, int personCount) {
        int responsesCount = 0;
        for (int i = 0; i < groupStatus.length; ++i) {
            if (groupStatus[i] == personCount) {
                responsesCount += 1;
            }
        }
        return responsesCount;
    }

    public static void main(String[] args) {
        List<String> lines = readFile("./input.txt");
        int[] groupStatus = getEmptyGroupStatus();
        int groupPersonCount = 0;
        int sum = 0;
        for (String rawLine : lines) {
            String line = rawLine.trim();
            if (line.length() == 0) {
                sum += getResponsesCount(groupStatus, groupPersonCount);
                groupPersonCount = 0;
                groupStatus = getEmptyGroupStatus();
                continue;
            }
            groupPersonCount++;
            for (int i = 0; i < line.length(); ++i) {
                char currentChar = line.charAt(i);
                int index = currentChar - 'a';
                groupStatus[index]++;
            }
        }
        sum += getResponsesCount(groupStatus, groupPersonCount);
        System.out.println("Sum = " + String.valueOf(sum));
    }
}
