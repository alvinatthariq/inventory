package restserver

import (
	"net/http"
	"sync"

	"inventory/stock"

	"github.com/julienschmidt/httprouter"
)

var once = &sync.Once{}

type REST interface{}

type rest struct {
	router *httprouter.Router
	stock  stock.StockItf
}

func Init(
	router *httprouter.Router,
	stock stock.StockItf,
) REST {
	var e *rest
	once.Do(func() {
		e = &rest{
			router: router,
			stock:  stock,
		}
		e.Serve()
	})
	return e
}

func (e *rest) Serve() {
	// post stock
	e.router.POST("/stock", e.PostStock)

	// get stock by id
	e.router.GET("/stock/:id", e.GetStockByID)

	http.ListenAndServe(":8080", e.router)
}
