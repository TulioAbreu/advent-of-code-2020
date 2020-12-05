package first

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isCompletePassport(passport map[string]bool) bool {
	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, key := range keys {
		if !passport[key] {
			return false
		}
	}
	return true
}

func createPassport() map[string]bool {
	return map[string]bool{
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false,
		"cid": false,
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Could not read input file")
		return
	}

	var validPassportCount int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var passport map[string]bool = createPassport()
		var line string = scanner.Text()

		lineChunks := strings.Split(line, " ")
		for _, chunk := range lineChunks {
			if len(strings.TrimSpace(chunk)) == 0 {
				continue
			}
			keyAndValue := strings.Split(chunk, ":")
			passport[keyAndValue[0]] = true
		}

		if isCompletePassport(passport) {
			validPassportCount++
		}
	}
	fmt.Println("Valid passports", validPassportCount)
}
