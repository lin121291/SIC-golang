package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//讀進檔案
func load_file(f *[]string) {

	var x string
	print("enter your .asm file \n")
	print(">>> ")
	fmt.Scanln(&x)
	file, err := os.Open(x)

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

//把每一行組語切割確認語法正確
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

//確認DIRECTIVE OPCODE存在
func checkOpDir(n string) bool {
	if check_DIRECTIVE(n) == true {
		return true
	}
	if check_OPTAB(n) == true {
		return true
	}
	return false
}

//確認symtab
func check_symtab(c2 string, sym map[string]int) bool {
	for i := range sym {
		if c2 == i {
			return true
		}
	}
	return false
}

//指令格式化
func generateInstruction(c1 string, c2 string, sym map[string]int) string {
	ins := take_opcode(c1) * 65536

	if c2 != " " {

		if check_symtab(c2, sym) == true {
			ins += int(sym[c2])
		} else {
			return ""
		}
	}

	return hexstr(strconv.FormatInt(int64(ins), 16)) //10進位轉16進位
}

//BYTE處理
func processBYTEC(n string) string {
	constant := ""
	for i := 2; i <= len(n)-1; i++ {
		t, _ := strconv.ParseInt(n, 16, 16)
		tmp := strconv.Itoa(int(t))

		if len(strconv.FormatInt(t, 16)) == 1 {
			tmp = "0" + tmp
		}
		tmp = strings.ToUpper(tmp)
		constant += tmp
	}
	return constant
}
