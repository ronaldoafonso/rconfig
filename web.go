/* File: web.go */
/* Description: Web Service for remote boxes */

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ronaldoafonso/rconfig/rconfig"
	"github.com/ronaldoafonso/rconfig/rmac"
	"net/http"
)

// Box ... An OpenWrt box struct
type Box struct {
	Boxname string
	rconfig.Config
}

var (
	boxes = make(map[string]Box)
	box1  = Box{
		Boxname: "box1",
		Config: rconfig.Config{
			SSID:      "SSID1",
			LeaseTime: "1m",
			AllowedMACs: rmac.AllowedMACs{
				"11:11:11:11:11:11",
			},
		},
	}
	box2 = Box{
		Boxname: "box2",
		Config: rconfig.Config{
			SSID:      "SSID2",
			LeaseTime: "2m",
			AllowedMACs: rmac.AllowedMACs{
				"22:22:22:22:22:22",
			},
		},
	}
	box3 = Box{
		Boxname: "box3",
		Config: rconfig.Config{
			SSID:      "SSID3",
			LeaseTime: "3m",
			AllowedMACs: rmac.AllowedMACs{
				"33:33:33:33:33:33",
			},
		},
	}
)

func init() {
	boxes["box1"] = box1
	boxes["box2"] = box2
	boxes["box3"] = box3
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/box/{box}", handleWebBoxesRequests)
	server := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}
	fmt.Println("Ready to handle rconfig web requests ...")
	server.ListenAndServe()
}

// showBox ... Show box struct to clients
func showBox(w http.ResponseWriter, box *Box) {
	fmt.Fprintf(w, "%s %s %s %v\n",
		box.Boxname,
		box.Config.SSID,
		box.Config.LeaseTime,
		box.Config.AllowedMACs)
}

// handleWebBoxesRequests ... Handle REST web boxes requests
func handleWebBoxesRequests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boxname := vars["box"]
	if box, ok := boxes[boxname]; ok {
		showBox(w, &box)
	} else if boxname == "all" {
		for _, box := range boxes {
			showBox(w, &box)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
	}
}
