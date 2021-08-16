package rbox

import (
	"io/ioutil"
	"testing"
)

func TestNewRBox(t *testing.T) {
	boxname := "788a20298f81.z3n.com.br"
	var b interface{} = NewRBox(boxname)
	if _, ok := b.(RemoteBox); !ok {
		t.Fatal("Error: not a RemoteBox type")
	}
}

func TestRBoxGetConfig(t *testing.T) {
	boxname := "788a20298f81.z3n.com.br"
	b := NewRBox(boxname)
	err := b.GetConfig()
	if err != nil {
		t.Fatalf("Error getting config for %v.", boxname)
	}

	f1, err := ioutil.ReadFile("./uci_show")
	if err != nil {
		t.Fatalf("Error reading uci_show: %v.", err)
	}

	f2, err := ioutil.ReadFile(boxname)
	if err != nil {
		t.Fatalf("Error reading %v: %v.", boxname, err)
	}

	if len(f1) != len(f2) {
		t.Fatal("Error config files are not the same size.")
	}

	for index, data := range f1 {
		if data != f2[index] {
			t.Fatal("Error: files are different.")
		}
	}
}

func TestRBoxGetSSIDs(t *testing.T) {
	boxname := "788a20298f81.z3n.com.br"
	b := NewRBox(boxname)

	SSIDs, err := b.GetSSIDs()
	if err != nil {
		t.Fatalf("Error getting SSIDs: %v.", err)
	}

	for _, SSID := range SSIDs {
		if SSID != "z3n" {
			t.Fatalf("Error getting SSID. Want: z3n, got: '%v'.", SSID)
		}
	}
}

func TestRBoxGetMACs(t *testing.T) {
	boxname := "788a20298f81.z3n.com.br"
	b := NewRBox(boxname)

	MACs, err := b.GetMACs()
	if err != nil {
		t.Fatalf("Error getting MACs: %v.", err)
	}

	tests := []struct {
		want string
	}{
		{"11:11:11:11:11:11"},
		{"11:11:11:11:11:22"},
		{"11:11:11:11:11:33"},
	}
	for i, test := range tests {
		if test.want != MACs[i] {
			t.Fatalf("Error getting MAC. Want: %v, got: %v.", test.want, MACs[i])
		}
	}
}
