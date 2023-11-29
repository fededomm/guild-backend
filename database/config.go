package database

type DbInfo struct {
	Host     string `json:"host" yaml:"host" mapstructure:"host"`
	Port     string `json:"port" yaml:"port" mapstructure:"port"`
	User     string `json:"user" yaml:"user" mapstructure:"user"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	Dbname   string `json:"dbname" yaml:"dbname" mapstructure:"dbname"`
	Sslmode  string `json:"sslmode" yaml:"sslmode" mapstructure:"sslmode"`
}
