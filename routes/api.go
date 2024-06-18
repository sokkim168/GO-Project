package routes

import (
	"net/http"

	"example.com/booking/internal/controller"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(r *mux.Router) {
	room := &controller.Rooms{}
	booking := &controller.RoomBooking{}
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/rooms", room.Index).Methods(http.MethodGet)
	api.HandleFunc("/rooms/{id:[0-9]+}", room.Show).Methods(http.MethodGet)
	api.HandleFunc("/rooms", room.Store).Methods(http.MethodPost)
	api.HandleFunc("/rooms/{id:[0-9]+}", room.Update).Methods(http.MethodPut)
	api.HandleFunc("/rooms/{id:[0-9]+}", room.Delete).Methods(http.MethodDelete)

	api.HandleFunc("/room-bookings", booking.Index).Methods(http.MethodGet)
	api.HandleFunc("/room-bookings/{id:[0-9]+}", booking.Show).Methods(http.MethodGet)
	api.HandleFunc("/room-bookings", booking.Store).Methods(http.MethodPost)
	api.HandleFunc("/room-bookings/{id:[0-9]+}", booking.Update).Methods(http.MethodPut)
	api.HandleFunc("/room-bookings/{id:[0-9]+}", booking.Delete).Methods(http.MethodDelete)
}
