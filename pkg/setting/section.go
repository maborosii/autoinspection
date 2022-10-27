package setting

type MetricType string
type Config struct {
	Endpoints map[string]DataSource `toml:"endpoints" mapstructure:"endpoints"`
	// Endpoints    map[string]string        `toml:"endpoints" mapstructure:"endpoints"`
	TimeOut      int                      `toml:"timeout"`
	MonitorItems MonitorItems             `toml:"monitoritems"`
	Rules        map[string][]interface{} `toml:"rules" mapstructure:"rules"`
	LogConfig    *LogConf                 `toml:"logconfig"`
	MailConfig   *MailConf                `toml:"mailconfig"`
}
type DataSource struct {
	Type    string `toml:"type"`
	Address string `toml:"address"`
}
type MailConf struct {
	Host     string   `toml:"host"`
	UserName string   `toml:"username"`
	Password string   `toml:"password"`
	Port     int      `toml:"port"`
	Subject  string   `toml:"subject"`
	To       []string `toml:"to"`
}
type LogConf struct {
	Level      string `toml:"level"`
	LogFile    string `toml:"logfile"`
	MaxSize    int    `toml:"maxsize"`
	MaxAge     int    `toml:"maxage"`
	MaxBackups int    `toml:"maxbackups"`
}

func (conf *Config) GetTimeOut() int {
	return conf.TimeOut
}

// func (conf *Config) GetMonitorItems() map[string]string {
// 	return conf.MonitorItems.ConvertToMap()
// }
func (conf *Config) GetLogConfig() *LogConf {
	return conf.LogConfig
}

func NewConfig() *Config {
	return &Config{}
}

type MonitorItems []*MonitorItem

func (m MonitorItems) Filter(mType string) MonitorItems {
	var newM MonitorItems
	if mType != "" {
		for _, i := range m {
			if i.MType == mType {
				newM = append(newM, i)
			}
		}
		return newM
	}
	return m
}

func (m MonitorItems) FindAdaptEndpoints(metricType string) map[string]struct{} {
	endPointSet := make(map[string]struct{}, 10)
	for _, item := range m.Filter(metricType) {
		for _, ed := range item.Endpoint {
			endPointSet[ed] = struct{}{}
		}
	}
	return endPointSet
}

type MonitorItem struct {
	MType   string `toml:"mtype"`
	Metrics string `toml:"metrics"`
	PromQL  string `toml:"promql"`
	// 从哪个数据源查询数据
	Endpoint []string `toml:"endpoint"`
}

// 将配置文件数据映射到结构体中
func (s *Setting) ReadConfig(value interface{}) error {
	err := s.vp.Unmarshal(value)
	if err != nil {
		return err
	}
	return nil
}
