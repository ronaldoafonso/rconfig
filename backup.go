/* File: backup.go */
/* Description: Configuration backup for remote boxes */

package main

import (
	"database/sql"
	"log"
	// PostgreSQL driver
	_ "github.com/lib/pq"
	"github.com/ronaldoafonso/rconfig/rbox"
	"sync"
)

type returnCode struct {
	boxname string
	err     error
}

var (
	returnCodes = make(chan returnCode)
	wgBackup    sync.WaitGroup
	wgResult    sync.WaitGroup
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

	for _, box := range boxes {
		wgBackup.Add(1)
		go backup(box)
	}

	wgResult.Add(1)
	go showResults()

	wgBackup.Wait()
	close(returnCodes)

	wgResult.Wait()
	log.Println("Backup done.")
}

// backup ... Backup a box configuration
func backup(box rbox.Box) {
	defer wgBackup.Done()

	rc := returnCode{
		boxname: box.Boxname,
		err:     nil,
	}

	log.Printf("Backing up %v.\n", rc.boxname)

	if err1 := box.GetRemoteSSID(); err1 != nil {
		rc.err = err1
	} else if err2 := box.GetRemoteAllowedMACs(); err2 != nil {
		rc.err = err2
	} else if err3 := box.UpdateConfig(); err3 != nil {
		rc.err = err3
	}

	returnCodes <- rc
}

// showResults ... Show backup results
func showResults() {
	defer wgResult.Done()

	for rc := range returnCodes {
        if rc.err != nil {
            log.Printf("Box %v: error {%v}.\n", rc.boxname, rc.err)
        } else {
            log.Printf("Box %v: backup OK.\n", rc.boxname)
        }
	}
}
