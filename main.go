package main

import "log"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: env. file not found:")
	}
}
