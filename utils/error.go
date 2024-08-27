package utils

import (
	"html/template"
	"net/http"
)

func HandleError(w http.ResponseWriter, msg string, code int) {
	tmpl, err := template.ParseFiles("./templates/html/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, struct {
		Msg  string
		Code int
	}{
		Msg:  msg,
		Code: code,
	})
}
