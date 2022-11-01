package email

import (
	"log"

	"github.com/pkg/errors"
	gomail "gopkg.in/gomail.v2"
)

// Mail easy way to send basic email
type Mail struct {
	host               string
	port               int
	username, password string
}

func NewMail(host string, port int) *Mail {
	log.Println("try to send mail,host:", host, "port:", port)
	return &Mail{
		host: host,
		port: port,
	}
}

// Login login to SMTP server
func (m *Mail) Login(username, password string) {
	log.Println("login username: ", username)
	m.username = username
	m.password = password
}

// BuildMessage implement
func (m *Mail) BuildMessage(msg string) string {
	return msg
}

// EmailDialer create gomail.Dialer
// 自定义接口用于兼容 gomail 的 dailer
type EmailDialer interface {
	DialAndSend(m ...*gomail.Message) error
}

type mailSendOpt struct {
	dialerFact func(host string, port int, username, passwd string) EmailDialer
}

// 定义默认的 dialer --> gomail.Dialer
func (o *mailSendOpt) fillDefault() *mailSendOpt {
	o.dialerFact = func(host string, port int, username, passwd string) EmailDialer {
		return gomail.NewDialer(host, port, username, passwd)
	}

	return o
}

// 执行一组应用于 mailSendOpt 的函数
func (o *mailSendOpt) applyOpts(optfs []MailSendOptFunc) *mailSendOpt {
	for _, optf := range optfs {
		optf(o)
	}
	return o
}

// MailSendOptFunc is a function to set option for Mail.Send
type MailSendOptFunc func(*mailSendOpt)

// WithMailSendDialer set gomail.Dialer
// mailSendOpt 赋值
func WithMailSendDialer(dialerFact func(host string, port int, username, passwd string) EmailDialer) MailSendOptFunc {
	return func(opt *mailSendOpt) {
		opt.dialerFact = dialerFact
	}
}

// Send send email
func (m *Mail) Send(subject, content string, toAddr []string, optfs ...MailSendOptFunc) (err error) {
	// 默认 gomail.dialer
	opt := new(mailSendOpt).fillDefault().applyOpts(optfs)
	log.Println("send email toAddr:", toAddr)
	s := gomail.NewMessage()
	s.SetHeader("From", m.username)
	s.SetHeader("To", toAddr...)
	s.SetHeader("Subject", subject)
	s.SetBody("text/html", content)
	// s.SetBody("text/plain", content)

	// 获取 dialer
	dialer := opt.dialerFact(m.host, m.port, m.username, m.password)
	// 执行接口方法，发送邮件
	if err := dialer.DialAndSend(s); err != nil {
		return errors.Wrap(err, "try to send email got error")
	}

	return nil
}
