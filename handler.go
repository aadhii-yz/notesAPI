package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HandleGetNotes(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
}

func HandleGetNoteByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
}

// POST /notes
// gets Request with json body and adds in the database
func HandlePostNote(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var note Note

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	note = Store.Append(note)
	fmt.Fprintf(w, "Request Recived: %v", note)
}

func HandlePutNoteByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
}

func HandleDeleteNoteByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
}
