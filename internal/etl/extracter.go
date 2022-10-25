package etl

import (
	"context"
	"regexp"
	"time"

	"node_metrics_go/global"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"go.uber.org/zap"
)

type QueryResult struct {
	label string
	value model.Value
}

func NewQueryResult() func(label string, value model.Value) *QueryResult {
	return func(label string, value model.Value) *QueryResult {
		return &QueryResult{label: label, value: value}
	}
}

func (q *QueryResult) GetLabel() string {
	return q.label
}
func (q *QueryResult) GetValue() model.Value {
	return q.value
}

// pattern: 正则表达式
// 抽取instance，label，metrics
func (q *QueryResult) CleanValue(pattern string) [][]string {
	var midResult = [][]string{}
	var re = regexp.MustCompile(pattern)
	matched := re.FindAllStringSubmatch(q.value.String(), -1)
	for _, match := range matched {
		midResult = append(midResult, []string{match[1], match[2]})
	}
	return midResult
}

func SendQueryResultToChan(label string, promql string, api v1.API, mChan chan<- *QueryResult) {
	defer func() {
		<-concurrencyChan
	}()
	concurrencyChan <- struct{}{}
	mChan <- QueryFromProm(label, promql, api)
	global.Logger.Info("metics gotten", zap.String("label", label))
}

func QueryFromProm(label string, promql string, api v1.API) *QueryResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(global.MonitorSetting.GetTimeOut())*time.Second)
	defer cancel()
	result, warnings, err := api.Query(ctx, promql, time.Now())

	if err != nil {
		global.Logger.Fatal("Error querying Prometheus: ", zap.Error(err))
		return nil
	}
	if len(warnings) > 0 {
		global.Logger.Warn("warning ", zap.Any("warnings: ", warnings))
	}
	return NewQueryResult()(label, result)
}
