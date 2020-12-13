package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	fields map[string]string
}

type PassportList struct {
	passports []Passport
}

func NewPassportList(reader io.Reader) (PassportList, error) {
	scanner := bufio.NewScanner(reader)
	passports := make([]Passport, 0)
	var fields = make(map[string]string)
	for scanner.Scan() {
		l := scanner.Text()
		// we have a line
		if len(l) != 0 {
			// first we split by spaces
			parts := strings.Split(l, " ")
			for _, value := range parts {
				// we then split by :
				separatedParts := strings.Split(value, ":")
				fields[separatedParts[0]] = separatedParts[1]
			}
		}
		if len(l) == 0 {
			// this is a new passport, so reset fields and save the passport into a variable
			p := Passport{
				fields: fields,
			}
			passports = append(passports, p)
			fields = map[string]string{}
		}
	}
	// Careful at the end we might still have some fields read in so we need to store them as a passport as well
	if len(fields) > 0 {
		p := Passport{
			fields: fields,
		}
		passports = append(passports, p)
	}
	pl := PassportList{
		passports: passports,
	}
	return pl, nil
}

func (p Passport) isValid() bool {
	_, ok := p.fields["byr"]
	if !ok {
		return false
	}
	_, ok = p.fields["iyr"]
	if !ok {
		return false
	}

	_, ok = p.fields["eyr"]
	if !ok {
		return false
	}
	_, ok = p.fields["hgt"]
	if !ok {
		return false
	}
	_, ok = p.fields["hcl"]
	if !ok {
		return false
	}
	_, ok = p.fields["ecl"]
	if !ok {
		return false
	}
	_, ok = p.fields["pid"]
	if !ok {
		return false
	}
	return true
}

func (p Passport) isBirthYearValid() bool {
	birthYear, err := strconv.Atoi(p.fields["byr"])
	return len(p.fields["byr"]) == 4 && err == nil && birthYear >= 1920 && birthYear <= 2002

}

func (p Passport) isIssueYearValid() bool {
	issueYear, err := strconv.Atoi(p.fields["iyr"])
	return len(p.fields["iyr"]) == 4 && err == nil && issueYear >= 2010 && issueYear <= 2020
}

func (p Passport) isExpirationYearValid() bool {
	expirationYear, err := strconv.Atoi(p.fields["eyr"])
	return len(p.fields["eyr"]) == 4 && err == nil && expirationYear >= 2020 && expirationYear <= 2030
}

func (p Passport) isHeightValid() bool {
	if len(p.fields["hgt"]) < 3 {
		return false
	}
	height := p.fields["hgt"]
	heightEnding := height[len(height)-2:]

	if heightEnding != "cm" && heightEnding != "in" {
		return false
	}
	heightNumbers, err := strconv.Atoi(height[:len(height)-2])
	if err != nil {
		return false
	}

	if heightEnding == "cm" {
		return 150 <= heightNumbers && heightNumbers <= 193
	}

	if heightEnding == "in" {
		return 59 <= heightNumbers && heightNumbers <= 76
	}
	return true
}

func (p Passport) isHairColorValid() bool {
	hairColorRegEx := regexp.MustCompile(`^#[a-f0-9]{6}$`)
	hairColorValid := hairColorRegEx.Match([]byte(p.fields["hcl"]))
	return hairColorValid
}

func (p Passport) isEyeColorValid() bool {
	eyeColorValid := p.fields["ecl"] == "amb" || p.fields["ecl"] == "blu" || p.fields["ecl"] == "brn" || p.fields["ecl"] == "gry" || p.fields["ecl"] == "grn" || p.fields["ecl"] == "hzl" || p.fields["ecl"] == "oth"
	return eyeColorValid
}

func (p Passport) isPidValid() bool {
	pidRegEx := regexp.MustCompile(`^[0-9]{9}$`)
	pidValid := pidRegEx.Match([]byte(p.fields["pid"]))
	return pidValid
}

func (p Passport) isStrictlyValid() bool {

	if !p.isValid() {
		return false
	}
	return p.isBirthYearValid() && p.isIssueYearValid() && p.isExpirationYearValid() && p.isHairColorValid() && p.isEyeColorValid() && p.isPidValid() && p.isHeightValid()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func partOne() int {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	var validPassports int = 0
	passportList, error := NewPassportList(file)
	check(error)
	for _, passport := range passportList.passports {
		if passport.isValid() {
			validPassports++
		}
	}
	return validPassports
}

func partTwo() int {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	var validPassports int = 0
	passportList, error := NewPassportList(file)
	validPassportList := make([]Passport, 0)
	check(error)
	for _, passport := range passportList.passports {
		if passport.isStrictlyValid() {
			validPassports++
			validPassportList = append(validPassportList, passport)
		}
	}
	return validPassports
}

func main() {
	fmt.Println("Valid Passports", partOne())
	fmt.Println("Strictly  Valid Passports", partTwo())

}
