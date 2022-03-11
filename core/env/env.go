package env

type Env struct {
	name string
}

func (t *Env) Value() string {
	return t.name
}

func (t *Env) Type() string {
	for _, e := range emptyEnvs {
		if t.name == e {
			return empty
		}
	}

	for _, e := range devEnvs {
		if t.name == e {
			return development
		}
	}

	for _, e := range prodEnvs {
		if t.name == e {
			return production
		}
	}

	for _, e := range testEnvs {
		if t.name == e {
			return testing
		}
	}

	return unknown
}

func (t *Env) IsEmpty() bool {
	return t.Type() == empty
}

func (t *Env) IsDev() bool {
	return t.Type() == development
}

func (t *Env) IsProd() bool {
	return t.Type() == production
}

func (t *Env) IsTest() bool {
	return t.Type() == testing
}

func (t *Env) IsUnknown() bool {
	return t.Type() == unknown
}

func CreateDevEnv() *Env {
	return &Env{name: defaultDevEnv}
}

func CreateProdEnv() *Env {
	return &Env{name: defaultProdEnv}
}

func CreateTestEnv() *Env {
	return &Env{name: defaultTestEnv}
}

func CreateEmptyEnv() *Env {
	return &Env{name: defaultEmptyEnv}
}

func CreateEnvByName(name string) *Env {
	return &Env{name: name}
}
