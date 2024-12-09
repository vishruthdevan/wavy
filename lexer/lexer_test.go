package lexer

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
	"wavy/token"
)

var filePath string

func init() {
	flag.StringVar(&filePath, "file", "/wavy/lexer/samples/sample_1.vy", "sample .vy file to process")
}

func TestLexerOutput(t *testing.T) {
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

	lexer := New(string(content))

	output_file, err := os.Create(filepath.Join(directory, fileName+".out"))
	if err != nil {
		t.Fatalf("error creating output file: %v", err)
	}
	defer output_file.Close()

	fmt.Print("\n==== Lexer Output Start ====\n\n")

	for {
		t := lexer.NextToken()
		fmt.Printf("<%s, \"%s\">\n", t.Type, t.Literal)
		output_file.WriteString(fmt.Sprintf("<%s, \"%s\">\n", t.Type, t.Literal))
		if t.Type == token.EOF {
			break
		}
	}
	if lexer.errors != nil {
		fmt.Printf("lexer has %d error(s):\n", len(lexer.errors))
		output_file.WriteString(fmt.Sprintf("lexer has %d error(s):\n", len(lexer.errors)))
		for _, msg := range lexer.errors {
			fmt.Println(msg)
			output_file.WriteString(msg + "\n")
		}
	}

	fmt.Print("\n==== Lexer Output End ====\n\n")
}
