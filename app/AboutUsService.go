package app

import "appengine/datastore"
import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"
import "time"
import "strings"
import "reflect"
import "log"

type AboutUsService struct {
}

func (gs *AboutUsService) AboutUsList(c endpoints.Context, r *AboutUsListReq) (*AboutUsList, error) {
	if r.Limit <= 0 {
		r.Limit = 10
	}

	q := datastore.NewQuery("AboutUs").Order("Email2").Limit(r.Limit)
	aboutus := make([]*AboutUs, 0, r.Limit)
	keys, err := q.GetAll(c, &aboutus)
	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		aboutus[i].Key = k
	}
	return &AboutUsList{aboutus}, nil
}

func (gs *AboutUsService) AboutUsAdd(c endpoints.Context, g *AboutUs) error {
	k := datastore.NewIncompleteKey(c, "AboutUs", nil)
	g.Date = time.Now()
	if g.Type == "" {
		g.Type = "team"
	}
	_, err := datastore.Put(c, k, g)
	return err
}

func (gs *AboutUsService) AboutUsDelete(c endpoints.Context, g *AboutUsDeleteReq) error {
	key, err := datastore.DecodeKey(g.Id)
	if err != nil {
		return err
	}
	err = datastore.Delete(c, key)
	if err != nil {
		return err
	}
	c.Infof("deleted %s", key)
	return nil
}

// Count returns the number of greetings.
func (gs *AboutUsService) AboutUsListCount(c endpoints.Context) (*Count, error) {
	n, err := datastore.NewQuery("AboutUs").Count(c)
	if err != nil {
		return nil, err
	}
	return &Count{n}, nil
}

func (fs *AboutUsService) _Register(name, version, apiname string) {
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
