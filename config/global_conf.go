package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var Config = GlobalConfig{}

func GetConfig() *GlobalConfig {
	return &Config
}

type (
	GlobalConfig struct {
		LineServer struct {
			ChannelSecret string `toml:"channel_secret"`
			ChannelToken  string `toml:"channel_token"`
		} `toml:"line_server"`
		GptConf struct {
			AuthToken string `toml:"auth_token"`
		} `toml:"gpt"`
		Server struct {
			Host string
			Port int64
		}
	}
)

func LoadGlobalConfig() (*GlobalConfig, error) {
	filePath := "./config.toml"
	_, err := toml.DecodeFile(filePath, &Config)
	if err != nil {
		return nil, fmt.Errorf("load Config file '%s' failed, %s", filePath, err)
	}

	//marshal, err := json.Marshal(Config)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(string(marshal))

	return &Config, nil
}
