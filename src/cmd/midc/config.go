package main

import (
	"encoding/json"
	"os"

	"github.com/midlang/mid/src/mid/build"
)

type Config struct {
	Suffix           string `cli:"suffix" usage:"source file suffix" dft".mid" name:"SUFFIX"`
	TemplatesRootDir string `cli:"template-root-dir" usage:"templates root directory" dft:"$MID_TEMP_ROOT_DIR"`

	Plugins       *build.PluginSet `json:"-" cli:"-"`
	LoadedPlugins []build.Plugin   `json:"plugins" cli:"-"`
}

func newConfig() *Config {
	cfg := new(Config)
	cfg.Plugins = build.NewPluginSet()
	return cfg
}

func (cfg *Config) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(cfg); err != nil {
		return err
	}
	for _, p := range cfg.LoadedPlugins {
		if err := cfg.Plugins.Register(&p); err != nil {
			return err
		}
	}
	return nil
}
