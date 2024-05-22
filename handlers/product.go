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

func StoreProduct(w http.ResponseWriter, r *http.Request) {
	payload, _ := io.ReadAll(r.Body)
	data := []byte(payload)

	product := new(model.Product)
	err := json.Unmarshal(data, &product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request gagal 1"))

		return
	}

	db := database.ConnectDatabase()
	result := db.Create(product)

	if result.RowsAffected < 0 {
		log.Fatal(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request gagal 2"))

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("product created"))
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDatabase()

	// user := new(app.User)
	var product []model.Product

	if err := db.Find(&product).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
	}

	output, err := json.MarshalIndent(product, "", " ")
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	payload, _ := io.ReadAll(r.Body)
	data := []byte(payload)

	var product model.Product
	// get data by id
	db := database.ConnectDatabase()
	db.First(&product, id)

	// decode json
	err := json.Unmarshal(data, &product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request gagal"))
	}

	// process Update
	db.Save(&product)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user updated"))
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	db := database.ConnectDatabase()
	db.Delete(&model.Product{}, id)

	w.WriteHeader(http.StatusOK)
	response := respond.RespDeleteSuccess{
		Message:    "Data berhasil dihapus",
		StatusCode: 200,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(response)

}
