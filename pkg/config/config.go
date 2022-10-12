package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ChengWu-NJ/yolosvc/pkg/drawbbox"
	"github.com/gookit/slog"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	DarknetConfigFile  string  `yaml:"darknetConfigFile"`
	ObjNamesFile       string  `yaml:"objNamesFile"`
	DarknetWeightsFile string  `yaml:"darknetWeightsFile"`
	DetectThreshold    float32 `yaml:"detectThreshold"`
	PortOfGrpcSvc      int     `yaml:"portOfGrpcSvc"`

	ObjClasses []*ObjClass `yaml:"objClasses"`
}

type ObjClass struct {
	Id          int    `yaml:"id"`
	Name        string `yaml:"name"`
	LabelColorR int    `yaml:"labelColorR"`
	LabelColorG int    `yaml:"labelColorG"`
	LabelColorB int    `yaml:"labelColorB"`
}

const (
	DEFAULT_DARKNET_FILES_DIR    = `darknetfiles`
	DEFAULT_DARKNET_CONFIGFILE   = `nn.cfg`
	DEFAULT_DARKNET_OBJNAMESFILE = `obj.names`
	DEFAULT_DARKNET_WEIGHTSFILE  = `nn.weights`
	DEFAULT_DETECT_THRESHOLD     = 0.8
	DEFAULT_PORT_OF_GRPCSVC      = 37658
)

var (
	ExecutablePath, ExecutableFile = getExecutablePath()
	GlobalConfig                   = newConfig()
)

func (cfg *Config) GetClassNames() []string {
	slc := make([]string, 0)
	for _, cls := range cfg.ObjClasses {
		slc = append(slc, cls.Name)
	}

	return slc
}

func (cfg *Config) GetClassNumber() int {
	return len(cfg.ObjClasses)
}

func getExecutablePath() (string, string) {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Dir(ex), filepath.Base(ex)
}

func newConfig() *Config {
	_configFile := ExecutablePath + `/config.yml`
	cfg := &Config{
		DarknetConfigFile:  ExecutablePath + `/` + DEFAULT_DARKNET_FILES_DIR + `/` + DEFAULT_DARKNET_CONFIGFILE,
		ObjNamesFile:       ExecutablePath + `/` + DEFAULT_DARKNET_FILES_DIR + `/` + DEFAULT_DARKNET_OBJNAMESFILE,
		DarknetWeightsFile: ExecutablePath + `/` + DEFAULT_DARKNET_FILES_DIR + `/` + DEFAULT_DARKNET_WEIGHTSFILE,
		DetectThreshold:    DEFAULT_DETECT_THRESHOLD,
		PortOfGrpcSvc:      DEFAULT_PORT_OF_GRPCSVC,
		ObjClasses:         make([]*ObjClass, 0),
	}

	if _, err := os.Stat(_configFile); err != nil {
		if os.IsNotExist(err) {
			cfg.setObjClasses()

			cfgstr, err := yaml.Marshal(cfg)
			if err != nil {
				log.Fatalf("marshall config object to yaml: [%v]", err)
			}

			if err := ioutil.WriteFile(_configFile, cfgstr, 0660); err != nil {
				log.Fatal(err)
			}

			return cfg
		}

		log.Fatal(err)
	}

	cfgstr, err := ioutil.ReadFile(_configFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(cfgstr, cfg); err != nil {
		log.Fatal(err)
	}

	validDarknetFile(cfg.DarknetConfigFile)
	validDarknetFile(cfg.DarknetWeightsFile)
	if cfg.DetectThreshold <= 0. {
		cfg.DetectThreshold = DEFAULT_DETECT_THRESHOLD
	}

	cfg.setObjClasses()
	return cfg
}

func validDarknetFile(filename string) {
	if _, err := os.Stat(filename); err != nil {
		log.Fatalf(`validDarknetFile: there is troubles with %s. Got err:[%+v]`, filename, err)
	}
}

func (cfg *Config) setObjClasses() {
	if len(cfg.ObjClasses) > 0 {
		slog.Info(`read object classes from config.yml`)
		return
	}

	bs, err := ioutil.ReadFile(cfg.ObjNamesFile)
	if err != nil {
		log.Fatalf(`not found informations of object classes from config.yml, `+
			`then try to read from file %s, but got err[%v]`, cfg.ObjNamesFile, err)
	}

	ss := strings.ReplaceAll(string(bs), "\r", "")
	slc := strings.Split(ss, "\n")

	idx := 0
	for _, obj := range slc {
		objStr := strings.TrimSpace(obj)
		if len(objStr) == 0 {
			continue
		}

		objClass := &ObjClass{
			Id:   idx,
			Name: objStr,
		}

		cfg.ObjClasses = append(cfg.ObjClasses, objClass)
		idx += 1
	}

	// calc label colors for obj classes
	total := cfg.GetClassNumber()
	for _, cls := range cfg.ObjClasses {
		// calc color
		offset := (idx * 123457) % total
		cls.LabelColorR = int(drawbbox.GetColor(2, offset, total) * 255)
		cls.LabelColorG = int(drawbbox.GetColor(1, offset, total) * 255)
		cls.LabelColorB = int(drawbbox.GetColor(0, offset, total) * 255)
	}
}
