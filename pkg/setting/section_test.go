package setting

import (
	"testing"
)

func TestMonitorConfig_GetTimeOut(t *testing.T) {
	type fields struct {
		Address   string
		TimeOut   int
		Items     MonitorItems
		Rules     []*NotifyRule
		LogConfig *LogConf
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &MonitorConfig{
				Address:      tt.fields.Address,
				TimeOut:      tt.fields.TimeOut,
				MonitorItems: tt.fields.Items,
				Thresholds:   tt.fields.Rules,
				LogConfig:    tt.fields.LogConfig,
			}
			if got := conf.GetTimeOut(); got != tt.want {
				t.Errorf("MonitorConfig.GetTimeOut() = %v, want %v", got, tt.want)
			}
		})
	}
}
