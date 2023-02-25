package controller

import (
	"api-go/src/database"
	"api-go/src/model"
	"api-go/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func SaveUsuario(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var usuario model.Usuario

	if err = json.Unmarshal(body, &usuario); err != nil {
		log.Fatal(err)
	}

	db, err := database.GetConnection()

	if err != nil {
		log.Fatal(err)
	}

	repository := repository.CreateRepository(db)
	id, err := repository.Save(usuario)

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Id: %d", id)))
}

func GetAllUsuarios(w http.ResponseWriter, r *http.Request) {

}

func GetUsuarioById(w http.ResponseWriter, r *http.Request) {

}

func UpdateUsuario(w http.ResponseWriter, r *http.Request) {

}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {

}
