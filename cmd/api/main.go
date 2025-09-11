package main

import "github.com/ahmdyaasiin/workshop-ci-cd/internal/bootstrap"

func main() {
	if err := bootstrap.Initialize(); err != nil {
		panic(err)
	}
}
