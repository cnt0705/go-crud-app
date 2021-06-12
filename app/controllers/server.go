package controllers

import (
	"fmt"
	"net/http"
	"text/template"
	"todo-module/app/models"
	"todo-module/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string

	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")

	if err != nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}

	return sess, err
}

func StartMainServer() error {
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
