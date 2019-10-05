package goa_files

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("files", func() {
	Title("File Management Service")

	Server("files", func() {
		Services("files")

		Host("development", func() {
			URI("http://localhost:8000/files")
			//URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("files", func( ){
	Method("add", func() {
		Payload(func() {
			Attribute("file", String)
		})

		HTTP(func() {
			POST("/")
			//MultipartRequest()
		})
	})
})
