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
			if check_symtab(c[0], symtab) {
				print("Your assembly code has problem.")
				continue
			}
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

	for i := range symtab {
		fmt.Printf("%s : %d\n", i, symtab[i])
	}

	//pass2
	//輸出objectfile
	reserveflag := false
	FileName := "" //
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
			FileName = writeHeader(c[0], c[2], r)

			//設定start locctr tstart
			tmp, _ := strconv.ParseInt(c[2], 16, 16) //16進制轉10進制
			locctr = int(tmp)
			tstart = int(tmp)
			start = int(tmp)
		}

		if c[1] == "END" {
			if len(tline) > 0 {
				writeText(FileName, tstart, tline)
			}

			proglen = locctr - start
			address := start
			if c[2] != " " {
				address = symtab[c[2]]
			}

			r := strconv.FormatInt(int64(address), 16) //start 10轉16
			WriteEnd(FileName, r)
			break
		}

		//指令們
		if check_OPTAB(c[1]) == true {
			instruction := generateInstruction(c[1], c[2], symtab)

			if locctr+3-tstart > 30 || reserveflag == true {

				writeText(FileName, tstart, tline)
				tstart = locctr
				tline = instruction
			} else {
				tline += instruction
			}

			reserveflag = false
			locctr += 3

		} else if c[1] == "WORD" {
			constant = hexstr(c[2])

			if locctr+3-tstart > 30 || reserveflag == true {
				writeText(FileName, tstart, tline)
				tstart = locctr
				tline = constant

			} else {
				tline += constant
			}

			reserveflag = false
			locctr += 3

		} else if c[1] == "RESW" {
			tmp, _ := strconv.ParseInt(c[2], 16, 16)
			locctr += (int(tmp) * 3)
			reserveflag = true

		} else if c[1] == "RESB" {
			tmp, _ := strconv.ParseInt(c[2], 16, 16)
			locctr += int(tmp)
			reserveflag = true

		} else if c[1] == "BYTE" {
			if string(c[2][0]) == "X" {
				operandlen = int((len(c[2]) - 3) / 2)
				constant = c[2][2 : len(c[2])-1]

			} else if string(c[2][0]) == "C" {
				operandlen = int(len(c[2]) - 3)
				constant = processBYTEC(c[2])

			}

			if locctr+3-tstart > 30 || reserveflag == true {
				writeText(FileName, tstart, tline)
				tstart = locctr
				tline = constant

			} else {
				tline += constant

			}

			reserveflag = false
			locctr += operandlen

		}

	}

}
