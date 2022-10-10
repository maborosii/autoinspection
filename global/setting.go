// 定义全局变量
package global

import (
	mail "node_metrics_go/pkg/email"
	rs "node_metrics_go/pkg/rules"
	"node_metrics_go/pkg/setting"

	"go.uber.org/zap"
)

var (
	MonitorSetting *setting.Config
	Logger         *zap.Logger
	ConfigPath     string
	Mailer         *mail.Mail
)

var PromQLForMap = "node_uname_info - 0"
var NotifyRules = make(map[string]rs.RuleItf, 100)
