package models

import (
	"context"
	"database/sql"
	"fmt"
	"lms/pkg/types"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

const configPath = "config.yaml"

func NewConfig(path string) (*types.Config, error) {
	// Create config structure
	config := &types.Config{}

	//open config file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

func dsn(config *types.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", config.UserNAME, config.Password, config.Host, config.DB)
}

func Connection() (*sql.DB, error) {
	config, err := NewConfig(configPath)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", dsn(config))
	if err != nil {
		log.Printf("Error: %s when opening DB", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	log.Printf("Connected to DB successfully\n")
	return db, err

}
