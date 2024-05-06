package handlers

import (
	"encoding/json"
	"os/exec"
	"fmt"
	"io/ioutil"
	"net/http"
)

func MakePdf(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		exec.Command("bash", "makepdf.sh")

		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode("Success!")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// https://tutorialedge.net/golang/go-file-upload-tutorial/
func UploadText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		// Get and retrieve bytes from multipart upload
		r.ParseMultipartForm(10 << 20)
		file, handler, err := r.FormFile("paper.tex")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)

		// Save into temp file
		tempFile, err := ioutil.TempFile("assets", "upload-*.tex")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		// Save into final file
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		tempFile.Write(fileBytes)
		fmt.Fprintf(w, "Successfully Uploaded File\n")

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode("Success!")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}