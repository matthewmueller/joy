package env

import "os"

// e.g.
// go run -ldflags="-X 'github.com/matthewmueller/joy/internal/env.awsAccessKey=hello'" _examples/env/env.go
var awsAccessKey = ""
var awsSecretKey = ""
var awsRegion = ""
var firehoseStream = ""

// Env is the environment
type Env struct {
	AWSAccessKey   string
	AWSSecretKey   string
	AWSRegion      string
	FirehoseStream string
}

// Get the environment
func Get() Env {
	return Env{
		AWSAccessKey:   getEnv("AWS_ACCESS_ID", awsAccessKey),
		AWSSecretKey:   getEnv("AWS_SECRET_ACCESS_KEY", awsSecretKey),
		AWSRegion:      getEnv("AWS_REGION", awsRegion),
		FirehoseStream: getEnv("AWS_FIREHOSE_STREAM", firehoseStream),
	}
}

func getEnv(name string, def string) string {
	if os.Getenv(name) != "" {
		return os.Getenv(name)
	}
	return def
}
