package factory

import "fmt"

type IAbstractRuleConfigParser interface {
	Parser()
}

func (a AbstractJsonRuleConfigParser) Parser() {
	fmt.Println("abstract json rule")
}

type IRuleSystemConfigParser interface {
	ParserSystem()
}

func (j JsonSystemConfigParser) ParserSystem() {
	fmt.Println("json system parser")
}

type IConfigParserFactory interface {
	CreateRuleParser()	IAbstractRuleConfigParser
	CreateSystemParser() IRuleSystemConfigParser
}

type AbstractJsonRuleConfigParser struct {}

type JsonSystemConfigParser struct {}

type JsonConfigParserFactory struct {}

func (j JsonConfigParserFactory) CreateRuleParser() IAbstractRuleConfigParser {
	return AbstractJsonRuleConfigParser{}
}

func (j JsonConfigParserFactory) CreateSystemParser() IRuleSystemConfigParser {
	return JsonSystemConfigParser{}
}