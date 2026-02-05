package compilerREPL

import (
    "bufio"
    "fmt"
    "io"
    "monkey/compiler"
    "monkey/lexer"
    "monkey/parser"
    "monkey/vm"
    "monkey/object"
)

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    
    constants := []object.Object{}
    globals := make([]object.Object, vm.GlobalSize)

    symbolTable := compiler.NewSymbolTable()

    for i, v := range object.Builtins {
        symbolTable.DefineBuiltin(i, v.Name)
    }
    
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
        
        comp := compiler.NewWithState(symbolTable, constants)
        err := comp.Compile(program)
        if err != nil {
            fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
            continue
        }
        
        code := comp.ByteCode()
        constants = code.Constants
        
        machine := vm.NewWithGlobalStore(code, globals)
        err = machine.Run()
        if err != nil {
            fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
        }
        
        stackTop := machine.LastPoppedStackElem()
        io.WriteString(out, stackTop.Inspect())
        io.WriteString(out, "\n")
    }
} 
