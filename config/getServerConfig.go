package config

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	serverConfig map[string]interface{}
)

func GetServerConfig(s string) interface{} {
	if serverConfig == nil {
		LoadServerConfig()
	}
	return serverConfig[s]
}

func LoadServerConfig() error {

	var opt ini.LoadOptions
	opt.IgnoreInlineComment = true
	cfg, err := ini.LoadSources(opt, "/opt/thunder/thunder.ini")
	if err != nil {
		fmt.Println(err)
		return err
	}
	serverConfig = map[string]interface{}{
		"host":     cfg.Section("MainServer").Key("DataBaseServerIp").String(),
		"passwd":   cfg.Section("MainServer").Key("Password").String(),
		"port":     "3306",
		"username": cfg.Section("MainServer").Key("UserName").String(),
		"dbname":   "karaok",
	}

	fmt.Println(serverConfig["passwd"])
	return nil
}
