package mail

import (
	"fmt"
	"node_metrics_go/global"
	"node_metrics_go/internal/message"
	mail "node_metrics_go/pkg/email"
	"strings"
)

type MailMessage struct {
	// toAddr           []string
	// subject, content string
	content string
}

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
	subject := strings.ToUpper(global.MetricsType) + "-" + global.Mailer.Subject
	if err := p.Mail.Send(subject, mm.GetContent(), global.Mailer.To); err != nil {
		return err
	}
	return nil
}
