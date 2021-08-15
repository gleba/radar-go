package sio

import (
	"log"
	"net/http"
)

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
//	pegas.SyncPresets()
//}

func StartHttp(handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc("/sync", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
