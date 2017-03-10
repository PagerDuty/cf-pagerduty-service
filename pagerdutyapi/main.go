package main

import (
	"cf-pagerduty-service/pagerdutyapi/route"
	"log"
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
)

// Credentials basic auth credentials
type Credentials struct {
	Username string
	Password string
}

const (
	defaultPort     = "8080"
	defaultUsername = "admin"
	defaultPassword = "pagerduty"
)

var c *config

type config struct {
	username string
	password string
	port     string
}

func main() {
	logger := lager.NewLogger("p-pagerduty-api")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
	logger.RegisterSink(lager.NewWriterSink(os.Stderr, lager.ERROR))

	c = configFromEnvironmentVariables()

	router := wrapper(route.NewRouter())

	logger.Fatal("http-listen", http.ListenAndServe(":"+getPort(), router))
}

func configFromEnvironmentVariables() *config {
	conf := &config{
		username: getEnv("basicauth_username", defaultUsername),
		password: getEnv("basicauth_password", defaultPassword),
		port:     getPort(),
	}

	return conf
}

func getPort() string {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = defaultPort
	}

	return port
}

func getEnv(env string, defaultValue string) string {
	var v string
	if v = os.Getenv(env); len(v) == 0 {
		log.Printf("Using default value for %v", env)
		return defaultValue
	}

	return env
}

func wrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (len(c.username) > 0) && (len(c.password) > 0) && !authorized(r, c.username, c.password) {
			w.Header().Set("WWW-Authenticate", `Basic realm="REALM"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func authorized(r *http.Request, user, pass string) bool {
	if username, password, ok := r.BasicAuth(); ok {
		return username == user && password == pass
	}

	return false
}
