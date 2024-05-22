package handlers

import (
	"encoding/json"
	"io"
	"log"
	"myproject/database"
	"myproject/model"
	"myproject/respond"
	"net/http"
)

func StoreUser(w http.ResponseWriter, r *http.Request) {
	payload, _ := io.ReadAll(r.Body)
	data := []byte(payload)

	user := new(model.User)
	err := json.Unmarshal(data, &user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request gagal 1"))

		return
	}

	db := database.ConnectDatabase()
	result := db.Create(user)

	if result.RowsAffected < 0 {
		log.Fatal(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request gagal 2"))

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDatabase()

	// user := new(app.User)
	var user []model.User

	if err := db.Find(&user).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
	}

	output, err := json.MarshalIndent(user, "", " ")
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	payload, _ := io.ReadAll(r.Body)
	data := []byte(payload)

	var user model.User
	// get data by id
	db := database.ConnectDatabase()
	db.First(&user, id)

	// decode json
	err := json.Unmarshal(data, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request gagal"))
	}

	// process Update
	db.Save(&user)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user updated"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	db := database.ConnectDatabase()
	db.Delete(&model.User{}, id)

	w.WriteHeader(http.StatusOK)
	response := respond.RespDeleteSuccess{
		Message:    "Data berhasil dihapus",
		StatusCode: 200,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(response)

}
