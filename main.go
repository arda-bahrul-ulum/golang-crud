package main

import (
	"golang-crud/controllers/pasiencontroller"
	"net/http"
)

func main()  {
	http.HandleFunc("/", pasiencontroller.Index)
	http.HandleFunc("/pasien", pasiencontroller.Index)
	http.HandleFunc("/pasien/index", pasiencontroller.Index)
	http.HandleFunc("/pasien/add", pasiencontroller.Add)
	http.HandleFunc("/pasien/edit", pasiencontroller.Edit)
	http.HandleFunc("/pasien/delete", pasiencontroller.Delete)

	http.ListenAndServe(":3000", nil)
}