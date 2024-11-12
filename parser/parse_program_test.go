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

func checkLexerErrors(t *testing.T, l *lexer.Lexer) {
	errors := l.Errors()
	if len(errors) == 0 {
		return
	}
	fmt.Printf("lexer has %d errors:\n", len(errors))
	for _, msg := range errors {
		fmt.Println(msg)
	}
	fmt.Print("\n==== Parser Output End ====\n\n")
	t.FailNow()

}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	fmt.Printf("parser has %d errors:\n", len(errors))
	for _, msg := range errors {
		fmt.Println(msg)
	}
	t.FailNow()

	fmt.Print("\n==== Parser Output End ====\n\n")
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

	l := lexer.Init(string(content))

	output_file, err := os.Create(filepath.Join(directory, fileName+".out"))
	if err != nil {
		t.Fatalf("error creating output file: %v", err)
	}
	defer output_file.Close()

	p := Init(l)

	fmt.Print("\n==== Parser Output Start ====\n\n")

	program := p.ParseProgram()
	checkLexerErrors(t, l)
	checkParserErrors(t, p)

	if program == nil {
		log.Fatalf("error parsing program")
	}

	fmt.Println(program.Tree("", true))

	output_file.WriteString(program.Tree("", true))

	fmt.Print("\n==== Parser Output End ====\n\n")

}
