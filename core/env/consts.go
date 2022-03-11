package env

const (
	EnvVar = "ENV"

	defaultEmptyEnv = ""
	defaultDevEnv   = "dev"
	defaultProdEnv  = "prod"
	defaultTestEnv  = "test"
)

var (
	emptyEnvs = []string{defaultEmptyEnv}
	devEnvs   = []string{defaultDevEnv, "development"}
	prodEnvs  = []string{defaultProdEnv, "production"}
	testEnvs  = []string{defaultTestEnv}
)
