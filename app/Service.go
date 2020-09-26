package app

type Service struct {
	Name          string
	BroadCastName string
	Method        string
	Path          string
	Desc          string
}

var (
	service = []Service{
		Service{
			Name:          "FeedbacksList",
			BroadCastName: "feedback.list",
			Method:        "GET",
			Path:          "feedback",
			Desc:          "List feedbacks",
		}, Service{
			Name:          "FeedbackAdd",
			BroadCastName: "feedback.add",
			Method:        "PUT",
			Path:          "feedback",
			Desc:          "add new feedback",
		}, Service{
			Name:          "FeedbacksListCount",
			BroadCastName: "feedback.count",
			Method:        "GET",
			Path:          "feedback/count",
			Desc:          "feedback count",
		}, /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/

		Service{
			Name:          "CategoriesList",
			BroadCastName: "category.list",
			Method:        "GET",
			Path:          "categories",
			Desc:          "List categories",
		}, Service{
			Name:          "CategoryAdd",
			BroadCastName: "category.add",
			Method:        "PUT",
			Path:          "categories",
			Desc:          "add new category",
		}, Service{
			Name:          "CategoriesListCount",
			BroadCastName: "category.count",
			Method:        "GET",
			Path:          "categories/count",
			Desc:          "category count",
		}, /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/

		Service{
			Name:          "DonorsList",
			BroadCastName: "donor.list",
			Method:        "GET",
			Path:          "donors",
			Desc:          "List donors",
		}, Service{
			Name:          "DonorAdd",
			BroadCastName: "donor.add",
			Method:        "PUT",
			Path:          "donors",
			Desc:          "add new donor",
		}, Service{
			Name:          "DonorsListCount",
			BroadCastName: "donor.count",
			Method:        "GET",
			Path:          "donors/count",
			Desc:          "donor count",
		}, /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/

		Service{
			Name:          "EventsList",
			BroadCastName: "event.list",
			Method:        "GET",
			Path:          "events",
			Desc:          "List events",
		}, Service{
			Name:          "EventAdd",
			BroadCastName: "event.add",
			Method:        "PUT",
			Path:          "events",
			Desc:          "add new event",
		}, Service{
			Name:          "EventsListCount",
			BroadCastName: "event.count",
			Method:        "GET",
			Path:          "events/count",
			Desc:          "event count",
		}, /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/

		Service{
			Name:          "JUsersList",
			BroadCastName: "user.list",
			Method:        "GET",
			Path:          "users",
			Desc:          "List users",
		}, Service{
			Name:          "JUserAdd",
			BroadCastName: "user.add",
			Method:        "PUT",
			Path:          "users",
			Desc:          "add new user",
		}, Service{
			Name:          "JUsersListCount",
			BroadCastName: "user.count",
			Method:        "GET",
			Path:          "users/count",
			Desc:          "user count",
		}, Service{
			Name:          "JUserValidate",
			BroadCastName: "user.validate",
			Method:        "POST",
			Path:          "users/validate",
			Desc:          "user validation",
		}, /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/

		Service{
			Name:          "AboutUsList",
			BroadCastName: "aboutus.get",
			Method:        "GET",
			Path:          "aboutus",
			Desc:          "Get aboutus",
		}, Service{
			Name:          "AboutUsAdd",
			BroadCastName: "aboutus.add",
			Method:        "PUT",
			Path:          "aboutus",
			Desc:          "Add aboutus",
		}, Service{
			Name:          "AboutUsDelete",
			BroadCastName: "aboutus.delete",
			Method:        "DELETE",
			Path:          "aboutus",
			Desc:          "delete aboutus",
		}, /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/

		Service{
			Name:          "GalleriesList",
			BroadCastName: "gallery.list",
			Method:        "GET",
			Path:          "galleries",
			Desc:          "List galleries",
		}, Service{
			Name:          "GalleryAdd",
			BroadCastName: "gallery.add",
			Method:        "PUT",
			Path:          "galleries",
			Desc:          "add new gallery",
		}, Service{
			Name:          "GalleriesListCount",
			BroadCastName: "event.count",
			Method:        "GET",
			Path:          "galleries/count",
			Desc:          "gallery count",
		}, /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/

		Service{
			Name:          "EventMembersList",
			BroadCastName: "eventmembers.get",
			Method:        "GET",
			Path:          "eventmembers",
			Desc:          "Get eventmembers",
		}, Service{
			Name:          "EventMemberAdd",
			BroadCastName: "eventmembers.add",
			Method:        "PUT",
			Path:          "eventmembers",
			Desc:          "Add eventmembers",
		}, Service{
			Name:          "EventMembersListCount",
			BroadCastName: "eventmembers.count",
			Method:        "GET",
			Path:          "eventmembers/count",
			Desc:          "count eventmembers",
		}, /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/ /***/
	}
)

func _getService(name string) Service {
	for _, ser := range service {
		if name == ser.Name {
			return ser
		}
	}
	return Service{}
}
