package app

import "appengine/datastore"

import "time"

type AboutUs struct {
	Key         *datastore.Key `json:"id" datastore:"-"`
	Name        string         `json:"name"`
	City        string         `json:"city"`
	State       string         `json:"state"`
	Image1      string         `json:"image1"`
	Detail      string         `json:"detail"`
	Phone1      string         `json:"phone1"`
	Phone2      string         `json:"phone2"`
	Email1      string         `json:"email1"`
	Email2      string         `json:"email2"`
	Description string         `json:"description"`
	Type        string         `json:"type"`
	Date        time.Time      `json:"date"`
}

type AboutUsList struct {
	Items []*AboutUs `json:"items"`
}

type AboutUsListReq struct {
	Limit int `json:"limit" endpoints:"d=10"`
}

type AboutUsDeleteReq struct {
	Id string `json:"id" endpoints:"d=''"`
}
