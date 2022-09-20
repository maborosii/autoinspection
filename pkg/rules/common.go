package rules

type BaseRule struct{}

type RuleItf interface {
	GetJob() string
}

func (b *BaseRule) GetJob() string {
	return "Basic Rules"
}
