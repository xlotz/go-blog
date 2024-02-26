package config

type Mysql struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	DB           string `yaml:"db"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	LogLevel     string `yaml:"log_level"`
	Config       string `yaml:"config"`
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns"`
	MaxIdealTime int    `json:"max_ideal_time" yaml:"max_ideal_time"`
}

func (m Mysql) DSN() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/?" + m.Config
}

func (m Mysql) GetLogLevel() string {
	return m.LogLevel
}
