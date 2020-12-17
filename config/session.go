package config

func GetSessionConfig() map[string]interface{}{
	sessionConfig := make(map[string]interface{})

	sessionConfig["key"] 	= "authority_management"
	sessionConfig["name"]	= "authority_session"

	return sessionConfig
}