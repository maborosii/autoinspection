package alert

import (
	"fmt"
	"node_metrics_go/global"

	"github.com/jedib0t/go-pretty/v6/table"
	"go.uber.org/zap"
)

type AlertInfo interface {
	PrintAlert(mType string) string
	PrintAlertFormatTable(mType string) table.Row
}
type baseAlertMessage struct {
	job, instance, alertMessage, alertMetricsLimit, alertMetricsUsage string
}

func NewBaseAlertMessage(job, instance, alertMessage, alertMetricsLimit, alertMetricsUsage string) *baseAlertMessage {
	return &baseAlertMessage{
		job:               job,
		instance:          instance,
		alertMessage:      alertMessage,
		alertMetricsLimit: alertMetricsLimit,
		alertMetricsUsage: alertMetricsUsage,
	}
}

func (n *baseAlertMessage) PrintAlert(mType string) string {
	return fmt.Sprintf("%s 指标异常 >>> job: %s, instance: %s,  告警信息:%s, 当前值:%s, 预警值：%.s\n", mType, n.job, n.instance, n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit)
}
func (n *baseAlertMessage) PrintAlertFormatTable(mType string) table.Row {
	return table.Row{fmt.Sprintf("%s 指标异常", mType), n.job, n.instance, "", n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit}
}

/* 合并告警信息
 */
// 文本信息
func MergeAlertInfo(a <-chan AlertInfo) string {
	var mm string
	alertInfoByKind := make(map[string]string, 5)
	for v := range a {
		switch v.(type) {
		case *nodeAlertMessage:
			alertInfoByKind["node"] += v.PrintAlert("Node")
		case *redisAlertMessage:
			alertInfoByKind["redis"] += v.PrintAlert("Redis")
		case *kafkaAlertMessage:
			alertInfoByKind["kafka"] += v.PrintAlert("Kafka")
		default:
			global.Logger.Warn("alert info not found suitable type", zap.String("info", v.PrintAlert("Unknown")))
		}
	}
	for _, infos := range alertInfoByKind {
		mm += infos
	}
	return mm
}

// 表格信息
func MergeAlertInfoFormatTable(a <-chan AlertInfo) []table.Row {
	var mm []table.Row
	alertInfoByKind := make(map[string][]table.Row, 5)
	for v := range a {
		switch v.(type) {
		case *nodeAlertMessage:
			alertInfoByKind["node"] = append(alertInfoByKind["node"], v.PrintAlertFormatTable("Node "))
		case *redisAlertMessage:
			alertInfoByKind["redis"] = append(alertInfoByKind["redis"], v.PrintAlertFormatTable("Redis"))
		case *kafkaAlertMessage:
			alertInfoByKind["kafka"] = append(alertInfoByKind["kafka"], v.PrintAlertFormatTable("Kafka"))
		default:
			global.Logger.Warn("alert info not found suitable type", zap.String("info", v.PrintAlert("Unknown")))
		}
	}
	for _, infos := range alertInfoByKind {
		mm = append(mm, infos...)
	}
	return mm
}

// render table to html
func RenderTable(rows []table.Row) string {
	t := table.NewWriter()
	t.AppendHeader(tableHeader)
	t.AppendRows(rows)
	t.SetAutoIndex(true)
	// 根据指标类型和 job 名称进行排序
	t.SortBy([]table.SortBy{sortedByKind, sortedByJob})

	t.Style().HTML = table.HTMLOptions{
		CSSClass:    "",
		EmptyColumn: "&nbsp;",
		EscapeText:  true,
		Newline:     "<br/>",
	}

	prefixMailHtml := fmt.Sprintf("<style>\n%s\n</style>\n", styleCss)
	htmlContext := prefixMailHtml + t.RenderHTML()
	return htmlContext
}
