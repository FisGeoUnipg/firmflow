package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type firmwarelist struct {
	curr  string
	queue []string
}

func (f *firmwarelist) show(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Refresh", "1; http://__IP__:9090/status")
		fmt.Fprint(w, "Firmware in esecuzione: "+f.curr+"<br>")
		fmt.Fprint(w, "Coda:<br>")
		for _, ot := range f.queue {
			fmt.Fprintln(w, " "+ot+"<br>")
		}
	}

}

func (f *firmwarelist) console(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Refresh", "1; http://__IP__:9090/console#fine")
		if dat, err := ioutil.ReadFile("/app/bitstreams/Console"); err == nil {
			lines:=strings.Split(string(dat),"\n")
			for i:=len(lines)-1;i>0;i-- {
				if lines[i]!= "" {
					fmt.Fprint(w, lines[i], "<br />")
				}
			}
		}
	}

}

func (f *firmwarelist) update() {
	for {
		f.curr = ""
		f.queue = make([]string, 0)

		root := "/app/bitstreams/"
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".bit" {
				f.queue = append(f.queue, filepath.Base(path))
			}
			return nil
		})

		if dat, err := ioutil.ReadFile(root+"Metadata"); err == nil {
			f.curr = string(dat)
		}

		time.Sleep(1000 * time.Millisecond)
	}

}

func upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		studente := r.FormValue("studente")
		esercizio := r.FormValue("esercizio")
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Refresh", "0; http://__IP__")
		fmt.Fprint(w, "Upload done! auto-reload in 5 seconds")
		filename := strings.ReplaceAll("/app/bitstreams/"+studente+"_"+esercizio+"_"+handler.Filename, " ", "")
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {

	f := new(firmwarelist)
	f.curr = ""
	f.queue = make([]string, 0)

	http.HandleFunc("/console", f.console)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/status", f.show)
	go f.update()
	log.Fatal(http.ListenAndServe(":9090", nil))
}
