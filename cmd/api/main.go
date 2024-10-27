package main

import (
	"database/sql"
	"log"

	"github.com/meetup/interals/env"
)

// type server struct {
// 	addr string
// }

// func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	switch r.Method {
// 	case http.MethodGet:
// 		switch r.URL.Path {
// 		case "/":
// 			w.Write([]byte("Hello From golang"))
// 			return
// 		case "/users":
// 			w.Write([]byte("Users Page"))
// 			return
// 		case "/data":
// 			w.Write([]byte("Data of users"))
// 			return
// 		}

// 	}
// }

// func main() {
// 	s := &server{addr: ":8080"}

// 	if err := http.ListenAndServe(s.addr, s); err != nil {
// 		panic(err.Error())
// 	}

// }

func main() {

	connStr := "http"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	app := App{config: Config{addr: env.GetEnv("ADDR", ":8080")}}
	mux := app.mount()
	if err := app.run(mux); err != nil {
		log.Fatal("Server Crashed")
	}
}
