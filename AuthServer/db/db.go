package db

import (
	"fmt"
	"github.com/leesper/couchdb-golang"
	"github.com/meik99/CoffeeToGO/AuthServer/credentials"
	"log"
	"os"
)

func connectionUrl() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	if dbUser == "" {
		log.Println("environment variable DB_USER is not set")
	} else {
		log.Printf("using user '%s' to connect to database", dbUser)
	}
	if dbPassword == "" {
		log.Println("environment variable DB_PASSWORD is not set")
	}
	if dbHost == "" {
		log.Println("environment variable DB_HOST is not set")
	} else {
		log.Printf("using host '%s' to connect to database", dbHost)
	}

	return fmt.Sprintf("http://%s:%s@%s:5984/authentication", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"))
}

func getDB() *couchdb.Database {
	db, err := couchdb.NewDatabase(connectionUrl())
	if err != nil {
		log.Fatalln(err.Error(), err)
	}
	return db
}

func SaveTokenForAccount(email string, token credentials.AuthToken) string {
	results, err := getDB().Query([]string{"_id"}, fmt.Sprintf(`email == "%s"`, email),
		nil, nil, nil, nil)
	if err != nil {
		log.Println(err.Error(), err)
		return ""
	}

	if len(results) > 0 {
		for _, result := range results {
			err = getDB().Delete(result["_id"].(string))
			if err != nil {
				log.Println(err.Error(), err)
			}
		}
	}

	id := couchdb.GenerateUUID()
	id, _, err = getDB().Save(map[string]interface{}{
		"_id":   id,
		"email": email,
		"token": token,
	}, nil)
	if err != nil {
		log.Println(err.Error(), err)
	}

	return id
}
