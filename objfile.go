package main

import (
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

func writeHeader(s string, l string, p string) {
	//使用指定的標誌(O_RDONLY 等)打開命名文件。如果文件不存在，並且傳遞了 O_CREATE 標誌
	f, err := os.OpenFile("example.obj", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	tmp := "H" + s + hexstr(l) + hexstr(p)

	f.WriteString(tmp)
	f.WriteString("\n")
	f.Close()
}

func WriteEnd(n string) {
	f, err := os.OpenFile("example.obj", os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	tmp := "E" + hexstr(n)

	f.WriteString(tmp)
	f.WriteString("\n")
	f.Close()
}

func writeText(s int, tline string) {
	f, err := os.OpenFile("example.obj", os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}

	textrecord := "T" + hexstr(strconv.FormatInt(int64(s), 16))
	l := strconv.FormatInt(int64(int(len(tline)/2)), 16)

	n := 2 - len(l)
	for i := 0; i <= n; i++ {
		l = "0" + l
	}

	l = strings.ToUpper(l)
	textrecord += l
	textrecord += tline
	textrecord += "\n"
	f.WriteString(textrecord)
	f.WriteString("\n")
	f.Close()
}
