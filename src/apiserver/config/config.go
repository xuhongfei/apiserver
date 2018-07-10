package config

import (
	"github.com/spf13/viper"
	"strings"
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 初始化日志包
	c.initLog()

	c.watchConfig()

	return nil
}

func (c *Config)initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("src/apiserver/conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APISERVER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config)watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		print("Config file changed: %s", e.Name)
	})
}

func (c *Config)initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers: viper.GetString("log.writers"),
		LoggerLevel: viper.GetString("log.logger_level"),
		LoggerFile: viper.GetString("log.logger_file"),
		LogFormatText: viper.GetBool("log.logger_format_text"),
		RollingPolicy: viper.GetString("log.rollingPolicy"),
		LogRotateDate: viper.GetInt("log.log_rotate_date"),
		LogRotateSize: viper.GetInt("log.rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}