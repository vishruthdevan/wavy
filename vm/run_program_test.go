package vm

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
	"wavy/compiler"
	"wavy/lexer"
	"wavy/parser"
)

var filePath string

func init() {
	flag.StringVar(&filePath, "file", "/wavy/vm/samples/sample_1.vy", "sample .vy file to process")
}

func TestVMOutput(t *testing.T) {
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
	p := parser.New(l)
	program := p.ParseProgram()

	if l.Errors() != nil {
		for _, msg := range l.Errors() {
			t.Fatalf("lexer error: %v", msg)
		}
	}

	if p.Errors() != nil {
		for _, msg := range p.Errors() {
			t.Fatalf("parser error: %v", msg)
		}
	}

	comp := compiler.New()
	err = comp.Compile(program)
	if err != nil {
		t.Fatalf("compiler error(s): %s", err)
	}

	vm := New(comp.Bytecode())
	err = vm.Run()
	if err != nil {
		t.Fatalf("vm error(s): %s", err)
	}

	output_file, err := os.Create(filepath.Join(directory, fileName+".out"))
	if err != nil {
		t.Fatalf("error creating output file: %v", err)
	}
	defer output_file.Close()

	fmt.Print("\n==== VM Output Start ====\n\n")

	fmt.Print(vm.LastPoppedStackElem().Inspect())

	fmt.Print("\n\n==== VM Output End ====\n\n")

	// write to file
	output_file.WriteString(vm.LastPoppedStackElem().Inspect())

}
