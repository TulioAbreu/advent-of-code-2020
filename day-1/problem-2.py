from itertools import combinations

def read_input():
    with open("./input.txt", "r") as file:
        file_lines = file.readlines()
        numbers = [int(line) for line in file_lines]
    return numbers

def main():
    numbers = read_input()
    all_combinations = [(a,b) for (a, b) in combinations(numbers, 2) if a < 1531 and b < 1531]
    for (a, b) in all_combinations:
        expected = 2020 - a - b
        if expected in numbers:
            print(a, b, expected)
            return
    print("No result")

if __name__ == "__main__":
    main()

