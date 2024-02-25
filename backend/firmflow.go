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
	"strconv"
)

type queueList struct {
	curr  string
	queue []string
	boardNumber int
}

type firmwarelist struct {
	queueList []queueList
}

func (f *firmwarelist) show(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		// get the board number from the r http request
		board := r.URL.Query().Get("board")
		boardNumber, _ := strconv.Atoi(board)

		// find inside f.queueList the element with the board number equal to boardNumber
		var targetBoard queueList

		for _, q := range f.queueList {
            if q.boardNumber == boardNumber {
                targetBoard = q
                break
            }
        }

		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Refresh", "1; http://__IP__:9090/status?board="+board)
		fmt.Fprint(w, "Firmware in esecuzione: "+targetBoard.curr+"<br>")
		fmt.Fprint(w, "Coda:<br>")
		for _, ot := range targetBoard.queue {
			fmt.Fprintln(w, " "+ot+"<br>")
		}
	}

}

func (f *firmwarelist) console(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// get the board number from the r http request
		board := r.URL.Query().Get("board")
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Refresh", "1; http://__IP__:9090/console#fine")
		if dat, err := ioutil.ReadFile("/app/bitstreams/"+board+"/Console"); err == nil {
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
		// iterate over queue list
		for i := 0; i < len(f.queueList); i++ {
			f.queueList[i].curr = ""
			f.queueList[i].queue = make([]string, 0)

			root := "/app/bitstreams/"+strconv.Itoa(f.queueList[i].boardNumber)+"/"
			filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				if filepath.Ext(path) == ".bit" {
					f.queueList[i].queue = append(f.queueList[i].queue, filepath.Base(path))
				}
				return nil
			})
			if dat, err := ioutil.ReadFile(root+"Metadata"); err == nil {
				f.queueList[i].curr = string(dat)
			}
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
		board := r.FormValue("board")
		file, _, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Refresh", "0; http://__IP__/page"+board+".html")
		fmt.Fprint(w, "Upload done! auto-reload in 5 seconds")

		// if folder /app/bitstreams/board does not exist, create it
		if _, err := os.Stat("/app/bitstreams/" + board); os.IsNotExist(err) {
			os.Mkdir("/app/bitstreams/"+board, 0755)
		}

		filename := strings.ReplaceAll("/app/bitstreams/"+board+"/"+studente+"_"+esercizio+".bit", " ", "")
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

	boardNumber := __BOARDNUMBER__
	f := new(firmwarelist)

	for i := 1; i < boardNumber+1; i++ {
		f.queueList = append(f.queueList, queueList{curr: "", queue: make([]string, 0), boardNumber: i})
	}

	http.HandleFunc("/console", f.console)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/status", f.show)
	go f.update()
	log.Fatal(http.ListenAndServe(":9090", nil))
}
