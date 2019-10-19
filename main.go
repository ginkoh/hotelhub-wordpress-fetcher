package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	//"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Posts struct {
	Author        int    `json:"author"`
	Categories    []int  `json:"categories"`
	CommentStatus string `json:"comment_status"`
	Content       struct {
		Protected bool   `json:"protected"`
		Rendered  string `json:"rendered"`
	} `json:"content"`
	Date    string `json:"date"`
	DateGmt string `json:"date_gmt"`
	Excerpt struct {
		Protected bool   `json:"protected"`
		Rendered  string `json:"rendered"`
	} `json:"excerpt"`
	FeaturedMedia int    `json:"featured_media"`
	Format        string `json:"format"`
	GUID          struct {
		Rendered string `json:"rendered"`
	} `json:"guid"`
	ID   int    `json:"id"`
	Link string `json:"link"`
	Meta struct {
		_etPbOldContent string `json:"_et_pb_old_content"`
		_etPbUseBuilder string `json:"_et_pb_use_builder"`
	} `json:"meta"`
	Modified    string        `json:"modified"`
	ModifiedGmt string        `json:"modified_gmt"`
	PingStatus  string        `json:"ping_status"`
	Slug        string        `json:"slug"`
	Status      string        `json:"status"`
	Sticky      bool          `json:"sticky"`
	Tags        []interface{} `json:"tags"`
	Links struct {
		About []struct {
			Href string `json:"href"`
		} `json:"about"`
		Author []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
		} `json:"author"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		Curies []struct {
			Href      string `json:"href"`
			Name      string `json:"name"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
		PredecessorVersion []struct {
			Href string `json:"href"`
			ID   int    `json:"id"`
		} `json:"predecessor-version"`
		Replies []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
		} `json:"replies"`
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		VersionHistory []struct {
			Count int    `json:"count"`
			Href  string `json:"href"`
		} `json:"version-history"`
		WpAttachment []struct {
			Href string `json:"href"`
		} `json:"wp:attachment"`
		WpFeaturedmedia []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
		} `json:"wp:featuredmedia"`
		WpTerm []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
			Taxonomy   string `json:"taxonomy"`
		} `json:"wp:term"`
	} `json:"_links"`
	Template    string        `json:"template"`
	Title       struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Type string `json:"type"`
}

func ServePostList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := "https://www.hotelflow.com.br/wp-json/wp/v2/posts"

	// Creates a client connection.
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// Read the response.
	responseJSON, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var postsList []Posts

	// Parse the response.
	jsonErr := json.Unmarshal(responseJSON, &postsList)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	json.NewEncoder(w).Encode(postsList)
	//w.Write([] postsList)

}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/posts", ServePostList).Methods("GET")

	// Print it to the screen.
	http.ListenAndServe(":5000", router)



}