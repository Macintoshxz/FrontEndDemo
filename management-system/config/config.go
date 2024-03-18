package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 配置文件结构体
type Config struct {
	DataSource  dataSource
	Appserver   appserver
	AdminInfo   adminInfo
	RedisConfig redisConfig
}

type appserver struct {
	AppName string
	AppPort int
}

type dataSource struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type redisConfig struct {
	Addr     string
	Password string
	DB       int
}

type adminInfo struct {
	AdmInEmail string
}

var appConfig = &Config{}

func Init() error {
	getwd, _ := os.Getwd()
	fmt.Printf("配置文件所在路径: %v\n", getwd)
	// separator := string(filepath.Separator)
	viper.AddConfigPath(getwd)         // 设置配置文件路径
	viper.SetConfigName("application") // 设置配置文件名
	viper.SetConfigType("yaml")        // 设置配置文件类型格式为YAML

	// 初始化配置文件
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}
	// 监控配置文件变化并热加载程序，即不重启程序进程就可以加载最新的配置
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		initConfig()
	})
	initConfig()

	return nil
}

func initConfig() {
	appConfig = &Config{
		dataSource{
			Host:     viper.GetString("datasource.host"),
			Port:     viper.GetString("datasource.port"),
			Database: viper.GetString("datasource.database"),
			Username: viper.GetString("datasource.username"),
			Password: viper.GetString("datasource.password"),
		},
		appserver{
			AppName: viper.GetString("app.name"),
			AppPort: viper.GetInt("app.port"),
		},
		adminInfo{
			AdmInEmail: viper.GetString("admin.email"),
		},
		redisConfig{
			Addr:     viper.GetString("redis.addr"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.DB"),
		},
	}
}
func GetCofnig() *Config {
	return appConfig
}
