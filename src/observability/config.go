package observability

type Observability struct {
	Enable      bool   `json:"enable" yaml:"enable" mapstructure:"enable"`
	ServiceName string `json:"serviceName" yaml:"serviceName" mapstructure:"serviceName"`
	Endpoint	string `json:"endpoint" yaml:"endpoint" mapstructure:"endpoint"`
}