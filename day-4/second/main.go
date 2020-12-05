package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValidByr(str string) bool {
	val, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return val >= 1920 && val <= 2002
}

func isValidIyr(str string) bool {
	val, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return val >= 2010 && val <= 2020
}

func isValidEyr(str string) bool {
	val, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return val >= 2020 && val <= 2030
}

func isAlphanumeric(char int) bool {
	return (char >= 'a' && char <= 'f') || (char >= '0' && char <= '9')
}

func isValidHcl(str string) bool {
	if str[0] != '#' {
		return false
	}
	for _, ch := range str[1:] {
		if !isAlphanumeric(int(ch)) {
			return false
		}
	}
	return true
}

func isValidEcl(str string) bool {
	return str == "amb" || str == "blu" || str == "brn" || str == "gry" || str == "grn" || str == "hzl" || str == "oth"
}

func isValidHgt(str string) bool {
	if !strings.HasSuffix(str, "cm") && !strings.HasSuffix(str, "in") {
		return false
	}
	suffix := str[len(str)-2:]
	preffix := str[:len(str)-2]
	value, err := strconv.Atoi(preffix)
	if err != nil {
		return false
	}

	if suffix == "cm" {
		return value >= 150 && value <= 193
	}
	return value >= 59 && value <= 76
}

func isNumber(str rune) bool {
	return int(str) >= '0' && int(str) <= '9'
}

func isValidPid(str string) bool {
	if len(str) != 9 {
		return false
	}
	for _, ch := range str {
		if !isNumber(ch) {
			return false
		}
	}
	return true
}

func isCompletePassport(passport map[string]string) bool {
	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, key := range keys {
		if passport[key] == "" {
			return false
		}
		switch key {
		case "byr":
			if !isValidByr(passport[key]) {
				return false
			}
			break
		case "iyr":
			if !isValidIyr(passport[key]) {
				return false
			}
			break
		case "eyr":
			if !isValidEyr(passport[key]) {
				return false
			}
			break
		case "hgt":
			if !isValidHgt(passport[key]) {
				return false
			}
			break
		case "hcl":
			if !isValidHcl(passport[key]) {
				return false
			}
			break
		case "ecl":
			if !isValidEcl(passport[key]) {
				return false
			}
			break
		case "pid":
			if !isValidPid(passport[key]) {
				return false
			}
			break
		}
	}
	fmt.Printf("\n")
	return true
}

func createPassport() map[string]string {
	return map[string]string{
		"byr": "",
		"iyr": "",
		"eyr": "",
		"hgt": "",
		"hcl": "",
		"ecl": "",
		"pid": "",
		"cid": "",
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
		var passport map[string]string = createPassport()
		var line string = scanner.Text()

		lineChunks := strings.Split(line, " ")
		for _, chunk := range lineChunks {
			if len(strings.TrimSpace(chunk)) == 0 {
				continue
			}
			keyAndValue := strings.Split(chunk, ":")
			passport[keyAndValue[0]] = keyAndValue[1]
		}

		if isCompletePassport(passport) {
			validPassportCount++
		}
	}
	fmt.Println("Valid passports", validPassportCount)
}
