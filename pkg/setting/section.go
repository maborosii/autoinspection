package setting

type MonitorConfig struct {
	Address      string        `toml:"address"`
	TimeOut      int           `toml:"timeout"`
	MonitorItems MonitorItems  `toml:"monitorItems"`
	Thresholds   []*NotifyRule `toml:"thresholds"`
	LogConfig    *LogConf      `toml:"logconfig"`
}
type NotifyRule struct {
	Job                  string `toml:"job"`
	CpuThreshold         string `toml:"cpu"`
	CpuIncreaseThreshold string `toml:"cpu_increase"`
	MemThreshold         string `toml:"mem"`
	MemIncreaseThreshold string `toml:"mem_increase"`
	// DiskThreshold        string `toml:"disk"`
}

type LogConf struct {
	Level      string `toml:"level"`
	LogFile    string `toml:"logfile"`
	MaxSize    int    `toml:"maxsize"`
	MaxAge     int    `toml:"maxage"`
	MaxBackups int    `toml:"maxbackups"`
}

func (conf *MonitorConfig) GetTimeOut() int {
	return conf.TimeOut
}
func (conf *MonitorConfig) GetAddress() string {
	return conf.Address
}
func (conf *MonitorConfig) GetMonitorItems() map[string]string {
	return conf.MonitorItems.ConvertToMap()
}

func (conf *MonitorConfig) GetLogConfig() *LogConf {
	return conf.LogConfig
}

func NewMonitorConfig() *MonitorConfig {
	return &MonitorConfig{}
}

type MonitorItems []*MonitorItem

type MonitorItem struct {
	Metrics string `toml:"metrics"`
	PromQL  string `toml:"promql"`
}

func (i MonitorItems) ConvertToMap() map[string]string {
	promQLs := make(map[string]string)
	for _, item := range i {
		promQLs[item.Metrics] = item.PromQL
	}
	return promQLs
}

// 将配置文件数据映射到结构体中
func (s *Setting) ReadConfig(value interface{}) error {
	err := s.vp.Unmarshal(value)
	if err != nil {
		return err
	}
	return nil
}
