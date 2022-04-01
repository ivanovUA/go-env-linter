package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	template := regexp.MustCompile(`^[A-Z]*=".+"$`)
	var data []string
	file, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !template.MatchString(scanner.Text()) {
			fmt.Println(scanner.Text())
			data = append(data, scanner.Text())
		}
	}
	if data != nil {
		os.Exit(1)
	} else {
		fmt.Println("No problems found")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
