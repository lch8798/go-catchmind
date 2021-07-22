package utils

import (
	"os"
	"strconv"
)

const DEFAULT_PORT int = 8000

func GetPort() (int) {
	// set port from env
	var port int;
	var err error;
	var envPort string = os.Getenv("PORT")
	port, err = strconv.Atoi(envPort)

	// set default port
	if err != nil {
		port = DEFAULT_PORT
		err = nil
	}

	return port
}