package main

import (
	"log"
	"os"
)

func writeHeader(s string, l string, p string) {
	//使用指定的標誌(O_RDONLY 等)打開命名文件。如果文件不存在，並且傳遞了 O_CREATE 標誌
	f, err := os.OpenFile("example.obj", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	f.Write([]byte(s))
	f.Write([]byte(l))
	f.Write([]byte(p))
	f.Write([]byte("\n"))
	f.Close()
}
