package lexer

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestNextToken(t *testing.T) {
	file, err := os.Open("../samples/sample_1.vy")
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
