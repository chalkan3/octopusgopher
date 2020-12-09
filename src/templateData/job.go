package templateData

type Jobs []Job
type Job struct {
	Name         string              `yaml:"name"`
	Stage        string              `yaml:"stage"`
	Image        string              `yaml:"image"`
	Variable     []map[string]string `yaml:"variable"`
	Only         []string            `yaml:"only"`
	BeforeScript []string            `yaml:"before_script"`
	Script       []string            `yaml:"script"`
}

func NewJob() *Job {
	return new(Job)
}
