//note : go run *.go

package main

func main() {

	var file []string //建立 Slice
	load_file(&file)

	//pass1
	//pass1 build symtab
	for _, line := range file {

		examine_line(line)

	}
}
