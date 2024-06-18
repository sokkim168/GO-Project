package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"example.com/booking/routes"
	"github.com/gorilla/mux"
)

func main() {
	// db := config.GetDB()  // TEST CONNECTION

	/* TEST GET ALL ROOM*/
	// res := models.All()
	// fmt.Println(res)

	/* TEST GET BY ID*/
	// getRoom := models.Find(8)
	// fmt.Println(getRoom)

	/* TEST GET DELETE*/
	// deleteRoom := models.Delete(9)
	// fmt.Println(deleteRoom)

	/* TEST ROUTES*/

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	app_url := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	listner, err := net.Listen("tcp", app_url)
	fmt.Printf("APP HOST: %v", app_url)
	if err != nil {
		log.Fatal("Listen:", err)
	}
	err2 := http.Serve(listner, nil)
	if err2 != nil {
		log.Fatal("http.Serve:", err2)
	}

}
