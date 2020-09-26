package app

import "appengine/datastore"
import "time"

type Feedback struct {
	Key     *datastore.Key `json:"id" datastore:"-"`
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	Message string         `json:"message"`
	Phone   string         `json:"phone"`
	Date    time.Time      `json:"date"`
}

type FeedbacksList struct {
	Items []*Feedback `json:"items"`
}

type FeedbacksListReq struct {
	Limit int `json:"limit" endpoints:"d=10"`
}
