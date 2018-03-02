package drs

import (
	"github.com/mitchellh/mapstructure"
)

/*SlotOrderStatus represents a single slot and the status of order on that slot*/
type SlotOrderStatus struct {
	OrderStatus string `json:"orderStatus"`
	SlotID      string `json:"slotId"`
}

/*SlotOrderStatuses is a container to hold a slice of SlotOrderStatuses returned after cancelling a test order*/
type SlotOrderStatuses struct {
	SlotOrderStatuses []SlotOrderStatus `json:"slotOrderStatuses"`
}

/*CancelTestOrder cancels a test order on a slot for a device. Note that only test orders can be cancelled. An order for a real device must be cancelled through the Amazon account of the user and is not exposed by the DRS API*/
func CancelTestOrder(deviceToken, slotID string) (*SlotOrderStatuses, error) {
	if deviceToken == "" || slotID == "" {
		return nil, &APIError{
			Code: 400,
			Data: map[string]string{
				"message": "deviceToken and slotID cannot be blank",
			},
		}
	}

	code, resp, err := makeCall("cancelTestOrder", []interface{}{slotID}, deviceToken, map[string]string{})
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, err
	}

	result := SlotOrderStatuses{}
	decodeErr := mapstructure.Decode(resp, &result)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return &result, nil
}

type OrderItem struct {
	ASIN                 string `json:"asin"`
	ExpectedDeliveryDate string `json:"expectedDeliveryDate"`
	Quantity             int    `json:"quantity"`
	SlotID               string `json:"slotID"`
	Status               string `json:"status"`
}

type OrderInfoData struct {
	InstanceID string      `json:"instanceId"`
	OrderItems []OrderItem `json:"orderItems"`
}

func GetOrderInfo(deviceToken, instanceID string) (*OrderInfoData, error) {
	if deviceToken == "" || instanceID == "" {
		return nil, &APIError{
			Code: 400,
			Data: map[string]string{
				"message": "deviceToken and instanceID cannot be blank",
			},
		}
	}
	code, resp, err := makeCall("getOrderInfo", []interface{}{instanceID}, deviceToken, map[string]string{})
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, err
	}

	result := OrderInfoData{}
	oid, oidOK := resp["orderInfoData"]
	if !oidOK || oid == "" {
		return nil, &APIError{
			Code: 500,
			Data: map[string]string{
				"message": "no orderInfoData was returned from Amazon",
			},
		}
	}
	decodeErr := mapstructure.Decode(oid, &result)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return &result, nil
}
