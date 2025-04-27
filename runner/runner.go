package runner

import (
	"io"
	"os"

	"github.com/Neel-shetty/monkey/evaluator"
	"github.com/Neel-shetty/monkey/lexer"
	"github.com/Neel-shetty/monkey/object"
	"github.com/Neel-shetty/monkey/parser"
)

func RunFile(filename string, out io.Writer) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	env := object.NewEnvironment()
	l := lexer.New(string(content))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			io.WriteString(out, "\t"+msg+"\n")
		}
		return nil
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}

	return nil
}
