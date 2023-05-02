package helper

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	hostEnvKey  = "HOST"
	portEnvKey  = "PORT"
	defaultHost = "0.0.0.0"
)

func GetAddress(envKeyPrefix, defaultPort string) string {
	host := GetEnv(strings.Join([]string{envKeyPrefix, hostEnvKey}, "_"), defaultHost)
	p := GetEnv(strings.Join([]string{envKeyPrefix, portEnvKey}, "_"), defaultPort)
	port, err := strconv.Atoi(p)
	if err != nil {
		log.Printf("failed to parse port: %s", err)
	}

	return fmt.Sprintf("%s:%d", host, port)
}

func GetEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultValue
	}
	return val
}

func GetQueryParam(r *http.Request, key string) (string, error) {
	value := r.URL.Query().Get(key)
	if len(value) == 0 {
		return "", fmt.Errorf("%s is required", key)
	}

	return value, nil
}

func HttpError(w http.ResponseWriter, status int, message string, err error) {
	log.Printf(message+": %s", err)
	w.WriteHeader(status)
	if _, err := w.Write([]byte(err.Error())); err != nil {
		log.Printf("failed to write response body: %s", err)
	}
}
