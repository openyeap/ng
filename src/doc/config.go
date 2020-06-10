package doc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"gopkg.in/yaml.v2"
)

type Toc struct {
	Name    string
	Host    string
	Title   string
	Date    string
	Authors []string
	Summary string
	Root    map[string]Doc
	Pages   map[string]Page
	Datas   map[string]map[string]string
}

type Page struct {
	Name       string
	Path       string
	Title      string
	Date       string
	Summary    string
	Order      string
	Extension  string
	Authors    []string
	Keywords   []string
	Tags       []string
	Categories []string
}
type Doc struct {
	Page
	Data  []byte
	Items map[string]Doc
}

func WalkDir(dir string) {
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		doc := Doc{}
		doc.Name = f.Name()
		doc.Path = path
		if f.IsDir() {
			WalkDir(path)
			path += ".setting"
		}

		//load raw data
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		doc.Data = data

		//load settings
		dict := make(map[string]interface{})
		err = yaml.Unmarshal([]byte(data), &dict)
		if err != nil {
			return err
		}
		//generate html
		// html := string(markdown.Run([]byte(data)))
		// fmt.Println(html)
		// engine := liquid.NewEngine()

		// out, err := engine.ParseAndRenderString(html, siteMap)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// fmt.Println(out)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

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

	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// 监听配置文件是否改变,用于热更新
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})
}

func getHandles(path string) []Config {

	var handles []Config
	file, err := os.Open(path)
	if err != nil {
		return handles
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return handles
	}
	json.Unmarshal([]byte(data), &handles)
	return handles
}
