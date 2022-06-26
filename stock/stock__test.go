package stock_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"inventory/entity"
	"inventory/stock"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	sqlClient *sql.DB
	err       error

	stockItf stock.StockItf
)

func TestMain(t *testing.M) {
	// init sql connection
	sqlClient, err = sql.Open("mysql", "root:password@tcp(127.0.0.1)/inventory_db")
	if err != nil {
		log.Fatalln(err)
	}

	stockItf = stock.InitStock(sqlClient)

	exitVal := t.Run()

	sqlClient.Close()

	os.Exit(exitVal)
}

func TestStock(t *testing.T) {
	stocks := testCreateStock(t)
	if len(stocks) > 0 {
		testGetStockByID(t, stocks[0])
	}
}

func testCreateStock(t *testing.T) []entity.Stock {
	stocks := []entity.Stock{}
	t.Run("TestCreateStock", func(t *testing.T) {
		Convey("TestCreateStock", t, FailureHalts, func() {
			testCases := []struct {
				testID   int
				testType string
				testDesc string
				args     struct {
					payload entity.CreateStockRequest
				}
			}{
				{

					testID:   1,
					testDesc: "success",
					testType: "P",
					args: struct{ payload entity.CreateStockRequest }{
						payload: entity.CreateStockRequest{
							Stocks: []entity.CreateStock{
								{
									Name:         "test1",
									Availability: 10,
									Price:        1000,
									IsActive:     true,
								},
							},
						},
					},
				},
				{

					testID:   2,
					testDesc: "error, price is 0",
					testType: "N",
					args: struct{ payload entity.CreateStockRequest }{
						payload: entity.CreateStockRequest{
							Stocks: []entity.CreateStock{
								{
									Name:         "test1",
									Availability: 10,
									Price:        0,
									IsActive:     true,
								},
							},
						},
					},
				},
				{

					testID:   3,
					testDesc: "error, availability is negative",
					testType: "N",
					args: struct{ payload entity.CreateStockRequest }{
						payload: entity.CreateStockRequest{
							[]entity.CreateStock{
								{
									Name:         "test1",
									Availability: -1,
									Price:        1000,
									IsActive:     true,
								},
							},
						},
					},
				},
			}

			for _, tc := range testCases {
				t.Logf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc)
				results, err := stockItf.CreateStock(tc.args.payload)
				if tc.testType == "P" {
					So(err, ShouldBeNil)

					for i := range results {
						So(results[i].ID, ShouldHaveLength, 36)
						So(results[i].Name, ShouldEqual, tc.args.payload.Stocks[i].Name)
						So(results[i].Availability, ShouldEqual, tc.args.payload.Stocks[i].Availability)
						So(results[i].Price, ShouldEqual, tc.args.payload.Stocks[i].Price)
						So(results[i].IsActive, ShouldEqual, tc.args.payload.Stocks[i].IsActive)
					}

					stocks = append(stocks, results...)
				} else {
					So(err, ShouldNotBeNil)
				}
			}
		})
	})

	return stocks
}

func testGetStockByID(t *testing.T, stockData entity.Stock) {
	t.Run("TestGetStockByID", func(t *testing.T) {
		Convey("TestGetStockByID", t, FailureHalts, func() {
			testCases := []struct {
				testID   int
				testType string
				testDesc string
				args     struct {
					id string
				}
			}{
				{

					testID:   1,
					testDesc: "success get valid data",
					testType: "P",
					args: struct{ id string }{
						id: stockData.ID,
					},
				},
				{

					testID:   2,
					testDesc: "error, stock not found",
					testType: "N",
					args: struct{ id string }{
						id: "invaliddd",
					},
				},
			}

			for _, tc := range testCases {
				t.Logf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc)
				result, err := stockItf.GetStockByID(tc.args.id)
				if tc.testType == "P" {
					So(err, ShouldBeNil)
					So(result, ShouldResemble, stockData)
				} else {
					So(err, ShouldNotBeNil)
				}
			}
		})
	})
}
