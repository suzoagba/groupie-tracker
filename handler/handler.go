package handler

import (
	"fmt"
	"groupie-tracker/artists"
	"groupie-tracker/variables"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("PANIC:", err)
			errorHandler(w, http.StatusInternalServerError)
		}
	}()
	if !validateRequest(r.Header) {
		errorHandler(w, http.StatusBadRequest)
		return
	}
	switch r.URL.Path {
	case "/":
		tmpl, err := template.ParseFiles("templates/index.gohtml", "templates/filters.gohtml", "templates/artists.gohtml", "templates/artistInfo.gohtml")
		if err != nil {
			panic(err)
		}
		artists.FilterArtists(r.FormValue("search"), r.FormValue("creationDateFrom"), r.FormValue("creationDateTo"), r.FormValue("firstAlbumDateFrom"), r.FormValue("firstAlbumDateTo"), r.Form["numMembers"], r.FormValue("locations"))
		formValues(r.FormValue("search"), r.FormValue("locations"), r.Form["numMembers"], r.FormValue("firstAlbumDateFrom"), r.FormValue("firstAlbumDateTo"), r.FormValue("creationDateTo"), r.FormValue("creationDateFrom"))
		if err := tmpl.Execute(w, variables.NewPage); err != nil {
			fmt.Fprint(w, err)
		}
	default:
		errorHandler(w, http.StatusNotFound)
	}
}

func errorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	errorStruct := make(map[string]string)
	errorStruct["ErrorNumber"] = strconv.Itoa(status)
	tmpl, err := template.ParseFiles("templates/error.gohtml")
	if err != nil {
		panic(err)
	}
	if status == http.StatusNotFound {
		errorStruct["ErrorText1"] = "Oops! Looks like you got lost."
		errorStruct["ErrorText2"] = "Page not found."
	} else if status == http.StatusBadRequest {
		errorStruct["ErrorText1"] = "Oops! Bad Request!"
		errorStruct["ErrorText2"] = "You can only request text/html."
	} else if status == http.StatusInternalServerError {
		errorStruct["ErrorText1"] = "Oops! This is awkward."
		errorStruct["ErrorText2"] = "Internal Server Error."
	}
	tmpl.Execute(w, errorStruct)
}

// Check if the header "Accept" contains text/html
func validateRequest(h http.Header) bool {
	reg := regexp.MustCompile(`(?m)text\/html`)
	for _, header := range h["Accept"] {
		match := reg.Match([]byte(header))
		if match {
			return true
		}
	}
	return false
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/favicon.ico")
}
