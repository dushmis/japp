package app

import "appengine/datastore"
import "time"

type EventMember struct {
	Key     *datastore.Key `json:"id" datastore:"-"`
	Name    string         `json:"name"`
	City    string         `json:"city"`
	State   string         `json:"state"`
	Event   string         `json:"event"`
	Address string         `json:"address"`
	Email   string         `json:"email"`
	Remarks string         `json:"remarks"`
	Details string         `json:"details"`
	Date    time.Time      `json:"date"`
}

type EventMembersList struct {
	Items []*EventMember `json:"items"`
}

type EventMembersListReq struct {
	Limit int `json:"limit" endpoints:"d=10"`
}
