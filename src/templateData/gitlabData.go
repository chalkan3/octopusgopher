package templateData

import (
	"log"
	"os"
	"strings"

	"github.com/ghodss/yaml"
)

type GitlabData struct {
	Jobs      Jobs                `yaml:"jobs"`
	Image     string              `yaml:"image"`
	Services  []string            `yaml:"services"`
	Variables []map[string]string `yaml:"variables"`
	Env       map[string]string
	Plugins   map[string]map[string]func(...interface{}) (interface{}, error)
}

func (gd *GitlabData) Unmarshall(body []byte) *GitlabData {
	if err := yaml.Unmarshal(body, gd); err != nil {
		log.Fatal(err)
	}

	return gd
}

func (gd *GitlabData) ParsingEnv() *GitlabData {
	envMap := make(map[string]string)

	for _, v := range os.Environ() {
		splitV := strings.Split(v, "=")
		envMap[splitV[0]] = splitV[1]
	}
	gd.Env = envMap
	return gd
}

func NewGitlabData() *GitlabData {
	return new(GitlabData)
}
