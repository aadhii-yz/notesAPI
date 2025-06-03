package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HandleGetNotes(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
}

func HandleGetNoteByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
}

func HandlePostNote(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
}

func HandlePutNoteByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
}

func HandleDeleteNoteByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
}
