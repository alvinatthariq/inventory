package restserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"inventory/entity"
	"inventory/stock"

	"github.com/julienschmidt/httprouter"
)

func PostStock(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	stockData, err := stock.CreateStock(requestBody)
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

func GetStockByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	stockID := params.ByName("id")
	if len(stockID) < 1 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{
			"error": "must specify id"
		}`))
		return
	}

	stockData, err := stock.GetStockByID(stockID)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{
			"error": "%s"
		}`, err.Error())))
		return
	}

	if (stockData == entity.Stock{}) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{
			"error": "stock not found"
		}`))
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
