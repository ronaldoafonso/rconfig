/* File: backup.go */
/* Description: Configuration backup for remote boxes */

package main

import (
	"database/sql"
	"log"
	// PostgreSQL driver
	_ "github.com/lib/pq"
	"github.com/ronaldoafonso/rconfig/rbox"
)

func main() {
	connString := "user=rconfig dbname=rconfig password=rconfig " +
		"host=rconfig_db sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT boxname FROM boxes;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	boxes := make(map[string]rbox.Box)

	for rows.Next() {
		var boxname string
		err = rows.Scan(&boxname)
		if err != nil {
			log.Fatal(err)
		}

		boxes[boxname] = rbox.Box{
			Boxname: boxname,
		}
	}

	if box, ok := boxes["boxname"]; ok {
		backup(box)
	}

	if box, ok := boxes["boxname1"]; ok {
		backup(box)
	}

	if box, ok := boxes["boxname2"]; ok {
		backup(box)
	}
}

func backup(box rbox.Box) {
	log.Printf("Doing %q.\n", box.Boxname)
	if err := box.GetRemoteSSID(); err != nil {
		log.Fatal(err)
	}

	if err := box.GetRemoteAllowedMACs(); err != nil {
		log.Fatal(err)
	}

	if err := box.UpdateConfig(); err != nil {
		log.Fatal(err)
	}
}
