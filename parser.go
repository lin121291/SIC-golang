package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func load_file(f *[]string) {

	file, err := os.Open("example.asm")
	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	//返回scanner型別
	fileScanner := bufio.NewScanner(file)

	// read line by line
	for fileScanner.Scan() {
		*f = append(*f, fileScanner.Text()) //新增每一行到tmp
	}
}

func examine_line(line string) {

	element := strings.Fields(strings.TrimSpace(line)) //切割line
}
