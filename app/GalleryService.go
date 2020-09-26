package app

import "appengine/datastore"
import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"
import "time"
import "strings"
import "reflect"
import "log"

type GalleryService struct {
}

func (gs *GalleryService) GalleriesList(c endpoints.Context, r *GalleriesListReq) (*GalleriesList, error) {
	if r.Limit <= 0 {
		r.Limit = 10
	}

	q := datastore.NewQuery("Gallery").Order("-Date").Limit(r.Limit)
	galleries := make([]*Gallery, 0, r.Limit)
	keys, err := q.GetAll(c, &galleries)
	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		galleries[i].Key = k
	}
	return &GalleriesList{galleries}, nil
}

func (gs *GalleryService) GalleryAdd(c endpoints.Context, g *Gallery) error {
	k := datastore.NewIncompleteKey(c, "Gallery", nil)
	g.Date = time.Now()
	_, err := datastore.Put(c, k, g)
	return err
}

// Count returns the number of greetings.
func (gs *GalleryService) GalleriesListCount(c endpoints.Context) (*Count, error) {
	n, err := datastore.NewQuery("Gallery").Count(c)
	if err != nil {
		return nil, err
	}
	return &Count{n}, nil
}

func (fs *GalleryService) _Register(name, version, apiname string) {
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
