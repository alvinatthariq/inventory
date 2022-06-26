package restserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"inventory/entity"

	"github.com/julienschmidt/httprouter"
)

func (e *rest) PostStock(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var requestBody entity.CreateStock
	err := decoder.Decode(&requestBody)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{
			"error": %s
		}`, err.Error())))
		return
	}

	stockData, err := e.stock.CreateStock(requestBody)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{
			"error": "%s"
		}`, err.Error())))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{
		"id": "%s"
	}`, stockData.ID)))
}

func (e *rest) GetStockByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	stockID := params.ByName("id")
	if len(stockID) < 1 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{
			"error": "must specify id"
		}`))
		return
	}

	stockData, err := e.stock.GetStockByID(stockID)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		if err.Error() == "stock not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(fmt.Sprintf(`{
			"error": "%s"
		}`, err.Error())))

		return
	}

	responseBody, err := json.Marshal(stockData)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{
			"error": "%s"
		}`, err.Error())))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
