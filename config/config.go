package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Values ci = &conf{}

type ci interface {
	Get() *conf
}

type conf struct {
	PortNumber string
	host       string
	dbname     string
	user       string
	password   string
	sslmode    string
	ConnStr    string
}

func (m *conf) Get() *conf {
	log.Println("=== config Get ===")

	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	m.PortNumber = os.Getenv("PORT")
	m.host = os.Getenv("DB_HOST")
	m.dbname = os.Getenv("DB_NAME")
	m.user = os.Getenv("DB_USER")
	m.password = os.Getenv("DB_PASSWORD")
	m.sslmode = os.Getenv("DB_SSL_MODE")
	m.ConnStr = fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s", m.host, m.dbname, m.user, m.password, m.sslmode)

	return m
}
