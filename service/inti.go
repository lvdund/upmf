package service

import "os"

func Start() (err error) {
	QueryPath()
	return
}

func Terminate() (err error) {
	os.Exit(0)
	return
}