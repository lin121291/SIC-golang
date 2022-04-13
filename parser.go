package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func load_file(f *[]string) {

	//這邊要加東西
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

func ExamineLine(line string, c *[3]string) {

	command := strings.Fields(strings.TrimSpace(line)) //切割line

	//先不處理

	//在只有一個指令的情況下
	//ex. RSUB
	if len(command) == 1 {
		if checkOpDir(command[0]) == false {
			println("Your assembly code has problem.")
		}
		c[1] = command[0]
	}
	if len(command) == 2 {
		if checkOpDir(command[0]) == true {
			c[1] = command[0]
			c[2] = command[1]
		} else if checkOpDir(command[1]) == true {
			c[0] = command[0]
			c[1] = command[1]
		} else {
			println("Your assembly code has problem.")
		}
	}
	if len(command) == 3 {
		if checkOpDir(command[1]) == true {
			c[0] = command[0]
			c[1] = command[1]
			c[2] = command[2]
		} else {
			println("Your assembly code has problem.")
		}
	}
}

func checkOpDir(n string) bool {
	if check_DIRECTIVE(n) == true {
		return true
	}
	if check_OPTAB(n) == true {
		return true
	}
	return false
}
