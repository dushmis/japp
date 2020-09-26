package app

import "appengine/datastore"
import "time"

type JUser struct {
	Key      *datastore.Key `json:"id" datastore:"-"`
	Username string         `json:"username"`
	Password string         `json:"password"`
	Email    string         `json:"email"`
	Phone    string         `json:"phone"`
	Admin    bool           `json:"admin"`
	Date     time.Time      `json:"date"`
}

type JUsersList struct {
	Items []*JUser `json:"items"`
}

type JUsersListReq struct {
	Limit int `json:"limit" endpoints:"d=10"`
}

type JUsersAuth struct {
	Items *JUser `json:"items"`
	Auth  bool   `json:"auth"`
}

type JUsersAuthReq struct {
	//Limit int `json:"limit" endpoints:"d=10"`
	Username string `json:"username"`
	Password string `json:"password"`
}
