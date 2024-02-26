package config

import "fmt"

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
