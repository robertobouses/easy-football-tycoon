//go:build dev

package main

import "os"

func init() {
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASS", "mysecretpassword")
	os.Setenv("DB_DATABASE", "easy_football_tycoon")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")

}
