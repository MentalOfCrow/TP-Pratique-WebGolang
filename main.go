package main

import (
	"html/template"
	"net/http"
)

type Student struct {
	Nom    string
	Prenom string
	Age    int
	Sexe   string
}

type Promotion struct {
	Nom       string
	Filiere   string
	Niveau    string
	Nombre    int
	Etudiants []Student
}

func main() {
	promotion := Promotion{
		Nom:     "Promotion 2023",
		Filiere: "Informatique",
		Niveau:  "Bac+3",
		Nombre:  5,
		Etudiants: []Student{
			{"Cyril", "RODRIGUES", 22, "homme"},
			{"Kheir-eddine", "MEDERREG", 22, "homme"},
			{"Alan", "PHILIPIERT", 26, "homme"},
		},
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("template/template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, promotion)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
