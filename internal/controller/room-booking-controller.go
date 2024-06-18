package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/booking/models"
	"github.com/gorilla/mux"
)

type RoomBooking struct {
	Model *models.RoomBooking
}

func (c *RoomBooking) Index(w http.ResponseWriter, r *http.Request) {
	model := c.Model.All()
	res, _ := json.Marshal(model)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c *RoomBooking) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Booking ID", http.StatusBadRequest)
	}
	booking := c.Model.Find(int32(id))
	res, _ := json.Marshal(booking)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c RoomBooking) Store(w http.ResponseWriter, r *http.Request) {
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

func (c *RoomBooking) Update(w http.ResponseWriter, r *http.Request) {
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

func (c *RoomBooking) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	get := c.Model.Delete(int32(id))
	res, _ := json.Marshal(get)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
