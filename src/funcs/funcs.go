package funcs

import (
	"gitlabparse/src/plugins"
	"text/template"
)

type Functions struct {
	Funcs template.FuncMap
}

func (f *Functions) LoadFunctionsByPlugin(plugins *plugins.Plugin) *Functions {
	funcs := template.FuncMap{}
	for _, value := range plugins.Plugins {
		for key, value := range value {
			funcs[key] = value
		}
	}

	f.Funcs = funcs
	return f
}

func (f *Functions) LoadEnvFunction(env *Env) *Functions {
	f.Funcs["getEnv"] = env.GetEnv
	return f
}

func NewFunctions() *Functions {
	return new(Functions)
}
