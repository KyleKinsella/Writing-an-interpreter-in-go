package main

import (
	"fmt"
	"monkey/ast"
	// "monkey/token"
	"monkey/lexer"
	"monkey/parser"
	"monkey/object"
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

// func add(a, b int) (int, error) {
// 	return a + b, nil
// }

// func understandingEnv(node *ast.Identifier, env *object.Environment) object.Object {
// 	get, ok := env.Get(node.Value)
// 	if !ok {
// 		panic(ok)
// 	}
// 	return get
// }

func main() {
	// input := "let x = 5;"

	// lexer := lexer.New(input)
	// fmt.Println("lexer:", lexer)

	// for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
	// 	fmt.Println(tok)
 	// }

	// parser := parser.New(lexer)
	// fmt.Println("\nparser:", parser)

	input := "let x = 5; return x;"

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

	// add, err := add(1, 2)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("add:", add)


	// I don't really understand what the "environment" is. this is for the bindings

	env := object.NewEnvironment()
	fmt.Println("\n\nenv:", env)

	get, ok := env.Get("kyle")
	if !ok {
		panic(ok)
	}
	fmt.Println("get:", get)

	// understandingEnv(program, env)

	// get, ok := env.Get(program.String())
	// if !ok {
	// 	panic(ok)
	// }
	// fmt.Println("get:", get)
}