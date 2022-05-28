package config

import (
	"os"
	"strconv"
)

/*************************************************************
*                         STRUCTS                            *
**************************************************************/
type Config struct {
	Host                    string
	Port                    string
	Type                    string
	MaxConnections          int
	MaxMessagesHandlers     int
	MaxOccourrencesHandlers int
	MaxStatsHandlers        int
}

/*************************************************************
*                         CONSTANTS                          *
**************************************************************/

const (
	CONNECTION_HOST          = "localhost" //HOST NAME
	CONNECTION_PORT          = "9000"      //PORT NUMBER
	CONNECTION_TYPE          = "tcp"       //CONNECTION TYPE
	MAX_CONNECTIONS          = 1000        //MAX CLIENT CONNECTIONS
	MAX_MESSAGES_HANDLERS    = 100         //MAX MESSAGES HANDLERS PER CLIENT
	MAX_OCCURRENCES_HANDLERS = 10          //MAX OCCURRENCES HANDLERS PER CLIENT
	MAX_STATS_HANDLERS       = 10          //MAX STATS HANDLERS PER CLIENT
)

/*************************************************************
*                        FUNCTIONS                            *
**************************************************************/
func LoadConfig() (config Config, err error) {
	config.Host = getEnv("HOST", CONNECTION_HOST)
	config.Port = getEnv("PORT", CONNECTION_PORT)
	config.Type = getEnv("TYPE", CONNECTION_TYPE)
	config.MaxConnections = getEnvInt("MAX_CONNECTIONS", MAX_CONNECTIONS)
	config.MaxMessagesHandlers = getEnvInt("MAX_MESSAGES_HANDLERS", MAX_MESSAGES_HANDLERS)
	config.MaxOccourrencesHandlers = getEnvInt("MAX_OCCURRENCES_HANDLERS", MAX_OCCURRENCES_HANDLERS)
	config.MaxStatsHandlers = getEnvInt("MAX_STATS_HANDLERS", MAX_STATS_HANDLERS)

	return
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i

}
