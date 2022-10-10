package message

type OutMessage interface {
	RealText()
}

type OutMessageTemplate interface {
	GetSentence() []string
}
type InputMessage interface {
	Spread(tagType string, keys ...string) (map[string]interface{}, error)
}
type Render interface {
	Rend(n InputMessage, omt OutMessageTemplate) (OutMessage, error)
}
