package main

import (
	customtemplate "gitlabparse/src/customTemplate"
	"gitlabparse/src/files"
	"gitlabparse/src/funcs"
	"gitlabparse/src/gitlab"
	"gitlabparse/src/plugins"
	"gitlabparse/src/templateData"
)

// Config consts
const filePath string = "output/.gitlab-ci.yml"

func main() {
	config := plugins.NewPluginConfig().Load()
	values := templateData.NewGitlabData().Unmarshall(files.NewYamlFile().GetFile("values.yml")).ParsingEnv()
	plugins := plugins.NewPlugin().Load(config)
	env := funcs.NewEnv(values.Env)

	funcs := funcs.NewFunctions().LoadFunctionsByPlugin(plugins).LoadEnvFunction(env)

	customtemplate.NewTemplate().CreateTemplate(funcs.Funcs).Execute(values)
	// gitlab
	file := files.NewYamlFile().Raw(filePath)
	fileName := ".gitlab-ci.yml"
	gitlab.NewGitlab(config.Gitlab.APIKey).StartClient(config.Gitlab.ProjectID, config.Gitlab.APIPath).Commit(&fileName, &file, config.Gitlab.Commit, config.Gitlab.BranchTarget)

}
