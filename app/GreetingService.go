package app

import "appengine/datastore"
import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"

type GreetingService struct {
}

func (gs *GreetingService) List(c endpoints.Context, r *GreetingsListReq) (*GreetingsList, error) {
	if r.Limit <= 0 {
		r.Limit = 10
	}

	q := datastore.NewQuery("Greeting").Order("-Date").Limit(r.Limit)
	greets := make([]*Greeting, 0, r.Limit)
	keys, err := q.GetAll(c, &greets)
	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		greets[i].Key = k
	}
	return &GreetingsList{greets}, nil
}

func (gs *GreetingService) Add(c endpoints.Context, g *Greeting) error {
	k := datastore.NewIncompleteKey(c, "Greeting", nil)
	_, err := datastore.Put(c, k, g)
	return err
}

type Count struct {
	N int `json:"count"`
}

// Count returns the number of greetings.
func (gs *GreetingService) Count(c endpoints.Context) (*Count, error) {
	n, err := datastore.NewQuery("Greeting").Count(c)
	if err != nil {
		return nil, err
	}
	return &Count{n}, nil
}
