package go_design

import "demo/go-design/factory"

func GoDesignMain() {
	var jsonRuleFactory = factory.NewIRuleConfigParserFactory("json")
	var jsonRule = jsonRuleFactory.CreateParser()
	var data []byte = []byte{11,22}
	jsonRule.Parse(data)
}
