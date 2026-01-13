package main

import (
	"monkey/lexer"
	"monkey/parser"
	"monkey/ast"
	//"monkey/code"
	"monkey/compiler"
	"monkey/vm"
	"fmt"
)

func learning(input string) {
	
	l := lexer.New(input)
	p := parser.New(l)
	
	pro := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return
	}
	
	com := compiler.New()
	
	compile := com.Compile(pro)
	
	machine := vm.New(com.ByteCode())
	
	run := machine.Run()
	
	fmt.Println("compile:", compile, "\n", "machine:", machine, "\n", "run:", run)
}

func setup(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func bytecodeStuff() {
}

func compilerStuff(input string) {	
	
	start := setup(input)
	
	//com := compiler.Compiler{}
	
	//by := compiler.ByteCode{}
	
	comNew := compiler.New()
	
	compile := comNew.Compile(start)
	
	fmt.Println(compile)
}

func vmStuff() {
}


func main() {
	//compilerStuff("let x = 5;")
	
	//learning("let x = 5;")
	
	learning("1 + 2")
}
