package configs

import "os"

//DatabaseConfig represent one database config
type DatabaseConfig struct {
	Username string
	Password string
	DBName   string
}

func SetupDatabase() DatabaseConfig {
	dbConfig := DatabaseConfig{}

	switch os.Getenv("DB") {
	case "PGQL":
		dbConfig.Username = "faruqisan"
		dbConfig.Password = "faruqisan"
		dbConfig.DBName = "faruqisan"
		break
	case "MY":
		dbConfig.Username = "root"
		dbConfig.Password = ""
		dbConfig.DBName = "modellingkol"
		break;
	default:
		dbConfig.Username = "faruqisan"
		dbConfig.Password = "faruqisan"
		dbConfig.DBName = "faruqisan"
		break
	}

	return dbConfig
}
