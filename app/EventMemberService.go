package app

import "appengine/datastore"
import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"
import "time"
import "strings"
import "reflect"
import "log"

type EventMemberService struct {
}

func (gs *EventMemberService) EventMembersList(c endpoints.Context, r *EventMembersListReq) (*EventMembersList, error) {
	if r.Limit <= 0 {
		r.Limit = 10
	}

	q := datastore.NewQuery("EventMember").Order("-Date").Limit(r.Limit)
	events := make([]*EventMember, 0, r.Limit)
	keys, err := q.GetAll(c, &events)
	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		events[i].Key = k
	}
	return &EventMembersList{events}, nil
}

func (gs *EventMemberService) EventMemberAdd(c endpoints.Context, g *EventMember) error {
	k := datastore.NewIncompleteKey(c, "EventMember", nil)
	g.Date = time.Now()
	_, err := datastore.Put(c, k, g)
	return err
}

// Count returns the number of greetings.
func (gs *EventMemberService) EventMembersListCount(c endpoints.Context) (*Count, error) {
	n, err := datastore.NewQuery("EventMember").Count(c)
	if err != nil {
		return nil, err
	}
	return &Count{n}, nil
}

func (fs *EventMemberService) _Register(name, version, apiname string) {
	api, err := endpoints.RegisterService(fs,
		name, version, apiname, true)
	if err != nil {
		log.Fatalf("Register service: %v", err)
	}
	service_ := reflect.TypeOf(fs)
	for i := 0; i < service_.NumMethod(); i++ {
		method := service_.Method(i)
		if !strings.HasPrefix(method.Name, "_") {
			s := _getService(method.Name)
			if s.Name != "" {
				m := api.MethodByName(method.Name)
				i := m.Info()
				i.Name, i.HTTPMethod, i.Path, i.Desc = s.BroadCastName, s.Method, s.Path, s.Desc
			}
		}
	}

}
