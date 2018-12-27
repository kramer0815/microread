package service

import (
	"os"
	"github.com/gorilla/mux"
	"github.com/SlyMarbo/rss"
	"net/http"
	"encoding/json"
	"strconv"
	"fmt"
	"github.com/callistaenterprise/goblog/accountservice/model"
)

// GetRssFeed fetches a atom.xml and gets the headline by id
func GetRssFeed(w http.ResponseWriter, r *http.Request) {

	// RSS Feed beziehen
	feed, err := rss.Fetch("http://www.heise.de/newsticker/heise-atom.xml")
	if err != nil {
		println("RSS Feed nicht erreichbar")
	}
	
	//ID aus Handler holen
	var commentid = mux.Vars(r)["number"]
	
	count, err := strconv.Atoi(commentid)
   	 if err != nil {
    	    fmt.Println(count)
   	 }
	
	fmt.Fprintf(os.Stderr, "feedentry: %d " + feed.Items[count].Title + "\n", count)

	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// If found, marshal into JSON, write headers and content
	var line = feed.Items[count].Title
	// Put Data into Struct
	var feedentry = model.FeedEntry{Id:commentid, Name:line}
	//Marshel
	data, _ := json.Marshal(feedentry)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}