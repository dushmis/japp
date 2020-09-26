package app

import "appengine/datastore"
import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"
import "time"
import "strings"
import "reflect"
import "log"

type DonorService struct {
}

func (gs *DonorService) DonorsList(c endpoints.Context, r *DonorsListReq) (*DonorsList, error) {
	if r.Limit <= 0 {
		r.Limit = 10
	}

	q := datastore.NewQuery("Donor").Order("-Date").Limit(r.Limit)
	donors := make([]*Donor, 0, r.Limit)
	keys, err := q.GetAll(c, &donors)
	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		donors[i].Key = k
	}
	return &DonorsList{donors}, nil
}

func (gs *DonorService) DonorAdd(c endpoints.Context, g *Donor) error {
	k := datastore.NewIncompleteKey(c, "Donor", nil)
	g.Date = time.Now()
	_, err := datastore.Put(c, k, g)
	return err
}

// Count returns the number of greetings.
func (gs *DonorService) DonorsListCount(c endpoints.Context) (*Count, error) {
	n, err := datastore.NewQuery("Donor").Count(c)
	if err != nil {
		return nil, err
	}
	return &Count{n}, nil
}

func (fs *DonorService) _Register(name, version, apiname string) {
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
