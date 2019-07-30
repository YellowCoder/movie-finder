package env

import (
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

func loadDotEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	_, filename, _, _ := runtime.Caller(1)
	file := path.Join(path.Dir(filename), "../.env."+env)
	godotenv.Load(file)
}
