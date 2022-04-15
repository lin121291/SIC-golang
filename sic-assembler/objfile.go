package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func hexstr(n string) string {
	h := strings.ToUpper(n)
	extra := 6 - len(h)
	for i := 0; i < extra; i++ {
		n = "0" + n
	}
	return n
}

//寫入H行
func writeHeader(s string, l string, p string) string {

	var name string
	print("enter your .obj file name \n")
	print(">>> ")
	fmt.Scanln(&name)

	//使用指定的標誌(O_RDONLY 等)打開命名文件。如果文件不存在，並且傳遞了 O_CREATE 標誌
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	tmp := "H" + s + hexstr(l) + hexstr(p)

	f.WriteString(tmp)
	f.WriteString("\n")
	f.Close()

	return name
}

//寫入E行
func WriteEnd(file string, n string) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	tmp := "E" + hexstr(n)

	f.WriteString(tmp)
	f.WriteString("\n")
	f.Close()
}

//寫入T行
func writeText(file string, s int, tline string) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}

	textrecord := "T" + hexstr(strconv.FormatInt(int64(s), 16))
	l := strconv.FormatInt(int64(int(len(tline)/2)), 16)

	n := 2 - len(l)
	for i := 0; i < n; i++ {
		l = "0" + l
	}

	textrecord += l
	textrecord += strings.ToUpper(tline)
	textrecord += "\n"
	textrecord = strings.ToUpper(textrecord)
	f.WriteString(textrecord)
	f.WriteString("\n")
	f.Close()
}
