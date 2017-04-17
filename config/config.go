package config

import ini "gopkg.in/ini.v1"

func Get(source interface{}, others ...interface{}) *ini.File {
	f, _ := ini.LoadSources(ini.LoadOptions{
		Loose:            true,
		Insensitive:      true,
		AllowBooleanKeys: true,
	}, source, others...)
	return f
}
