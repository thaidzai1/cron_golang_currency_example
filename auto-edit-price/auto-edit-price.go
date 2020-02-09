package editpricejobs

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	httpreq "gido.vn/gic/cron/auto-edit-price/http-request"
	cm "gido.vn/gic/cron/auto-edit-price/models"

	"gido.vn/gic/cron/auto-edit-price/models/admin"
	"gido.vn/gic/cron/auto-edit-price/models/pship"
)

var (
	totalBoxedParcel  int64 = 0
	totalUpdatedOrder int64 = 0
	adminHost               = os.Getenv("HOST")
)

// Main ...
func Main() {
	fmt.Println("Ahoy!!!")
	adminInfo, err := adminLogin()
	panic(err, "Admin login failed")
	fmt.Println(adminInfo)

	var fetchingMinute time.Duration
	if os.Getenv("CRON_SCHEDULE") == "" {
		fetchingMinute = 15
	} else {
		parsedTime, err := strconv.ParseInt(os.Getenv("CRON_SCHEDULE"), 10, 64)
		if err != nil {
			fmt.Printf("parse time failed: %v", err)
		}
		fetchingMinute = time.Duration(parsedTime)
	}
	boxedParcels, err := getBoxedParcelByTime(fetchingMinute, 0, 0, adminInfo.AccessToken)
	panic(err, "Query boxed parcels by time failed")
	if boxedParcels.PageInfo.Total == 0 {
		return
	}

	boxedParcelsChan := make(chan int64)
	stopBoxedParcelsChan := make(chan bool)
	updatedOrderChan := make(chan bool)
	defer close(boxedParcelsChan)
	defer close(stopBoxedParcelsChan)
	defer close(updatedOrderChan)

	go processGetBoxedParcel(boxedParcelsChan, stopBoxedParcelsChan, adminInfo.AccessToken)
	processEditBoxedParcel(boxedParcelsChan, updatedOrderChan, stopBoxedParcelsChan, adminInfo.AccessToken)

	for {
		switch {
		case <-updatedOrderChan:
			totalUpdatedOrder++
			if totalUpdatedOrder == totalBoxedParcel {
				fmt.Println("DONE †")
				return
			}
		}
	}
}

func panic(err error, message string) {
	if err != nil {
		fmt.Printf("Panic: %v, %v\n", err, message)
		os.Exit(0)
	}
}

func handleError(err error, message string) {
	if err != nil {
		fmt.Printf("Error: %v: %v", message, err)
	}
}

func adminLogin() (*admin.Admin, error) {
	loginAPI := httpreq.API{
		Host:   adminHost,
		URI:    "/account/login",
		Method: "POST",
		Body: &admin.LoginRequest{
			Phone:    os.Getenv("ADMIN_USERNAME"),
			Password: os.Getenv("ADMIN_PASSWORD"),
		},
	}

	jsonLoginResponse, err := loginAPI.RequestAPI()
	if err != nil {
		return nil, err
	}
	loginResponse := &admin.Admin{}
	err = json.Unmarshal(jsonLoginResponse, loginResponse)
	if err != nil {
		return nil, err
	}

	return loginResponse, nil
}

func getBoxedParcelByTime(minute time.Duration, offet int64, size int64, token string) (*pship.ListOrderItemsResponse, error) {
	now := time.Now().UTC()
	cielCheckTime := now.UnixNano() / 1000000
	floorCheckTime := now.Add(-time.Second*60*minute).UnixNano() / 1000000

	getBoxedParcelAPI := httpreq.API{
		Host:   adminHost,
		URI:    "/boxed_parcels",
		Method: "GET",
		Query: httpreq.Query{
			Paging: cm.Paging{
				First:  size,
				Offset: offet,
			},
			Filter: cm.Filter{
				UpdatedAt: strconv.FormatInt(cielCheckTime, 10) + "," + strconv.FormatInt(floorCheckTime, 10),
			},
		},
		Token: token,
	}

	jsonGetBoxedParcelResponse, err := getBoxedParcelAPI.RequestAPI()
	if err != nil {
		handleError(err, "RequestAPI get boxed parcels failed")
		return nil, err
	}
	fmt.Println(string(jsonGetBoxedParcelResponse))

	boxedParcelResponse := &pship.ListOrderItemsResponse{}
	err = json.Unmarshal(jsonGetBoxedParcelResponse, boxedParcelResponse)
	if err != nil {
		handleError(err, "Decode json to struct OrderItem failed")
		return nil, err
	}

	return boxedParcelResponse, nil
}

func getMerchantOrderInfo(id int64, token string) (*pship.MerchantOrdersResponse, error) {
	getMerchantOrderAPI := httpreq.API{
		Host:   adminHost,
		URI:    "/merchant_order/id/" + strconv.FormatInt(id, 10),
		Method: "GET",
		Token:  token,
	}

	jsonGetMerchantOrderAPI, err := getMerchantOrderAPI.RequestAPI()
	if err != nil {
		handleError(err, "RequestAPI get merchant order failed")
		return nil, err
	}
	fmt.Println(string(jsonGetMerchantOrderAPI))

	merchantOrders := &pship.MerchantOrdersResponse{}
	err = json.Unmarshal(jsonGetMerchantOrderAPI, merchantOrders)
	if err != nil {
		handleError(err, "Decode json to struct MerchantOrdersReponse failed")
		return nil, err
	}

	return merchantOrders, nil
}

func editPriceMerchantOrder(merchantOrderID int64, token string, updatedOrderChan chan bool) {
	editPriceAPI := httpreq.API{
		Host:   adminHost,
		URI:    "/order_price/edit",
		Method: "POST",
		Body: &pship.EditOrderPriceRequest{
			ID: merchantOrderID,
		},
		Token: token,
	}

	jsonEditPriceResponse, err := editPriceAPI.RequestAPI()
	handleError(err, "Request edit price api failed")
	fmt.Printf("editedPrice: OrderID: %v\n", merchantOrderID)
	fmt.Println(string(jsonEditPriceResponse))

	editPriceResponse := &pship.OrderPriceResponse{}
	err = json.Unmarshal(jsonEditPriceResponse, editPriceResponse)
	handleError(err, "Decoding json failed")
	updatedOrderChan <- true
}

func processGetBoxedParcel(boxedParcelsChan chan int64, quit chan bool, token string) {
	var fetchingMinute time.Duration
	if os.Getenv("CRON_SCHEDULE") == "" {
		fetchingMinute = 15
	} else {
		parsedTime, err := strconv.ParseInt(os.Getenv("CRON_SCHEDULE"), 10, 64)
		if err != nil {
			fmt.Printf("parse time failed: %v", err)
		}
		fetchingMinute = time.Duration(parsedTime)
	}
	var offset int64 = 0
	var size int64 = 10
	for {
		boxedParcels, err := getBoxedParcelByTime(fetchingMinute, offset, size, token)
		fmt.Printf("Number of boxed parcels: %v\n", len(boxedParcels.OrderItems))
		panic(err, "get boxed parcels failed")

		if offset == 0 && totalBoxedParcel == 0 {
			totalBoxedParcel = boxedParcels.PageInfo.Total
		}

		for _, boxedParcel := range boxedParcels.OrderItems {
			merchantOrders, err := getMerchantOrderInfo(boxedParcel.MerchantOrderID, token)
			if err != nil {
				handleError(err, "get merchant order info failed")
				continue
			}
			if !merchantOrders.MerchantOrder[0].IsCBE {
				boxedParcelsChan <- boxedParcel.MerchantOrderID
			}
		}
		if boxedParcels.PageInfo.Next {
			offset += size
		} else {
			break
		}
	}
	quit <- true
}

func processEditBoxedParcel(boxedParcelChan chan int64, updatedOrderChan chan bool, quit chan bool, token string) {
	for {
		select {
		case merchantOrderID := <-boxedParcelChan:
			fmt.Printf("MerchantOrderID: %v\n", merchantOrderID)
			go editPriceMerchantOrder(merchantOrderID, token, updatedOrderChan)
		case <-quit:
			fmt.Println("GET FULLY BOXED PARCELS DONE †")
			return
		}
	}
}
