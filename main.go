//note : go run *.go

package main

import (
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
	proglen := 0

	for _, line := range file {

		c := [3]string{" ", " ", " "}
		ExamineLine(line, &c)

		if c[1] == "START" {
			tmp, _ := strconv.ParseInt(c[2], 16, 16)
			start = int(tmp)
			locctr = start
		}
		if c[1] == "END" {
			proglen = locctr - start
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
			tmp, _ := strconv.ParseInt(c[2], 16, 16)
			locctr = locctr + int(tmp)*3
		} else if c[1] == "RESB" {
			tmp, _ := strconv.ParseInt(c[2], 16, 16)
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
	/*
		for i := range symtab {
			fmt.Printf("%s : %d\n", i, symtab[i])
		}
	*/

	//pass2
	//輸出objectfile
	locctr = 0
	tline := ""
	tstart := 0
	constant := "" //
	operandlen := 0

	for _, line := range file {

		c := [3]string{" ", " ", " "}
		ExamineLine(line, &c)

		if c[1] == "START" {

			//把第一行組好
			r := strconv.FormatInt(int64(proglen), 16) //proglen轉16
			writeHeader(c[0], c[2], r)

			//設定start locctr tstart
			locctr, _ := strconv.ParseInt(c[2], 16, 16) //16進制轉10進制
			tstart = int(locctr)
			start = int(locctr)
		}

		if c[1] == "END" {

			if len(tline) > 0 {
				writeText(tstart, tline) //重要
			}
			r := strconv.FormatInt(int64(start), 16) //start 10轉16
			WriteEnd(r)
		}

		//指令們
		if check_OPTAB(c[1]) == true {
			//
			//會輸出這行指令
			instruction := generateInstruction(c[1], c[2], symtab)

			//缺
			if locctr+3-tstart > 30 {

				writeText(tstart, tline)
				tstart = locctr
				tline = instruction
			} else {
				tline += instruction
			}
			locctr += 3

		} else if c[1] == "WORD" {
			//
			constant = hexstr(c[2])

			if locctr+3-tstart > 30 {
				//
				writeText(tstart, tline)
				tstart = locctr
				tline = constant
			} else {
				tline += constant
			}

			locctr += 3

		} else if c[1] == "RESW" {
			tmp, _ := strconv.ParseInt(c[2], 16, 16)
			locctr += int(tmp * 3)
		} else if c[1] == "RESB" {
			tmp, _ := strconv.ParseInt(c[2], 16, 16)
			locctr += int(tmp)
		} else if c[1] == "BYTE" {
			if string(c[2][0]) == "X" {
				//
				operandlen = int((len(c[2]) - 3) / 2)
				constant = c[2][2 : len(c[2])-1]
			} else if string(c[2][0]) == "C" {
				//
				operandlen = int(len(c[2]) - 3)
				constant = processBYTEC(c[2])
			}

			if locctr+3-tstart > 30 {
				//
				writeText(tstart, tline)
				tstart = locctr
				tline = constant
			} else {
				tline += constant
			}

			locctr += operandlen

		}

	}

}
