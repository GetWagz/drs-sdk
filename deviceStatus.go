package drs

import (
	"net/http"
	"time"
)

// UpdateDeviceStatus updates the device status. According to the docs, you will want to call this at least once every 24 hours.
//
// If lastStatus is an empty string, we will replace it with the current timestamp in ISO8601
func UpdateDeviceStatus(deviceToken string, lastStatus string) (bool, error) {
	if deviceToken == "" {
		return false, &APIError{
			Code: http.StatusBadRequest,
			Data: map[string]string{
				"message": "deviceToken cannot be blank",
			},
		}
	}

	if lastStatus == "" {
		lastStatus = time.Now().Format(time.RFC3339)
	} else {
		//parse and make sure everything looks right
		_, timeErr := time.Parse(time.RFC3339, lastStatus)
		if timeErr != nil {
			return false, &APIError{
				Code: http.StatusBadRequest,
				Data: map[string]string{
					"message": "lastStatus is not valid",
				},
			}
		}
	}

	code, _, err := makeCall("deviceStatus", nil, deviceToken, map[string]string{
		"mostRecentlyActiveDate": lastStatus,
	})
	if err != nil || code != http.StatusOK {
		return false, err
	}

	return true, nil
}
