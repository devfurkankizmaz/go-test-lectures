package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

var pathToTemplates = "./templates/"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home", &templateData{})
}

type templateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) error {
	// Parse the template from the file
	tmpl, err := template.ParseFiles(path.Join(pathToTemplates, name+".gohtml"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	td.IP = app.ipFromCtx(r.Context())

	err = tmpl.Execute(w, td)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	log.Println(email, password)

	fmt.Fprintf(w, "Email: %s, Password: %s", email, password)
}
