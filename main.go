//note : go run *.go

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var file []string //建立 Slice
	load_file(&file)

	//pass1
	//pass1 build symtab

	symtab := map[string]int{}
	start := 0
	locctr := 0
	//proglen := 0

	for _, line := range file {

		c := [3]string{" ", " ", " "}
		ExamineLine(line, &c)

		if c[1] == "START" {
			tmp, _ := strconv.ParseUint(c[2], 16, 32)
			start = int(tmp)
			locctr = start
		}
		if c[1] == "END" {
			//proglen = locctr - start
			break
		}
		if c[0] != " " { //最左邊的要加入symtab
			//缺
			symtab[c[0]] = locctr
		}

		if check_OPTAB(c[1]) == true {
			locctr = locctr + 3
		} else if c[1] == "WORD" {
			locctr = locctr + 3
		} else if c[1] == "RESW" {
			tmp, _ := strconv.ParseUint(c[2], 16, 32)
			locctr = locctr + int(tmp)*3
		} else if c[1] == "RESB" {
			tmp, _ := strconv.ParseUint(c[2], 16, 32)
			locctr = locctr + int(tmp)
		} else if c[1] == "BYTE" {
			if string(c[2][0]) == "X" {
				locctr = locctr + (len(c[2]) - 3)
			}
			if string(c[2][0]) == "C" {
				locctr = locctr + ((len(c[2]) - 3) / 2)
			}
		}
	}

	for i := range symtab {
		fmt.Printf("%s : %d\n", i, symtab[i])
	}

	/*
		//pass2
		//輸出objectfile
		locctr=0
		for _, line := range file {

			c := [3]string{" ", " ", " "}
			examine_line(line)

			if c[1]=="START"{

				tmp, _ := strconv.ParseUint(c[2], 16, 32)

				str:=c[0]
				locctr := int(tmp)
				writeHeader(str,string(locctr),string(proglen))
			}

			if c[1]=="END"{

			}

			if check_OPTAB(c[1]) == True {

			}else if c[1]=="WORD"{

			}else if c[1]=="RESW"{

			}else if c[1]=="RESB"{

			}else if c[1]=="BYTE"{
				if {

				}
				if{

				}
			}

		}
	*/
}
