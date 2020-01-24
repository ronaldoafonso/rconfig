/* File: web.go */
/* Description: Web Service for remote boxes */

package main

import (
	"fmt"
	"github.com/ronaldoafonso/rconfig/rconfig"
	"github.com/ronaldoafonso/rconfig/rmac"
	"net/http"
	"strings"
)

// Box ... An OpenWrt box struct
type Box struct {
	Boxname string
	rconfig.Config
}

var (
	box1 = Box{
		Boxname: "box1",
		Config: rconfig.Config{
			SSID:      "SSID1",
			LeaseTime: "1m",
			AllowedMACs: rmac.AllowedMACs{
				"11:11:11:11:11:11",
			},
		},
	}
)

func main() {
	server := &http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", handleWebBoxesRequests)
	fmt.Println("Ready to handle rconfig web requests ...")
	server.ListenAndServe()
}

// handleWebBoxesRequests ... Handle REST web boxes requests
func handleWebBoxesRequests(w http.ResponseWriter, r *http.Request) {
	urlPathElements := strings.Split(r.URL.Path, "/")
	if urlPathElements[1] == "boxes" {
		if urlPathElements[2] == "box1" {
			fmt.Fprintf(w, "%s %s %s %v\n",
				box1.Boxname,
				box1.Config.SSID,
				box1.Config.LeaseTime,
				box1.Config.AllowedMACs)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not Found"))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
	}
}
