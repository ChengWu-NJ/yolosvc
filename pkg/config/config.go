package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	DarknetConfigFile  string  `yaml:"darknetConfigFile"`
	ObjNamesFile       string  `yaml:"objNamesFile"`
	DarknetWeightsFile string  `yaml:"darknetWeightsFile"`
	DetectThreshold    float32 `yaml:"detectThreshold"`
	PortOfGrpcSvc      int     `yaml:"portOfGrpcSvc"`

	ClassNames []string `yaml:"-"`
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
		ClassNames:         make([]string, 0),
	}

	if _, err := os.Stat(_configFile); err != nil {
		if os.IsNotExist(err) {
			cfgstr, err := yaml.Marshal(cfg)
			if err != nil {
				log.Fatalf("marshall config object to yaml: [%v]", err)
			}

			if err := ioutil.WriteFile(_configFile, cfgstr, 0660); err != nil {
				log.Fatal(err)
			}

			cfg.loadObjNames()
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
	validDarknetFile(cfg.ObjNamesFile)
	if cfg.DetectThreshold <= 0. {
		cfg.DetectThreshold = DEFAULT_DETECT_THRESHOLD
	}

	cfg.loadObjNames()
	return cfg
}

func validDarknetFile(filename string) {
	if _, err := os.Stat(filename); err != nil {
		log.Fatalf(`validDarknetFile: there is troubles with %s. Got err:[%+v]`, filename, err)
	}
}

func (c *Config) loadObjNames() {
	bs, err := ioutil.ReadFile(c.ObjNamesFile)
	if err != nil {
		log.Fatal(err)
	}

	ss := strings.ReplaceAll(string(bs), "\r", "")
	slc := strings.Split(ss, "\n")

	for _, obj := range slc {
		objStr := strings.TrimSpace(obj)
		if len(objStr) == 0 {
			continue
		}

		c.ClassNames = append(c.ClassNames, objStr)
	}
}
