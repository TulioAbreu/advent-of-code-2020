import java.io.File
import java.util.Collections.max
import java.util.Collections.min

fun readInputFile(filepath: String): List<Long> {
    return File(filepath).readLines().map { line: String -> line.toLong()};
}

fun checkRule(targetValue: Long, values: List<Long>): Boolean {
    for ((index1, val1) in values.withIndex()) {
        for ((index2, val2) in values.withIndex()) {
            if (index1 == index2) {
                continue;
            }
            if (val1 + val2 == targetValue) {
                return true;
            }
        }
    }
    return false;
}

fun partOne(encryptedCode: List<Long>): Long {
    var result: Long = 0;
    val PREAMBLE_SIZE: Int = 25;
    for (i in PREAMBLE_SIZE until encryptedCode.size) {
        val value = encryptedCode[i];
        var lastValues = encryptedCode.slice(i-PREAMBLE_SIZE until i);
        if (!checkRule(value, lastValues)) {
            result = value;
            break;
        }
    }
    return result;
}

fun partTwo(encryptedCode: List<Long>, targetValue: Long): Long {
    for (startIndex in encryptedCode.indices) {
        for (endIndex in encryptedCode.indices) {
            if (startIndex == endIndex) {
                continue;
            }
            var codeSlice = encryptedCode.slice(startIndex..endIndex);
            if (targetValue == codeSlice.sum()) {
                return min(codeSlice) + max(codeSlice);
            }
        }
    }
    return 0;
}

fun main() {
    var encryptedCode: List<Long> = readInputFile("./input.prod.txt");
    val targetNumber = partOne(encryptedCode);
    println("Part 1 ] Response is $targetNumber");
    val encryptionWeakness = partTwo(encryptedCode, targetNumber);
    println("Part 2 ] Response is $encryptionWeakness");
}