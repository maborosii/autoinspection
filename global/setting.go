// 定义全局变量
package global

import (
	rs "node_metrics_go/internal/rules"
	"node_metrics_go/pkg/setting"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"go.uber.org/zap"
)

var (
	MonitorSetting *setting.Config
	Logger         *zap.Logger
	ConfigPath     string
	MetricsType    string
	Mailer         *setting.MailConf
)

// var PromQLForKafkaInfo = "redis_instance_info - 0"
// var PromQLForEsInfo = "redis_instance_info - 0"

var NotifyRules = make(map[string]rs.RuleItf, 100)
var PromClients = make(map[string]v1.API, 10)
