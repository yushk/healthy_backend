package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// viper keys
const (
	envPrefix = "healthy"
	EnvKey    = "env"

	DNSServer    = "dns_server"
	PodName      = "pod.name"
	PodNamespace = "pod.namespace"

	// oauth2
	ClientID      = "oauth2.client.id"
	ClientSecret  = "oauth2.client.secret"
	RedirectURL   = "oauth2.redirect.url"
	AuthServerURL = "oauth2.server.url"

	ManagerProtocol = "manager.protocol"
	ManagerHost     = "manager.host"
	ManagerPort     = "manager.port"

	DBType     = "db.type" // json;mongodb;mysql;sqlite3
	DBAddress  = "db.address"
	DBUsername = "db.username"
	DBPassword = "db.password"

	// MQTT协议
	MqttProtocol = "mqtt.protocol"
	MqttHost     = "mqtt.host"
	MqttPort     = "mqtt.port"

	RootUserID   = "root.userid"
	RootUsername = "root.username"
	RootPassword = "root.password"

	// 产品ID
	ProductID = "product.Id"
)

var viperInitOnce sync.Once
var config *viper.Viper

func init() {
	viperInitOnce.Do(func() {
		config = viper.New()
		config.SetConfigType("yaml")

		config.SetEnvPrefix(envPrefix) // will be uppercased automatically
		config.AutomaticEnv()
		config.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

		config.SetDefault(EnvKey, "prod")
		config.SetDefault(DNSServer, "local")
		config.SetDefault(PodName, "manager")
		config.SetDefault(PodNamespace, "healthy")

		config.SetDefault(ClientID, "11111")
		config.SetDefault(ClientSecret, "22222")
		config.SetDefault(RedirectURL, "http://localhost:8080/v1/oauth/registered/callback")
		config.SetDefault(AuthServerURL, "http://localhost:8085")

		config.SetDefault(DBType, "mongodb")
		config.SetDefault(DBAddress, "mongodb://localhost:27017")
		config.SetDefault(DBUsername, "")
		config.SetDefault(DBPassword, "")

		config.SetDefault(ManagerProtocol, "tcp")
		config.SetDefault(ManagerHost, "localhost")
		config.SetDefault(ManagerPort, "30031")

		config.SetDefault(MqttProtocol, "tcp")
		config.SetDefault(MqttHost, "192.168.1.118")
		config.SetDefault(MqttPort, "31883")

		config.SetDefault(RootUserID, "")
		config.SetDefault(RootUsername, "kaisawind")
		config.SetDefault(RootPassword, "kaisawind")

		config.SetDefault(ProductID, "c7agh5h0vo40rfaddtdg")

		configFile := fmt.Sprintf("config-%s", config.Get(EnvKey))
		logrus.Debugf("config file is %s", configFile+".yaml")
		config.SetConfigName(configFile) // name of config file (without extension)
		config.AddConfigPath(fmt.Sprintf("/%s/", envPrefix))
		config.AddConfigPath(fmt.Sprintf("/etc/%s/", envPrefix))
		config.AddConfigPath(fmt.Sprintf("$HOME/.%s", envPrefix))
		config.AddConfigPath(".")
		config.AddConfigPath("./bin")

		err := config.ReadInConfig()
		if err != nil {
			switch err := err.(type) {
			case viper.ConfigFileNotFoundError:
				logrus.Debugf("No config file '%s' found. Using environment variables only.", configFile)
			case viper.ConfigParseError:
				logrus.Panicf("Cannot read config file: %s: %s", configFile, err)
			default:
				logrus.Debugf("Read config file error: %s: %s", configFile, err)
			}
		} else {
			logrus.Infof("Loading config from file %s", config.ConfigFileUsed())
		}
	})
}

// Set  sets the value for the key in the override register.
func Set(key string, value interface{}) {
	config.Set(key, value)
}

// GetString returns the value associated with the key as a string.
func GetString(key string) string {
	return config.GetString(key)
}

// GetInt returns the value associated with the key as a int.
func GetInt(key string) int {
	return config.GetInt(key)
}
