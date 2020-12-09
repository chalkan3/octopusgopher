package files

import (
	"io/ioutil"
	"path"
	"runtime"
)

type YamlFile struct{}

func (*YamlFile) GetFile(file string) []byte {
	_, filename, _, _ := runtime.Caller(1)

	yaml, err := ioutil.ReadFile(path.Join(path.Dir(filename), "template/"+file))
	if err != nil {
		panic(err)
	}
	return yaml
}

func (*YamlFile) GetFilePath(paths string) []byte {
	_, filename, _, _ := runtime.Caller(1)

	yaml, err := ioutil.ReadFile(path.Join(path.Dir(filename), paths))
	if err != nil {
		panic(err)
	}

	return yaml
}

func (*YamlFile) Raw(paths string) string {
	_, filename, _, _ := runtime.Caller(1)

	file, err := ioutil.ReadFile(path.Join(path.Dir(filename), paths))
	if err != nil {
		panic(err)
	}
	return string(file)
}
func NewYamlFile() *YamlFile {
	return new(YamlFile)
}
