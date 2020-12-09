package customtemplate

import (
	"gitlabparse/src/files"
	"log"
	"path"
	"runtime"
	"text/template"
)

type GitLabTemplate struct {
	template *template.Template
}

func (gotpl *GitLabTemplate) CreateTemplate(funcs template.FuncMap) *GitLabTemplate {
	_, filename, _, _ := runtime.Caller(1)
	gotpl.template, _ = template.New("gitlabci.yml").Funcs(funcs).ParseFiles(path.Join(path.Dir(filename), "template/gitlabci.yml"))
	return gotpl
}

func (gotpl *GitLabTemplate) Execute(data interface{}) {

	file := files.NewFile().CreateFile("./output/.gitlab-ci.yml")
	if err := gotpl.template.ExecuteTemplate(file, "gitlabci.yml", data); err != nil {
		log.Fatal(err)

	}

}
func NewTemplate() *GitLabTemplate {
	return new(GitLabTemplate)
}
