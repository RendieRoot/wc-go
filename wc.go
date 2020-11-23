package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var args = os.Args[1:]

	if len(args) == 0 || includes(args, "-") {
		fmt.Println("without arguments")
	} else if includes(args, "--help") {
		help()
	} else if includes(args, "--version") {
		version()
	} else {
		if len(args) == 1 {
			file := args[0]
			if _, err := os.Stat(file); err == nil {
				var result []int

				result = append(result, bytes(file))
				result = append(result, chars(file))
				result = append(result, lines(file))
				result = append(result, words(file))

				resultString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), " "), "[]")
				fmt.Println(resultString, file)

			} else {
				fmt.Println("No such file or directory:", file)
			}
		} else {
			file := args[1]

			if _, err := os.Stat(file); err == nil {
				argument := args[0]

				switch argument {
				case "-c", "--bytes":
					bytes := bytes(file)
					fmt.Println(bytes, file)
				case "-m", "--chars":
					chars := chars(file)
					fmt.Println(chars, file)
				case "-l", "--lines":
					lines := lines(file)
					fmt.Println(lines, file)
				case "-w", "--words":
					words := words(file)
					fmt.Println(words, file)
				default:
					var result []int

					if strings.Contains(argument, "c") {
						result = append(result, bytes(file))
					}

					if strings.Contains(argument, "m") {
						result = append(result, chars(file))
					}

					if strings.Contains(argument, "l") {
						result = append(result, lines(file))
					}

					if strings.Contains(argument, "w") {
						result = append(result, words(file))
					}

					if len(result) == 0 {
						warning := fmt.Sprintf("wc: invalid option -- '%v'\nTry '--help' for more information.", argument)
						fmt.Println(warning)
					} else {
						resultString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), " "), "[]")
						fmt.Println(resultString, file)
					}
				}
			} else {
				fmt.Println("No such file or directory:", file)
			}
		}
	}
}

func includes(arr []string, str string) bool {
	for _, argument := range arr {
		if argument == str {
			return true
		}
	}
	return false
}

func bytes(file string) int {
	info, error := os.Stat(file)

	if error == nil {
		bytes := int(info.Size())
		return bytes
	}

	return 0
}

func lines(file string) int {
	content, error := os.Open(file)

	if error == nil {
		scanner := bufio.NewScanner(content)
		lines := 0
		for scanner.Scan() {
			lines++
		}
		return lines
	}

	return 0
}

func chars(file string) int {
	content, error := ioutil.ReadFile(file)

	if error == nil {
		chars := len(content)
		return chars
	}

	return 0
}

func words(file string) int {
	content, error := ioutil.ReadFile(file)

	if error == nil {
		text := string(content)
		array := strings.Split(text, " ")
		words := len(array)
		return words
	}

	return 0
}

func version() {
	version := "version: SuperPuperShitGo wc 1.0"
	fmt.Println(version)
}

func help() {
	help := `Usage: go run wc.go [OPTION]... [FILE]...

	-c, --bytes            print the byte counts
	-m, --chars            print the character counts
	-l, --lines            print the newline counts
	-w, --words            print the word counts
		--help             display this help and exit
		--version          output version information and exit`
	fmt.Println(help)
}
