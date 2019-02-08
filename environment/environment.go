package environment

import "github.com/OlympBMSTU/test-second-service/config"

// import "github.com//config"

type ServerEnvironment struct {
	FirstExercisesService Serivice
	FirstExercisesService Serivice
	Logger                ILogger
	Configurator          Config
}

func New() ServerEnvironment {
	// logger := logger.New()
	cfg, err := config.GetConfigInstance()
	// todo read config file
}

func (env *Environment) Configure() {

}
