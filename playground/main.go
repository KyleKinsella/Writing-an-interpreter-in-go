package main

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/parser"
	// "monkey/token"
)

func printNode(node ast.Node, indent string) {
	switch n := node.(type) {
	case *ast.Program:
		for _, stmt := range n.Statements {
			printNode(stmt, indent)
		}
	case *ast.LetStatement:
		fmt.Println(indent + "LetStatement:")
		fmt.Println(indent + "  Name:", n.Name.Value)
		fmt.Println(indent + "  Value:")
		printNode(n.Value, indent + "    ")
	
	case *ast.ReturnStatement:
		fmt.Println(indent, "ReturnStatement:")
		fmt.Println(indent, "  ReturnValue:")
		printNode(n.ReturnValue, indent + "    ")
	
	case *ast.ExpressionStatement:
		printNode(n.Expression, indent)
	
	case *ast.InfixExpression:
		fmt.Println(indent, "InfixExpression", n.Operator)
		fmt.Println(indent + "  Left:")
		printNode(n.Left, indent + "   ")
		fmt.Println(indent + "  Right:")
		printNode(n.Right, indent + "   ")

	case *ast.IntegerLiteral:
		fmt.Println(indent + "IntegerLiteral:", n.Value)
	
	case *ast.Identifier:
		fmt.Println(indent + "Identifier:", n.Value)

	default:
		fmt.Println(indent + "Unknown node type:", fmt.Sprintf("%T",n))
	}
}

func main() {
	// input := "let x = 5; return x;"
	input := "5 + 2 * 3"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram() // parse into the ast

	if len(p.Errors()) != 0 {
		fmt.Println("Parser errors:")
		for _, err := range p.Errors() {
			fmt.Println("\t", err)
		}
		return
	}

	fmt.Println("AST:")
	printNode(program, "")



	// fmt.Println("\n\n\n\n\nprogram:\n", program, "\n")

	// for _, p := range program.Statements {
		// fmt.Println("program.Statement contains:", p)
	// }

	// let := ast.LetStatement{Token: token.Token{Type: token.LET, Literal: "x"}}
	// fmt.Println(let)

	// let.Name = &ast.Identifier{Value: "test"}
	// fmt.Println(let.Name)

	// let.Value = let.Token.Literal()

	// let := ast.LetStatement{Token: token.Token{Type: token.LET, Literal: "x"}}
	// let.Name = &ast.Identifier{Value: "test"}
	// let.Value = let.

	// fmt.Println(let.String())
}