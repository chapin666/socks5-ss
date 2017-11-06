package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
)

const (
	// ConfigFileName .
	ConfigFileName = ".socks5-ss.json"
)

// Config struct.
type Config struct {
	// ListenAddr.
	ListenAddr string `json:"listen"`
	// RemoteAddr.
	RemoteAddr string `json:"remote"`
	// Password.
	Password string `json:"password"`
}

var configPath string

func init() {
	home, _ := homedir.Dir()
	configPath = path.Join(home, ConfigFileName)
}

// SaveConfig method
func (config *Config) SaveConfig() {
	configJSON, _ := json.MarshalIndent(config, "", "		")
	err := ioutil.WriteFile(configPath, configJSON, 0644)
	if err != nil {
		fmt.Errorf("保存到配置文件 %s 出错：%s", configPath, err)
	}
	log.Printf("保存到配置文件：%s 成功\n", configPath)
}

// ReadConfig method.
func (config *Config) ReadConfig() {
	// 如果配置文件存在，就读取配置文件中的配置 assign 到 config
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		log.Printf("从文件 %s 中读取配置\n", configPath)
		file, err := os.Open(configPath)
		if err != nil {
			log.Fatalf("打开配置文件 %s 出错：%s", configPath, err)
		}
		defer file.Close()
		err = json.NewDecoder(file).Decode(config)
		if err != nil {
			log.Fatalf("格式不合法的 JSON 配置文件：\n%s", file)
		}
	}
}
