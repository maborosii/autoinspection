package mail

import (
	"fmt"
	"node_metrics_go/internal/message"
	mail "node_metrics_go/pkg/email"
)

type MailMessage struct {
	// toAddr           []string
	// subject, content string
	content string
}

// func (m *MailMessage) GetToAddr() []string {
// 	return m.toAddr
// }

// func (m *MailMessage) GetSubject() string {
// 	return m.subject
// }
func NewMailMessage(c string) *MailMessage {
	return &MailMessage{
		content: c,
	}
}

func (m *MailMessage) GetContent() string {
	return m.content
}

func (m *MailMessage) RealText() {
}

type MailPusher struct {
	Mail *mail.Mail
}

func (p *MailPusher) Type() string {
	return "mail"
}

func (p *MailPusher) Push(m message.OutMessage) error {
	mm, ok := m.(*MailMessage)
	if !ok {
		return fmt.Errorf("mail message asset failed")
	}
	if err := p.Mail.Send("主机巡检详情", mm.GetContent(), []string{"liushuai@leyaoyao.com"}); err != nil {
		// if err := p.Mail.Send(mm.GetSubject(), mm.GetContent(), mm.GetToAddr()); err != nil {
		return err
	}
	return nil
}
