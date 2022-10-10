package starters

import (
	"log"
	"node_metrics_go/global"
	"node_metrics_go/infra"
	ph "node_metrics_go/internal/pusher"
	mai "node_metrics_go/internal/pusher/mail"
	"node_metrics_go/pkg/email"
	"node_metrics_go/pkg/setting"
)

type MailStarter struct {
	infra.BaseStarter
}

func (d *MailStarter) Setup(conf *setting.Config) {
	d.setupMail(conf)
}

func (d *MailStarter) setupMail(conf *setting.Config) {
	log.Println("init Mail setting ...")
	global.Mailer = conf.MailConfig
	m := email.NewMail(global.Mailer.Host, global.Mailer.Port)
	m.Login(global.Mailer.UserName, global.Mailer.Password)
	ph.PusherList.RegisterPusher(&mai.MailPusher{
		Mail: m,
	})
}
