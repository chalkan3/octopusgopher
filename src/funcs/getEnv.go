package funcs

type Env struct {
	Envs map[string]string
}

func (e *Env) GetEnv(k string) string { return e.Envs[k] }

func NewEnv(envs map[string]string) *Env {
	return &Env{
		Envs: envs,
	}
}
