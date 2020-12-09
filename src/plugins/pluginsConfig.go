package plugins

import (
	"gitlabparse/src/files"
	"log"

	"gopkg.in/yaml.v2"
)

type PluginConfig struct {
	Plugins []struct {
		Name    string   `yaml:"name"`
		Methods []string `yaml:"methods"`
	} `yaml:"plugins"`
	Gitlab struct {
		ProjectID    int    `yaml:"project_id"`
		BranchTarget string `yaml:"branch_target"`
		Commit       string `yaml:"commit"`
		APIKey       string `yaml:"api_key"`
		APIPath      string `yaml:"api_path"`
	} `yaml:"gitlab"`
}

func (pc *PluginConfig) Load() *PluginConfig {
	if err := yaml.Unmarshal(files.NewYamlFile().GetFilePath("../../config/config.yaml"), pc); err != nil {
		log.Fatal(err)
	}
	return pc
}
func NewPluginConfig() *PluginConfig {
	return new(PluginConfig)
}
