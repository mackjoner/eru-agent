package defines

type DockerConfig struct {
	Endpoint string
	Ca       string
	Key      string
	Cert     string
	Health   int
}

type EruConfig struct {
	Endpoint string
}

type LenzConfig struct {
	Routes   string
	Forwards []string
	Stdout   bool
}

type MetricsConfig struct {
	Root      string
	Step      int64
	Timeout   int64
	Transfers []string
}

type RedisConfig struct {
	Host string
	Port int
	Min  int
	Max  int
}

type VLanConfig struct {
	Physical []string
}

type AgentConfig struct {
	HostName string `yaml:"hostname"`
	PidFile  string

	Docker  DockerConfig
	Eru     EruConfig
	Lenz    LenzConfig
	Metrics MetricsConfig
	VLan    VLanConfig
	Redis   RedisConfig
}
