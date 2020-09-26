package app

import "appengine/datastore"
import "time"

type Donor struct {
	Key    *datastore.Key `json:"id" datastore:"-"`
	Name   string         `json:"name"`
	City   string         `json:"city"`
	Amount int            `json:"amount"`
	Mobile string         `json:"mobile"`
	Date   time.Time      `json:"date"`
}

type DonorsList struct {
	Items []*Donor `json:"items"`
}

type DonorsListReq struct {
	Limit int `json:"limit" endpoints:"d=10"`
}
