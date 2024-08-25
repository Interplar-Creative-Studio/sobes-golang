package task2

type AppConfig struct {
	Host string
	Port string

	DB1 DbConfig
	DB2 DbConfig `yaml:"db2"`
}

type DbConfig struct {
	Host                     string
	Port                     string
	Type                     string
	Table                    string
	User                     string
	Password                 string
	ConnectionTimeoutSeconds int `yaml:"timeout"`
}

func ParseConfig(dst any, data []byte) error {
	// TODO impl
	panic("implement me")
}
