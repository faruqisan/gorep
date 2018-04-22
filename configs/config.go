package configs

type Config struct {
	DBConf DatabaseConfig
}

func InitConfig() Config {

	return Config{
		DBConf: SetupDatabase(),
	}
}
