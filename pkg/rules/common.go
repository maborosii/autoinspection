package rules

type BaseRule struct{}

type RuleItf interface {
	GetRuleJob() string
}

type RuleOption func(RuleItf) bool

func (b *BaseRule) GetRuleJob() string {
	return "Basic Rules"
}
