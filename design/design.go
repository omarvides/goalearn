package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("user", func() {
	Title("User access API")
	Description("A users management system API for any resource")
	Scheme("http")
	Host("localhost:8081")
})

var _ = Resource("user", func() {
	BasePath("/users")
	DefaultMedia(UserMedia)

	Action("show", func() {
		Description("Get user by id")
		Routing(GET("/:userID"))
		Params(func() {
			Param("userID", Integer, "User ID")
		})
		Response(OK)
		Response(NotFound)
	})
})

// UserMedia defines the media type for users
var UserMedia = MediaType("application/user.api+json", func() {
	Description("A user")
	Attributes(func() {
		Attribute("id", Integer, "Unique user ID")
		Attribute("age", Integer, "Age of the user")
		Attribute("name", String, "The name of the user")
		Required("id", "age", "name")
	})
	View("default", func() {
		Attribute("id")
		Attribute("age")
		Attribute("name")
	})
})
