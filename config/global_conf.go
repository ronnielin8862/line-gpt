package config

import (
	"github.com/BurntSushi/toml"
	"log"
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

func LoadGlobalConfig() {
	filePath := "./config.toml"
	_, err := toml.DecodeFile(filePath, &Config)
	if err != nil {
		log.Fatal("LoadGlobalConfig error: ", err)
	}

	//marshal, err := json.Marshal(Config)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(string(marshal))
}
