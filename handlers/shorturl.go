package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"example.com/m/db"
	"github.com/gorilla/mux"
)

type shorturl struct {
	l *log.Logger
}

func NewShorturl(l *log.Logger) *shorturl {
	return &shorturl{l}
}

func (s *shorturl) GetShorturl(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("Handle GET shorturl")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["short"])
	if err != nil {
		http.Error(rw, "Unable to get original URL", http.StatusBadRequest)
		return
	}

	original, err := db.GetOriginal(id)
	if err != nil {
		http.Error(rw, "Unable to get original URL", http.StatusBadRequest)
		return
	}

	http.Redirect(rw, r, original, http.StatusFound)
}

func (s *shorturl) PostShorturl(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("Handle POST shorturl")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Unable to read body", http.StatusBadRequest)
		return
	}

	id, err := db.AddNew(string(data))
	if err != nil {
		http.Error(rw, "Unable to add URL", http.StatusInternalServerError)
		return
	}

	fmt.Println(id)
}
