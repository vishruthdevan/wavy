package lexer

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

var filePath string

func init() {
	flag.StringVar(&filePath, "file", "../samples/sample_1.vy", "sample .vy file to process")
}

func TestLexer(t *testing.T) {
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

	lexer := Init(string(content))

	output_file, err := os.Create(filepath.Join(directory, fileName+".out"))
	if err != nil {
		t.Fatalf("error creating output file: %v", err)
	}
	defer output_file.Close()

	fmt.Print("\n==== Lexer Output Start ====\n\n")

	for {
		token := lexer.NextToken()
		fmt.Printf("<%s, \"%s\">\n", token.Type, token.Value)
		output_file.WriteString(fmt.Sprintf("<%s, \"%s\">\n", token.Type, token.Value))
		if token.Type == EOF {
			break
		}
	}

	fmt.Print("\n==== Lexer Output End ====\n\n")
}
