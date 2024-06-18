package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/booking/models"
	"github.com/gorilla/mux"
)

type Rooms struct {
	Model *models.Room
}

func (c *Rooms) Index(w http.ResponseWriter, r *http.Request) {
	rooms := c.Model.All()
	res, _ := json.Marshal(rooms)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c *Rooms) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
	}
	rooms := models.Find(int32(id))
	res, _ := json.Marshal(rooms)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c Rooms) Store(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&c.Model); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c.Model.Create()
	res, _ := json.Marshal(&c.Model)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c *Rooms) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := json.NewDecoder(r.Body).Decode(&c.Model); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(vars["id"])
	get := c.Model.Update(int32(id))
	res, _ := json.Marshal(get)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c *Rooms) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	get := models.Delete(int32(id))
	res, _ := json.Marshal(get)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
