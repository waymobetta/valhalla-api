package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/waymobetta/valhalla-api/app"
)

var (
	home, _ = homedir.Dir()
	fil     = fmt.Sprintf("%s/go/src/github.com/waymobetta/valhalla-api/db/%s", home, "db.json")
)

func Vis() (app.GodkjentCollection, *app.StandardError) {
	var liste []*app.Godkjent
	var g app.GodkjentCollection

	// lese fil
	dataArr, err := ioutil.ReadFile(fil)
	if err != nil {
		return g, &app.StandardError{
			Code:    500,
			Message: "lese fil error",
		}
	}

	// unmarshal
	err = json.Unmarshal(dataArr, &liste)
	if err != nil {
		return g, &app.StandardError{
			Code:    500,
			Message: "JSON unmarshal error",
		}
	}

	for _, kjent := range liste {
		g = append(g, kjent)
	}

	return g, nil
}

func LeggeTil(navn, adresse string) *app.StandardError {
	// nye Godkjent struct
	gk := &app.Godkjent{
		Navn:    navn,
		Adresse: adresse,
	}

	// nye liste av Godkjent struct
	var liste []*app.Godkjent

	// lese db fil
	dataArr, err := ioutil.ReadFile(fil)
	if err != nil {
		return &app.StandardError{
			Code:    500,
			Message: "fil lese error",
		}
	}

	// unmarshal
	err = json.Unmarshal(dataArr, &liste)
	if err != nil {
		return &app.StandardError{
			Code:    500,
			Message: "JSON unmarshal error",
		}
	}

	// legge til liste
	liste = append(liste, gk)

	// marshal
	dataArr, err = json.Marshal(&liste)
	if err != nil {
		return &app.StandardError{
			Code:    500,
			Message: "JSON marshal error",
		}
	}

	// legge til fil
	err = ioutil.WriteFile(fil, dataArr, 0644)
	if err != nil {
		return &app.StandardError{
			Code:    500,
			Message: "skriver fil error",
		}
	}

	return nil
}
