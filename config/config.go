package config

import (
	"io/ioutil"
	`os`
	`path/filepath`
	
	`github.com/sirupsen/logrus`
	"gopkg.in/yaml.v2"
)

func InitConf(){
	confPath := os.Getenv("CONFIG_PATH")
	pwd , _ := os.Getwd()
	defaultConfPath  := filepath.Join(pwd, "config.yaml")
	if confPath == ""{
		confPath = defaultConfPath
	}
	logrus.Infof("confPath = %s",confPath)
	Init(confPath)
}



var Conf Config

type Config struct {
	HostPort       string         `yaml:"hostport"`
	DataPath       string         `yaml:"datapath"`
	DataBaseConfig DataBaseConfig `yaml:"database"`
}

type DataBaseConfig struct {
	HostPort     string `yaml:"hostport"`
	UserPassword string `yaml:"userpassword"`
	DB           string `yaml:"db"`
}

func Init(path string) error {
	byte, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(byte, &Conf)
	if err != nil {
		panic(err)
	}
	return err
}
