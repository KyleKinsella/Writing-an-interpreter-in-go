package compilerREPL

import (
	"bufio"
	"fmt"
	"io"
	"monkey/compiler"
	"monkey/lexer"
	"monkey/parser"
	"monkey/vm"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	
	for {
		fmt.Fprintf(out, ">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			//printParserErrors(out, p.Errors())
			continue
		}
		
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
			continue
		}
		
		machine := vm.New(comp.ByteCode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
		}
		
		stackTop := machine.StackTop()
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")
	}
} 
