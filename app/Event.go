package app

import "appengine/datastore"
import "time"

type Event struct {
	Key     *datastore.Key `json:"id" datastore:"-"`
	Title   string         `json:"title"`
	City    string         `json:"city"`
	Venue   string         `json:"venue"`
	Details string         `json:"details"`
	Date    time.Time      `json:"date"`
}

type EventsList struct {
	Items []*Event `json:"items"`
}

type EventsListReq struct {
	Limit int `json:"limit" endpoints:"d=10"`
}
