package types

type RequestYaml struct {
	Host          string           `yaml:"host"`
	Path          string           `yaml:"path"`
	MatchRules    []*MatchRuleYaml `yaml:"matchRules"`
	ParserName    string           `yaml:"parserName"`
	PathRule      string           `yaml:"pathRule"`
	NeedMatchPath string           `yaml:"needMatchPath"`
}

type MatchRuleYaml struct {
	Rule string `yaml:"rule"`
	Keys string `yaml:"keys"`
}

type Parser interface {
	Parse(contents []byte, request *Request) (*ParseResult, error)
	InitRegexp(request *Request) error
	Serialize() (name string, args interface{})
}
