package plugins

import (
	"fmt"
	"path"
	ps "plugin"
	"runtime"
)

type Plugin struct {
	Plugins map[string]map[string]func(...interface{}) (interface{}, error)
}

func (pl *Plugin) Load(config *PluginConfig) *Plugin {
	_, filename, _, _ := runtime.Caller(1)

	for _, plugin := range config.Plugins {
		p, err := ps.Open(path.Join(path.Dir(filename), "plugins/"+plugin.Name+".so"))
		if err != nil {
			fmt.Println(err)
		}

		plugins := make(map[string]func(...interface{}) (interface{}, error))
		for _, methods := range plugin.Methods {
			s, err := p.Lookup(methods)
			if err != nil {
				fmt.Println(err)
			}
			plugins[methods] = s.(func(...interface{}) (interface{}, error))
		}

		pl.Plugins[plugin.Name] = plugins
	}

	return pl
}
func NewPlugin() *Plugin {
	return &Plugin{
		Plugins: make(map[string]map[string]func(...interface{}) (interface{}, error)),
	}
}
