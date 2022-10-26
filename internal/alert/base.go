package alert

import (
	"fmt"
	"node_metrics_go/global"

	"github.com/jedib0t/go-pretty/v6/table"
	"go.uber.org/zap"
)

type AlertInfo interface {
	PrintAlert() string
	PrintAlertFormatTable() table.Row
}

/* 合并告警信息
 */
// 文本信息
func MergeAlertInfo(a <-chan AlertInfo) string {
	var mm string
	alertInfoByKind := make(map[string]string, 5)
	for v := range a {
		switch v.(type) {
		case *NodeAlertMessage:
			alertInfoByKind["node"] += v.PrintAlert()
		case *RedisAlertMessage:
			alertInfoByKind["redis"] += v.PrintAlert()
		default:
			global.Logger.Warn("alert info not found suitable type", zap.String("info", v.PrintAlert()))
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
		case *NodeAlertMessage:
			alertInfoByKind["node"] = append(alertInfoByKind["node"], v.PrintAlertFormatTable())
		case *RedisAlertMessage:
			alertInfoByKind["redis"] = append(alertInfoByKind["redis"], v.PrintAlertFormatTable())
		default:
			global.Logger.Warn("alert info not found suitable type", zap.String("info", v.PrintAlert()))
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
