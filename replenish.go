package drs

import (
	"net/http"

	"github.com/mitchellh/mapstructure"
)

/*ReplenishResult represents the end result of a request to replenish a slot
 */
type ReplenishResult struct {
	EventInstanceID string `json:"eventInstanceId"`
	DetailCode      string `json:"detailCode"`
}

/*ReplenishSlot asks Amazon to submit an order to replenish a specific slot for the device
 */
func ReplenishSlot(deviceToken, slotID string) (*ReplenishResult, *APIError) {
	if deviceToken == "" {
		return nil, &APIError{
			Code: http.StatusBadRequest,
			Data: map[string]string{
				"message": "deviceToken cannot be blank",
			},
		}
	}

	code, body, err := makeCall("replenishSlot", []interface{}{slotID}, deviceToken, map[string]string{})
	if err != nil || code != http.StatusOK {
		return nil, err
	}

	ret := ReplenishResult{}
	repErr := mapstructure.Decode(body, &ret)
	if repErr != nil {
		return nil, &APIError{
			Code: http.StatusInternalServerError,
			Data: map[string]string{
				"message": "could not parse SDK response",
			},
		}
	}
	return &ret, nil
}
