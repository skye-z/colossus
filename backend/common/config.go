package common

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	// 获取用户配置文件目录
	configDir, _ := os.UserConfigDir()
	configDir = fmt.Sprintf("%s/Colossus", configDir)
	// 判断应用目录是否存在
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		// 目录不存在,创建目录
		os.Mkdir(configDir, os.ModePerm)
	}
	log.Println("Config path: " + configDir + "/config.ini")
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath(configDir)
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createDefault()
		} else {
			// 配置文件被找到，但产生了另外的错误
			fmt.Println(err)
		}
	}
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
	viper.WriteConfig()
}

func GetAll() map[string]string {
	var objMap map[string]string
	objMap = make(map[string]string)
	for _, key := range viper.AllKeys() {
		if key == "security.secret" {
			continue
		}
		objMap[key] = viper.GetString(key)
	}

	return objMap
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

const (
	VersionCode   = "0.0.1"
	VersionStage  = "alpha"
	VersionNumber = "4"
)

func createDefault() {
	// 终端背景颜色
	viper.SetDefault("terminal.background_color", "rgba(0, 0, 0, 1)")
	// 终端文字颜色
	viper.SetDefault("terminal.text_color", "rgba(255, 255, 255, 1)")
	// 终端文字大小
	viper.SetDefault("terminal.text_size", "14")
	// 下载目录
	downloadDir, _ := os.UserHomeDir()
	downloadDir = fmt.Sprintf("%s/%s", downloadDir, "Downloads/colossus")
	viper.SetDefault("download.directory", downloadDir)
	// 打包下载后自动解压
	viper.SetDefault("download.auto_unzip", "false")
	// 上传自动压缩打包
	viper.SetDefault("file.show_hide", "true")
	// AES密钥
	secret, err := generateSecret()
	if err != nil {
		panic(err)
	}
	viper.SetDefault("security.secret", secret)
	// VI前缀
	viper.SetDefault("security.prefix", "betax-")
	// 写入配置
	viper.SafeWriteConfig()
}

func generateSecret() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}
