package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	Name     string    `json:"slack_name"`
	Day      string    `json:"current_day"`
	Time     time.Time `json:"utc_time"`
	Track    string    `json:"track"`
	File_url string    `json:"github_file_url"`
	Repo_url string    `json:"github_rep_url"`
	Status   int       `json:"status_code"`
}

var user = User{}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//Parse query parameter
		name := r.URL.Query().Get("slack_name")
		track := r.URL.Query().Get("track")

		//Checking if name and track parameters are provided
		if name == "" || track == "" {
			http.Error(w, "Both 'slack_name' and 'track' Parameters are required", http.StatusBadRequest)
			return
		}

		//Handling current time and day of the week
		t := time.Now().UTC()
		tDay, tTime := t.Weekday(), t

		//Formatting the return integer as weekday
		currentDay := dayFormat(tDay)

		//Handling github url
		source_url := "https://github.com/Zekeriyyah/hngx/blob/main/stageone/main.go"
		repo_url := "https://github.com/Zekeriyyah/hngx"

		user = User{
			Name:     name,
			Day:      currentDay,
			Time:     tTime,
			Track:    track,
			File_url: source_url,
			Repo_url: repo_url,
			Status:   http.StatusOK,
		}

		//Serializing the user struct to JSON
		res, err := json.Marshal(&user)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		//Setting the header and writing the JSON data to the response writer
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(res)
		if err != nil {
			http.Error(w, "Error writing response.", http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func dayFormat(t time.Weekday) string {
	switch t {
	case time.Sunday:
		return "Sunday"
	case time.Monday:
		return "Monday"
	case time.Tuesday:
		return "Tuesday"
	case time.Wednesday:
		return "Wednesday"
	case time.Thursday:
		return "Thursday"
	case time.Friday:
		return "Friday"
	case time.Saturday:
		return "Saturday"
	default:
		return "Invalid Day of the Week"
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9070"
	}

	http.HandleFunc("/api", UserHandler)

	fmt.Printf("Server is running on port %s....\n", port)
	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
