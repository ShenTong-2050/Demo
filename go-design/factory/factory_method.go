package factory

// IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
	CreateParser()		IRuleConfigParser
}

type jsonRuleConfigParserFactory struct {}

func (J jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

type yamlRuleConfigParserFactory struct {}

func (Y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

func NewIRuleConfigParserFactory(s string) IRuleConfigParserFactory {
	switch s {
		case "json":
			return jsonRuleConfigParserFactory{}
		case "yaml":
			return yamlRuleConfigParserFactory{}
	}
	return nil
}


