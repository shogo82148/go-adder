package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("adder", func() {
	Title("The adder API")
	Description("goaで遊んでみるテスト")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("operands", func() {
	Action("add", func() {
		Routing(GET("add/:left/:right"))
		Description("leftとrightを足した値をレスポンスボディーに入れて返します")
		Params(func() {
			Param("left", Integer, "左値")
			Param("right", Integer, "右値")
		})
		Response(OK, "text/plain")
	})
})
