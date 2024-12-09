package parser

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"
	"wavy/lexer"
)

// func checkLexerErrors(t *testing.T, l *lexer.Lexer, output_file *os.File) {
// 	errors := l.Errors()
// 	if len(errors) == 0 {
// 		return
// 	}
// 	fmt.Printf("lexer has %d error(s):\n", len(errors))
// 	output_file.WriteString(fmt.Sprintf("lexer has %d error(s):\n", len(errors)))
// 	for _, msg := range errors {
// 		fmt.Println(msg)
// 		output_file.WriteString(msg + "\n")
// 	}
// 	fmt.Print("\n==== Parser Output End ====\n\n")
// 	t.FailNow()

// }

func checkParserErrors_(t *testing.T, p *Parser, output_file *os.File) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	fmt.Printf("parser has %d error(s):\n", len(errors))
	output_file.WriteString(fmt.Sprintf("parser has %d error(s):\n", len(errors)))
	for _, msg := range errors {
		fmt.Println(msg)
		output_file.WriteString(msg + "\n")
	}
	fmt.Print("\n==== Parser Output End ====\n\n")
	t.FailNow()
}

var filePath string

func init() {
	flag.StringVar(&filePath, "file", "/wavy/parser/samples/sample_1.vy", "sample .vy file to process")
}

func TestParserOutput(t *testing.T) {
	flag.Parse()

	directory := filepath.Dir(filePath)
	fileName := filepath.Base(filePath)

	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		t.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	l := lexer.New(string(content))

	output_file, err := os.Create(filepath.Join(directory, fileName+".out"))
	if err != nil {
		t.Fatalf("error creating output file: %v", err)
	}
	defer output_file.Close()

	p := New(l)

	fmt.Print("\n==== Parser Output Start ====\n\n")

	program := p.ParseProgram()
	// checkLexerErrors(t, l, output_file)
	checkParserErrors_(t, p, output_file)

	if program == nil {
		log.Fatalf("error parsing program")
	}

	// fmt.Println(program.Tree("", true))

	// output_file.WriteString(program.Tree("", true))

	fmt.Print("\n==== Parser Output End ====\n\n")

}
