package lexer

import (
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "file", "../samples/sample_5.vy", "sample .vy file to process")
}

func TestNextToken(t *testing.T) {
	flag.Parse()

	file, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	lexer := Init(string(content))

	for {
		token := lexer.NextToken()
		fmt.Printf("<%s, \"%s\">\n", token.Type, token.Value)
		if token.Type == EOF {
			break
		}
	}
}
