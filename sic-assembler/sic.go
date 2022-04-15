package main

var OPTAB = map[string]int{
	"ADD":  0x18,
	"AND":  0x40,
	"COMP": 0x28,
	"DIV":  0x24,
	"J":    0x3C,
	"JEQ":  0x30,
	"JGT":  0x34,
	"JLT":  0x38,
	"JSUB": 0x48,
	"LDA":  0x00,
	"LDCH": 0x50,
	"LDL":  0x08,
	"LDX":  0x04,
	"MUL":  0x20,
	"OR":   0x44,
	"RD":   0xD8,
	"RSUB": 0x4C,
	"STA":  0x0C,
	"STCH": 0x54,
	"STL":  0x14,
	"STSW": 0xE8,
	"STX":  0x10,
	"SUB":  0x1C,
	"TD":   0xE0,
	"TIX":  0x2C,
	"WD":   0xDC,
	//26
}

var DIRECTIVE = [6]string{
	"START",
	"END",
	"WORD",
	"BYTE",
	"RESW",
	"RESB",
}

func check_DIRECTIVE(tmp string) bool {
	for _, d := range DIRECTIVE {
		if d == tmp {
			return true
		}
	}
	return false
}

func check_OPTAB(tmp string) bool {
	for op := range OPTAB {
		if tmp == op {
			return true
		}
	}
	return false
}

func take_opcode(n string) int {
	return OPTAB[n]
}
