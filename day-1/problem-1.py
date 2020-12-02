from time import time

def read_input():
    with open("./input.txt", "r") as file:
        file_lines = file.readlines()
        numbers = [int(line) for line in file_lines]
    return numbers

def main():
    numbers = read_input()
    expected_values = []
    TARGET_VALUE = 2020
    for number in numbers:
        expected = TARGET_VALUE - number
        if number in expected_values:
            print("%d * %d = %d" % (number, expected, number * expected))
            return
        expected_values.append(expected)

if __name__ == "__main__":
    main()

