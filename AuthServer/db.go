package main

import (
	"fmt"
	"github.com/leesper/couchdb-golang"
	"log"
	"os"
)

func connectionUrl() string {
	return fmt.Sprintf("http://%s:%s@%s:5984/authentication", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"))
}

func getDB() *couchdb.Database {
	db, err := couchdb.NewDatabase(connectionUrl())
	if err != nil {
		log.Fatalln(err.Error(), err)
	}
	return db
}

func saveTokenForAccount(email string, token string) string {
	results, err := getDB().Query([]string{"_id"}, fmt.Sprintf(`email == "%s"`, email),
		nil, nil, nil, nil)
	if err != nil {
		log.Println(err.Error(), err)
		return ""
	}
	log.Println(results)
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
