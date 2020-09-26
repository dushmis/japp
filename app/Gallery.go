package app

import "appengine/datastore"
import "time"

type Gallery struct {
	Key    *datastore.Key `json:"id" datastore:"-"`
	Title  string         `json:"title"`
	Path   string         `json:"path"`
	File   string         `json:"file"`
	Name   string         `json:"name"`
	Detail string         `json:"detail"`
	Date   time.Time      `json:"date"`
}

type GalleriesList struct {
	Items []*Gallery `json:"items"`
}

type GalleriesListReq struct {
	Limit int `json:"limit" endpoints:"d=10"`
}
