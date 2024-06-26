package handlers

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"os"
	"fmt"
	"io/ioutil"
	"io"
	"net/http"
)

func MakePdf(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		outputChannel := make(chan string)

		go func() {
			cmd := exec.Command("bash", "makepdf.sh")
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
			    fmt.Println(err)
			}		
			outputChannel <- out.String()
		}()
	
		result := <-outputChannel
		fmt.Println(result)

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
// https://opensource.com/article/18/6/copying-files-go
func UploadText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		// Get and retrieve bytes from multipart upload
		r.ParseMultipartForm(10 << 20)
		file, handler, err := r.FormFile("paper.tex")
		// Should add file type (and parsing) validation in prod
		// no database but writes to disk

		// Ideally, the files would be written to temp external file store instead
		// but for the sake of this example and make 
		// it useable partially offline I'll keep it this way
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)

		// Save into temp file
		// Must be specified like this - w/ comma
		tempFile, err := ioutil.TempFile("assets", "paper.tex")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		//defer tempFile.Close()
		
		// Cleanup file
		defer os.Remove(tempFile.Name())
		
		// Save submitted data into tempfile
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		tempFile.Write(fileBytes)
		fmt.Fprintf(w, "Successfully Uploaded File\n")

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode("Success!")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		// Copy tempfile into fixed file
		input, err := ioutil.ReadFile(tempFile.Name())
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		destinationFile := "assets/paper.tex"
		err = ioutil.WriteFile(destinationFile, input, 0644)
		if err != nil {
			fmt.Println("Error creating", destinationFile)
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func SaveLaTeX(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		// https://www.practical-go-lessons.com/post/go-how-to-send-post-http-requests-with-a-json-body-cbhvuqa220ds70kp2lkg
		defer r.Body.Close()
		// This is a preferred way to extract textual info sent via XHR
		input, err := io.ReadAll(r.Body)
		// https://pkg.go.dev/io#example-ReadAll
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode("Successfully Uploaded LaTeX Text!")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		destinationFile := "assets/paper.tex"
		err = ioutil.WriteFile(destinationFile, input, 0644)
		if err != nil {
			fmt.Println("Error creating", destinationFile)
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}