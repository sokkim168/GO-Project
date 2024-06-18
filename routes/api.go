package routes

import (
	"net/http"

	"example.com/booking/internal/controller"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(r *mux.Router) {
	room := &controller.Rooms{}
	booking := &controller.RoomBooking{}
	r.HandleFunc("/rooms", room.Index).Methods(http.MethodGet)
	r.HandleFunc("/rooms/{id:[0-9]+}", room.Show).Methods(http.MethodGet)
	r.HandleFunc("/rooms", room.Store).Methods(http.MethodPost)
	r.HandleFunc("/rooms/{id:[0-9]+}", room.Update).Methods(http.MethodPut)
	r.HandleFunc("/rooms/{id:[0-9]+}", room.Delete).Methods(http.MethodDelete)

	r.HandleFunc("/room-bookings", booking.Index).Methods(http.MethodGet)
	r.HandleFunc("/room-bookings/{id:[0-9]+}", booking.Show).Methods(http.MethodGet)
	r.HandleFunc("/room-bookings", booking.Store).Methods(http.MethodPost)
	r.HandleFunc("/room-bookings/{id:[0-9]+}", booking.Update).Methods(http.MethodPut)
	r.HandleFunc("/room-bookings/{id:[0-9]+}", booking.Delete).Methods(http.MethodDelete)
}
