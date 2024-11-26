package main

import (
	"fmt"
	"net/http"

	v "github.com/root27-dev/htmx-auth/views"
)

func main() {

	users := make(map[string]string)

	// Pages

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		index := v.Index()

		index.Render(r.Context(), w)

	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {

		email := r.FormValue("email")
		password := r.FormValue("password")

		if _, ok := users[email]; ok {
			//USer already exists

			register := v.Register(true, false)

			register.Render(r.Context(), w)

		}

		users[email] = password

		register := v.Register(false, true)

		register.Render(r.Context(), w)

	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		login := v.Login()

		login.Render(r.Context(), w)

	})

	http.HandleFunc("/servelogin", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Hx-Redirect", "/login")

	})

	// Handlers

	fmt.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", nil)

}
