package main

import (
	"fmt"

	"github.com/bamchoh/envlist"
)

func main() {
	fmt.Println("=== ALL ENVIRONMENT ===")
	for k, v := range envlist.GetAllEnvList() {
		fmt.Println(k, "\t===\t", v)
	}

	fmt.Println("=== SYSTEM ENVIRONMENT ===")
	sysenvs, _ := envlist.GetSystemEnvList()
	for k, v := range sysenvs {
		fmt.Println(k, "\t===\t", v)
	}

	fmt.Println("=== CURRENT USER ENVIRONMENT ===")
	cuenvs, _ := envlist.GetCurrentUserEnvList()
	for k, v := range cuenvs {
		fmt.Println(k, "\t===\t", v)
	}
}
