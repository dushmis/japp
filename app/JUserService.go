package app

import "appengine/datastore"
import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"
import "time"
import "strings"
import "reflect"
import "log"

type JUserService struct {
}

func (gs *JUserService) JUsersList(c endpoints.Context, r *JUsersListReq) (*JUsersList, error) {
	if r.Limit <= 0 {
		r.Limit = 10
	}

	q := datastore.NewQuery("JUser").Order("-Date").Limit(r.Limit)
	jusers := make([]*JUser, 0, r.Limit)
	keys, err := q.GetAll(c, &jusers)
	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		jusers[i].Key = k
		jusers[i].Password = "**"
	}
	return &JUsersList{jusers}, nil
}

func (gs *JUserService) JUserAdd(c endpoints.Context, g *JUser) error {
	k := datastore.NewIncompleteKey(c, "JUser", nil)
	g.Date = time.Now()
	_, err := datastore.Put(c, k, g)
	return err
}

// Count returns the number of greetings.
func (gs *JUserService) JUsersListCount(c endpoints.Context) (*Count, error) {
	n, err := datastore.NewQuery("JUser").Count(c)
	if err != nil {
		return nil, err
	}
	return &Count{n}, nil
}

func (gs *JUserService) JUserValidate(c endpoints.Context, r *JUsersAuthReq) (*JUsersAuth, error) {
	if r.Username == "" || r.Password == "" {
		return &JUsersAuth{Auth: false}, nil
	}

	q := datastore.NewQuery("JUser").Filter("Username  =", r.Username).Filter("Password =", r.Password).Limit(1)
	juser := &JUser{}
	jusers := make([]*JUser, 0, 1)
	keys, err := q.GetAll(c, &jusers)
	if err != nil {
		return nil, err
	}
	if len(keys) > 0 {
		juser = jusers[0]
		for _, k := range keys {
			juser.Key = k
			// juser.Email = string(k.IntID())
			juser.Password = "**"
		}
		return &JUsersAuth{Items: juser, Auth: true}, nil
	} else {
		return &JUsersAuth{Auth: false}, nil
	}
}

func (fs *JUserService) _Register(name, version, apiname string) {
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
