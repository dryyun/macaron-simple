package config

import (
	ini "gopkg.in/ini.v1"
)

type ConfigFile struct {
	*ini.File
}

func (c *ConfigFile) GetConfigKey(name string) *ini.Key {
	cfg := Get("config/app.ini")
	section := cfg.Section("").Key("env").MustString("")

	return cfg.Section(section).Key(name)
}

func Get(source interface{}, others ...interface{}) *ConfigFile {
	f, _ := ini.LoadSources(ini.LoadOptions{
		Loose:            true,
		Insensitive:      true,
		AllowBooleanKeys: true,
	}, source, others...)
	return &ConfigFile{f}
}
