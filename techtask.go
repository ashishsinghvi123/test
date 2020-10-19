package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Meeting struct {
	Id                int           `json:"id"`
	Title             string        `json:"title"`
	Participants      []Participant `json:"participants"`
	StartTime         string        `json:"starttime"`
	EndTime           string        `json:"endtime "`
	CreationTimestamp string        `json:"creationtimestamp"`
}

type Participant struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	RSVP  bool   `json:"rsvp"`
}

type Meetings []Meeting

type Part []Participant

func allMeetings(w http.ResponseWriter, r *http.Request) {
	Party := Part{
		Participant{Name: "Ashish", Email: "ashish.singh2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Ruchit", Email: "ruchit.nigam2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Mansi", Email: "mansi.singh2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Sneha", Email: "sneha.singh2018@vitbhopal.ac.in", RSVP: true},
	}
	meetings := Meetings{
		Meeting{Id: 101, Title: "codevit1", Participants: Party, StartTime: "10:00am", EndTime: "11:00", CreationTimestamp: "9:45"},
		Meeting{Id: 102, Title: "codevit2", Participants: Party, StartTime: "11:04am", EndTime: "12:00", CreationTimestamp: "9:45"},
		Meeting{Id: 103, Title: "codevit3", Participants: Party, StartTime: "10:00am", EndTime: "11:00", CreationTimestamp: "9:45"},
		Meeting{Id: 104, Title: "codevit4", Participants: Party, StartTime: "11:04am", EndTime: "12:00", CreationTimestamp: "9:45"},
	}

	fmt.Println("end point hit all meetings")
	json.NewEncoder(w).Encode(meetings)

}

func allMeetings1(w http.ResponseWriter, r *http.Request) {
	Party := Part{
		Participant{Name: "Ashish", Email: "ashish.singh2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Ruchit", Email: "ruchit.nigam2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Mansi", Email: "mansi.singh2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Sneha", Email: "sneha.singh2018@vitbhopal.ac.in", RSVP: true},
	}

	meetings := Meetings{
		Meeting{Id: 101, Title: "codevit1", Participants: Party, StartTime: "10:00am", EndTime: "10:04", CreationTimestamp: "9:45"},
	}

	json.NewEncoder(w).Encode(meetings)
}
func allMeetings2(w http.ResponseWriter, r *http.Request) {
	Party := Part{
		Participant{Name: "Ashish", Email: "ashish.singh2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Ruchit", Email: "ruchit.nigam2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Mansi", Email: "mansi.singh2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Sneha", Email: "sneha.singh2018@vitbhopal.ac.in", RSVP: true},
	}

	meetings := Meetings{
		Meeting{Id: 102, Title: "codevit2", Participants: Party, StartTime: "10:00am", EndTime: "10:04", CreationTimestamp: "9:45"},
	}

	json.NewEncoder(w).Encode(meetings)
}

func allMeetings3(w http.ResponseWriter, r *http.Request) {
	Party := Part{
		Participant{Name: "Ashish", Email: "ashish.singh2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Ruchit", Email: "ruchit.nigam2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Mansi", Email: "mansi.singh2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Sneha", Email: "sneha.singh2018@vitbhopal.ac.in", RSVP: true},
	}

	meetings := Meetings{
		Meeting{Id: 103, Title: "codevit3", Participants: Party, StartTime: "10:00am", EndTime: "10:04", CreationTimestamp: "9:45"},
	}

	json.NewEncoder(w).Encode(meetings)
}

func allMeetings4(w http.ResponseWriter, r *http.Request) {
	Party := Part{
		Participant{Name: "Ashish", Email: "ashish.singh2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Ruchit", Email: "ruchit.nigam2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Mansi", Email: "mansi.singh2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Sneha", Email: "sneha.singh2018@vitbhopal.ac.in", RSVP: true},
	}

	meetings := Meetings{
		Meeting{Id: 104, Title: "codevit4", Participants: Party, StartTime: "10:00am", EndTime: "10:04", CreationTimestamp: "9:45"},
	}

	json.NewEncoder(w).Encode(meetings)
}
func allMeetingstime1(w http.ResponseWriter, r *http.Request) {
	Party := Part{
		Participant{Name: "Ashish", Email: "ashish.singh2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Ruchit", Email: "ruchit.nigam2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Mansi", Email: "mansi.singh2018@vitbhopal.ac.in", RSVP: true},
		Participant{Name: "Sneha", Email: "sneha.singh2018@vitbhopal.ac.in", RSVP: true},
	}

	meetings := Meetings{
		Meeting{Id: 101, Title: "codevit1", Participants: Party, StartTime: "10:00am", EndTime: "10:04", CreationTimestamp: "9:45"},
		Meeting{Id: 103, Title: "codevit3", Participants: Party, StartTime: "10:00am", EndTime: "10:04", CreationTimestamp: "9:45"},
	}

	json.NewEncoder(w).Encode(meetings)
}

func handleRequests() {

	http.HandleFunc("/meetings", allMeetings)
	http.HandleFunc("/meeting/101", allMeetings1)
	http.HandleFunc("/meeting/102", allMeetings2)
	http.HandleFunc("/meeting/103", allMeetings3)
	http.HandleFunc("/meeting/104", allMeetings4)
	http.HandleFunc("/meetings?start=10:00am & end=11.00", allMeetingstime1)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	handleRequests()

}
