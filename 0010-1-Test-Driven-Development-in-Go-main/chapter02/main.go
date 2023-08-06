package main

import (
	"flag"
	"log"

	"github.com/PacktPublishing/Test-Driven-Development-in-Go/chapter02/calculator"
	"github.com/PacktPublishing/Test-Driven-Development-in-Go/chapter02/input"
)

// go run chapterXX/main.go -expression "2 + 3"
func main() {
	expr := flag.String("expression", "", "mathematical expression to parse")
	flag.Parse()

	engine := calculator.NewEngine()
	validator := input.NewValidator(engine.GetNumOperands(), engine.GetValidOperators())
	parser := input.NewParser(engine, validator)
	result, err := parser.ProcessExpression(*expr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*result)
}
