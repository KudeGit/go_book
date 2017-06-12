package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func dup2() {
	// open a file
	if file, err := os.Open("test.txt"); err == nil {

		// make sure it gets closed
		defer file.Close()

		// create a new scanner and read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			log.Println(scanner.Text())
		}

		// check for errors
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

	} else {
		log.Fatal(err)
	}
	file2, err := os.Open("test.txt")
	if err == nil {
		log.Println("Fuck")
	}
	scanner2 := bufio.NewScanner(file2)
	for scanner2.Scan() {
		log.Println(scanner2.Text())
	}

}

func countDuplicate(file *os.File, lines map[string]int) {
	if file == nil || lines == nil {
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines[scanner.Text()]++
	}
}

type conditionFunc func(k string, v int) bool

func notDuplicate(k string, v int) bool {
	if v <= 1 {
		return true
	}
	return false
}

func printMap(M *map[string]int, F conditionFunc) {
	for k, v := range *M {
		if F(k, v) {
			fmt.Printf("%s: %d\n", k, v)
		}
	}
}

func openFilesAndCount(filenames *[]string) {
	lines := make(map[string]int)
	for _, filename := range *filenames {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
			return
		}
		countDuplicate(file, lines)
	}
	notDupF := notDuplicate
	printMap(&lines, notDupF)
}

func dup1() {
	lines := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		lines[input.Text()]++
	}

	for line, count := range lines {
		fmt.Printf("%s: %d\n", line, count)
	}
}

func readFileAndSplit(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range strings.Split(string(data), "\n") {
		fmt.Println(line)
	}

}

func main() {
	//filenames := os.Args[1:]
	//openFilesAndCount(&filenames)
	readFileAndSplit("test.txt")
}
