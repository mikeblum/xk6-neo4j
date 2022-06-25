package neo4j

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	NEO4J_ADDR = "NEO4J_ADDR"
	NEO4J_USERNAME = "NEO4J_USERNAME"
	NEO4J_PASSWORD = "NEO4J_PASSWORD"
	NEO4J_DATABASE = "NEO4J_DATABASE"
	NEO4J_DEFAULT_ADDR = "neo4j://localhost:7687"
	NEO4J_DEFAULT_USERNAME = "neo4j"
	// note this is here for consistancy but 
	// by default neo4j will reject all cmds using the default password
	NEO4J_DEFAULT_PASSWORD = "neo4j"
	NEO4J_DEFAULT_DATABASE = "neo4j"
)

type Conf struct {
	*viper.Viper
}

func NewConf() *Conf {
	conf := viper.New()
	conf.AddConfigPath(".")
	conf.AddConfigPath("..")
	conf.AddConfigPath("../..")
	conf.SetConfigName(".env")
	conf.SetConfigType("env")
	conf.AutomaticEnv()
	err := conf.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	return &Conf{
		conf,
	}
}
