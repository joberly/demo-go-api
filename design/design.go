package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("demo", func() {
	Title("Demo Service")
	Description("Demonstration Go Service")
	Server("demo", func() {
		Host("localhost", func() {
			URI("http://:8000")
		})
	})
	Error("invalid_arguments")
})

var RandResult = ResultType("application/vnd.demo.rand-result", func() {
	Description("Result of a rand operation.")
	TypeName("RandResult")
	Attributes(func() {
		Attribute("result", Int64)
	})
	View("default", func() {
		Attribute("result")
	})
})

var _ = Service("demo", func() {
	Description("Demonstration Go Service")

	Method("rand", func() {
		Payload(func() {
			Attribute("min", Int64)
			Attribute("max", Int64)
		})
		Result(RandResult)
		Error("invalid_arguments")
		HTTP(func() {
			GET("/rand")
			Param("min")
			Param("max")
			Response(StatusOK)
			Response("invalid_arguments", StatusBadRequest, func() {
				Description("Invalid arguments to API")
			})
		})
	})
})
