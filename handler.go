package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GET /notes
func HandleGetNotes(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	outputStr := "Notes {\n"
	for _, value := range Store.data {
		jsonStr, _ := json.MarshalIndent(value, string(value.ID), "	")
		outputStr = fmt.Sprintf("%s\n%s\n", outputStr, jsonStr)
	}
	outputStr = fmt.Sprintf("%s\n}", outputStr)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, outputStr)
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
	w.WriteHeader(http.StatusCreated)
	jsonStr, _ := json.MarshalIndent(note, "", "	")
	fmt.Fprintf(w, "Request Recived: %s", jsonStr)
}

func HandlePutNoteByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
}

func HandleDeleteNoteByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
}
