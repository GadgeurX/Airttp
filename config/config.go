package config

import (
	"github.com/olebedev/config"
	"io/ioutil"
	"Airttp/logger"
)

var m_Config *Config = nil

type Config struct {
	mainConfig *config.Config
	modulesConfig *config.Config
}

func NewConfig() *Config {
	configObj := new(Config)

	json, err := ioutil.ReadFile("./config.json")
	if err != nil {
		logger.GetInstance().Error(err.Error())
		return nil
	} else {
		var err error
		configObj.mainConfig, err = config.ParseJson(string(json))
		if err != nil {
			logger.GetInstance().Error("config.json : " + err.Error())
			return nil
		}
		modulesPath, err := configObj.mainConfig.String("modules_conf_path")
		if (err != nil) {
			modulesPath = "./modules.json"
		}
		json, err := ioutil.ReadFile(modulesPath)
		if err != nil {
			logger.GetInstance().Error(err.Error())
		} else {
			configObj.modulesConfig, err = config.ParseJson(string(json))
			if err != nil {
				logger.GetInstance().Error(modulesPath + " : " + err.Error())
			}
		}
	}
	return configObj
}


func GetConfigInstance() *Config {
	if (m_Config != nil) {
		return m_Config
	} else {
		m_Config = NewConfig()
		return m_Config
	}
}

func (conf *Config) GetServerPort(def int) int {
	port, err := conf.mainConfig.Int("server.port")
	if (err != nil) {
		return def
	}
	return port
}

func (conf *Config) GetModulesList() map[string]map[string]string {
	modulesListInterface, err := conf.modulesConfig.Map("modules")
	if (err != nil) {
		return nil
	}
	modulesList := make(map[string]map[string]string)
	for key, value := range modulesListInterface {
		modulesList[key] = make(map[string]string)
		tmpList := value.(map[string]interface{})
		for tmpkey, tmpvalue := range tmpList {
			modulesList[key][tmpkey] = tmpvalue.(string)
		}
	}
	return (modulesList)
}