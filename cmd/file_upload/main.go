package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func fileHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Print("error getting a file for the provided form key: ", err)
		return
	}
	defer file.Close()
	out, pathError := os.Create("/tmp/uploadedFile")
	if pathError != nil {
		log.Print("error creating a file for writing: ", pathError)
		return
	}

	defer out.Close()
	_, copyFileError := io.Copy(out, file)
	if copyFileError != nil {
		log.Print("error occurred while file copy: ", copyFileError)
	}
	fmt.Fprint(w, "File uploaded successfully: "+header.Filename)
}

func index(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("../../web/template/upload-file.html")
	//parsedTemplate, err := template.ParseFiles("upload-file.html")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(parsedTemplate.Name())
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", fileHandler)
	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}
