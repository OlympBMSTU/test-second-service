package config

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// config - one file

const (
	LOG_PATH   = "logPath"
	LOG_LEVEL  = "logLevel"
	LOG_OUTPUT = "logOutput"
	TERMINATOR = "="
)

type Config struct {
	FileStorageDir    string
	ListenerHost      string
	ListenerPort      string
	DbHost            string
	DbPort            string
	Database          string
	DbUser            string
	DbPassword        string
	AuthCookieName    string
	AuthSecret        string
	TestVersion       string
	LogType           string
	LogPath           string
	LogLevel          int
	IgnoreAnswersMark string
}

func (cfg Config) GetFileStorageName() string {
	return cfg.FileStorageDir
}

func (cfg Config) GetDBHost() string {
	return cfg.DbHost
}

func (cfg Config) GetDBPort() string {
	return cfg.DbPort
}

func (cfg *Config) GetDatabase() string {
	return cfg.Database
}

func (cfg Config) GetDBUser() string {
	return cfg.DbUser
}

func (cfg Config) GetDBPassword() string {
	return cfg.DbPassword
}

func (cfg Config) IsTest() bool {
	return cfg.TestVersion == "test"
}

func (cfg Config) GetAuthCookieName() string {
	return cfg.AuthCookieName
}

func (cfg Config) GetAuthSecret() string {
	return cfg.AuthSecret
}

func (cfg Config) GetListenerHost() string {
	return cfg.ListenerHost
}

func (cfg Config) GetListenerPort() string {
	return cfg.ListenerPort
}

func (cfg Config) GetLoggerPath() string {
	return cfg.LogPath
}

func (cfg Config) GetLoggerType() string {
	return cfg.LogType
}

func (cfg Config) GetLoggerLevel() int {
	return cfg.LogLevel
}

func (cfg Config) IgnoreTestMark() bool {
	return cfg.IgnoreAnswersMark == "ignore"
}

// it works but need to get path to dir
// error handling, maybe return struct string, err
// check

func read(cfgPath string) (*Config, error) {
	file, err := os.Open(cfgPath)

	fbytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Cant start server without initial file\n"+
			"Please creaate init file and put it to "+cfgPath, err)
		return nil, err
	}

	fileData := string(fbytes)
	configs := strings.Split(fileData, "\n")

	countFields := reflect.ValueOf(Config{}).NumField()
	if len(configs) < countFields {
		log.Println("Not enough fields")
		return nil, errors.New("Not Enough fields for config")
	}

	logLevel, err := strconv.Atoi(configs[13])
	if err != nil {
		log.Print("Error logLevel, must be integer")
		return nil, err
	}

	return &Config{
		configs[0],
		configs[1],
		configs[2],
		configs[3],
		configs[4],
		configs[5],
		configs[6],
		configs[7],
		configs[8],
		configs[9],
		configs[10],
		configs[11],
		configs[12],
		logLevel,
		configs[14],
	}, nil
}

func Read(cfgPath string) (*Config, error) {
	return read(cfgPath)
}

// var configInstance *Config = nil

// func GetConfigInstance() (*Config, error) {
// 	if configInstance != nil {
// 		return configInstance, nil
// 	}

// 	var err error
// 	configInstance, err = Init()
// 	if err != nil {
// 		return nil, err
// 	}
// 	// fmt.Print(configInstance.GetFileStorageName())
// 	return configInstance, nil
// }

// logExist := false
// for configLine := range config {

// 	if strings.Contains(configLine, logLevel) {

// 		level := strings.Split()
// 	}
// }

// if
