package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DataTable struct {
	Id   int
	Data string
}

func main() {
	dsn := "root:example@tcp(localhost:3306)/test?autocommit=0"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
		os.Exit(1)
	}
	var row DataTable
	log.Println("About to fetch record for update")
	db.Exec("SET innodb_lock_wait_timeout=0")
	db.Transaction(func(tx *gorm.DB) error {
		tx.Exec("SET innodb_lock_wait_timeout=0")
		tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&row, 1)
		if tx.Error != nil {
			log.Printf("Could not fetch row %v", tx.Error)
			return db.Error
		}
		if row.Id == 0 { // need a way to detect if update lock could not be acquired
			log.Printf("Could not get lock on row with id 1. Skipping")
			return errors.New("could not acquire lock")
		}
		var b byte
		fmt.Printf("Enter a character: ")
		fmt.Scanf("%c", &b)
		row.Data = fmt.Sprintf("%v-%v", time.Now().UnixNano(), os.Getpid())
		tx.Save(&row)
		return nil
	})
}
