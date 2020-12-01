# autodoc
to generate document automatically

### test
```
package main

import (
	"github.com/ailose/autodoc"
	"log"
)


func main() {
	sec := make(map[string]DocSection)
	sec["test"] = DocSection{
		Index: 0,
		Name:  "test",
	}
	var infos []*DocInfo
	info := DocInfo{
		Name: "test",
		Desc: "test des",
	}
	infos = append(infos, &info)
	GenDoc("test.md", "test", sec, infos)
}

```

### example

use on gin:

type Server struct {
	docs              []*autodoc.DocInfo
}

type Handler struct {
	handler gin.HandlerFunc
	doc     *autodoc.DocInfo
}

r.setupRoute(gr, dr, "GET", "/ecs", 0, s.TestHandler())

func (s *Server) setupRoute(router, docRouter *gin.RouterGroup, method, uri string, roles uint32, handler Handler) {
	router.Handle(method, uri, handler.handler...)
	s.setupRouteDoc(docRouter, method, router.BasePath()+uri, roles, handler.doc)
}

func (s *Server) setupRouteDoc(router *gin.RouterGroup, method, uri string, roles uint32, info *autodoc.DocInfo) {
	if info != nil {
		info.Roles = roles
		info.Method = method
		info.URI = uri
		//log.Println(info)
		s.docs = append(s.docs, info)

		if router != nil {
			router.GET(uri, func(c *gin.Context) {
				c.JSON(http.StatusOK, info.Req)
				c.JSON(http.StatusOK, info.Rsp)
			})
		}
	}
}

