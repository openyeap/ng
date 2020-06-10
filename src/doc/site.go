package doc

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Site struct {
	// Name    string
	// Host    string
	// Title   string
	// Date    string
	// Authors []string
	// Summary string
	Root    Doc
	Pages   map[string]Page
	SiteMap map[string]string
	PageMap map[string]string
}

func loadLayoutContent(dir string, path string) string {
	//load raw data
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	html := string(data)
	//if using layout load file
	s := loadLayoutContent(dir, "new path")
	return strings.Replace(s, "{{content}}", html, 0)
}
func (this *Site) LoadLayout() {
	dir := this.SiteMap["site.layout"]
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {

			return err
		}
		if f.IsDir() {
			return nil
		}
		html := loadLayoutContent(dir, path)
		this.SiteMap["site.layout."+f.Name()] = html
		return nil
	})
	if err != nil {
		fmt.Println("Load Layout Error:", err)
	}
}
func (this *Site) WalkDoc(dir string) {
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		doc := Doc{}
		doc.Name = f.Name()
		doc.Path = path
		if f.IsDir() {
			this.WalkDoc(path)
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

func (this *Site) Public(dir string) {
	// for _, page := range this.Pages {

	// }
}
func NewSite(cfg string) (Site, error) {
	site := Site{SiteMap: make(map[string]string), PageMap: make(map[string]string), Pages: make(map[string]Page), Root: Doc{}}
	//load site settings
	viper.SetConfigFile(cfg)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if nil != err {
		log.Fatalln("ERROR:", err.Error())
		return site, err
	}

	for _, key := range viper.AllKeys() {
		site.SiteMap["site."+key] = viper.GetString(key)
	}
	return site, nil
}

func (this *Site) init() {

	this.WalkDoc(this.SiteMap["site.docs"])
	this.Public(this.SiteMap["site.public"])
}
