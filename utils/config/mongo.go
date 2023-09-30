package config

type MongoDB struct {
	Host   string `yaml:"host"`
	Dbname string `yaml:"dbname"`
}
