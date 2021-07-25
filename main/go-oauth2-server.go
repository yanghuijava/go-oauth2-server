package main

import (
	"github.com/sirupsen/logrus"
	"go-oauth2-server/server"
)

func init() {
	logrus.SetReportCaller(true)
}

func main() {
	server.Run()
}
