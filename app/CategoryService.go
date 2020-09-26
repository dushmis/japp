package app

import "appengine/datastore"
import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"
import "time"
import "strings"
import "reflect"
import "log"

type CategoryService struct {
}

func (gs *CategoryService) CategoriesList(c endpoints.Context, r *CategoriesListReq) (*CategoriesList, error) {
	if r.Limit <= 0 {
		r.Limit = 10
	}

	q := datastore.NewQuery("Category").Order("-Date").Limit(r.Limit)
	categories := make([]*Category, 0, r.Limit)
	keys, err := q.GetAll(c, &categories)
	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		categories[i].Key = k
	}
	return &CategoriesList{categories}, nil
}

func (gs *CategoryService) CategoryAdd(c endpoints.Context, g *Category) error {
	k := datastore.NewIncompleteKey(c, "Category", nil)
	g.Date = time.Now()
	_, err := datastore.Put(c, k, g)
	return err
}

// Count returns the number of greetings.
func (gs *CategoryService) CategoriesListCount(c endpoints.Context) (*Count, error) {
	n, err := datastore.NewQuery("Category").Count(c)
	if err != nil {
		return nil, err
	}
	return &Count{n}, nil
}

func (fs *CategoryService) _Register(name, version, apiname string) {
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
