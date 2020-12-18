package config

func GetDbConfig() map[string]interface{} {

	// init db config
	dbConfig := make(map[string]interface{})

	dbConfig["hostname"] 	= "lxh001.top"
	dbConfig["port"] 		= "3306"
	dbConfig["database"] 	= "authority"
	dbConfig["username"] 	= "614"
	dbConfig["password"] 	= "661144"
	dbConfig["charset"]		= "utf8"
	dbConfig["parseTime"]	= "True"

	return dbConfig
}
