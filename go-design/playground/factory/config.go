package factory

import "fmt"

type IRuleConfigParser interface {
	Parse()
}

type IRuleConfigFactory interface {
	RuleFactory()		IRuleConfigParser
}

type JsonRuleConfigParser struct {}
type YamlRuleConfigParser struct {}
type XmlRuleConfigParser struct {}
type JsonRuleConfigFactory struct {}
type YamlRuleConfigFactory struct {}
type XmlRuleConfigFactory struct {}

func (j JsonRuleConfigParser) Parse() {
	fmt.Println("json rule config parser")
}

func (y YamlRuleConfigParser) Parse() {
	fmt.Println("yaml rule config parser")
}

func (x XmlRuleConfigParser) Parse() {
	fmt.Println("xml rule config parser")
}

func (jf JsonRuleConfigFactory) RuleFactory() IRuleConfigParser {
	return JsonRuleConfigParser{}
}

func (yf YamlRuleConfigFactory) RuleFactory() IRuleConfigParser {
	return YamlRuleConfigParser{}
}

func (xf XmlRuleConfigFactory) RuleFactory() IRuleConfigParser {
	return XmlRuleConfigParser{}
}

func CreateFactory(t string) IRuleConfigFactory {
	switch t {
	case "json":
		return JsonRuleConfigFactory{}
	case "yaml":
		return YamlRuleConfigFactory{}
	case "xml":
		return XmlRuleConfigFactory{}
	}
	return nil
}

