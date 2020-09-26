package app

import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"

func init() {
	categotiesService := &CategoryService{}
	donorService := &DonorService{}
	eventService := &EventService{}
	feedbackService := &FeedbackService{}
	juserService := &JUserService{}
	aboutUsService := &AboutUsService{}
	galleryService := &GalleryService{}
	eventMemberService := &EventMemberService{}

	feedbackService._Register("feedback", "v2", "Feedback API")
	categotiesService._Register("category", "v2", "Category API")
	donorService._Register("donor", "v2", "Donor API")
	eventService._Register("event", "v2", "Event API")
	juserService._Register("user", "v2", "User API")
	aboutUsService._Register("aboutus", "v2", "AboutUs API")
	galleryService._Register("gallery", "v2", "Gallery API")
	eventMemberService._Register("eventMember", "v2", "Event Member API")

	endpoints.HandleHTTP()
}
