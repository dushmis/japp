package app

import "appengine/datastore"
import "time"

type Category struct {
	Key           *datastore.Key `json:"id" datastore:"-"`
	Title         string         `json:"title"`
	Icon          string         `json:"icon"`
	Description   string         `json:"description"`
	Image1        string         `json:"image1"`
	Image2        string         `json:"image2"`
	Subcategory   Subcategory    `json:"subcategory"`
	LoginRequired bool           `json:"loginRequired"`
	AdminRequired bool           `json:"adminRequired"`
	Date          time.Time      `json:"date"`
}

type Subcategory struct {
	Title       string `json:"title"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Image1      string `json:"image1"`
	Image2      string `json:"image2"`
}

type CategoriesList struct {
	Items []*Category `json:"items"`
}

type CategoriesListReq struct {
	Limit int `json:"limit" endpoints:"d=10"`
}
