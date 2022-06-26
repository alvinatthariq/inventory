package restserver

import (
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

var once = &sync.Once{}

type REST interface{}

type rest struct {
	router *httprouter.Router
}

func Init(router *httprouter.Router) REST {
	var e *rest
	once.Do(func() {
		e = &rest{
			router: router,
		}
		e.Serve()
	})
	return e
}

func (e *rest) Serve() {
	e.router.POST("/stock", PostStock)

	e.router.GET("/stock/:id", GetStockByID)

	http.ListenAndServe(":8080", e.router)
}
