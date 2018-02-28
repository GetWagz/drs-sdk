package drs

import (
	"time"
)

// UpdateDeviceStatus updates the device status. According to the docs, you will want to call this at least once every 24 hours.
//
// If lastStatus is an empty string, we will replace it with the current timestamp in ISO8601
func UpdateDeviceStatus(deviceToken string, lastStatus string) (bool, *APIError) {
	if deviceToken == "" {
		err := APIError{
			Code: 400,
			Data: map[string]string{
				"message": "deviceToken cannot be blank",
			},
		}
		return false, &err
	}

	if lastStatus == "" {
		lastStatus = time.Now().Format(time.RFC3339)
	} else {
		//parse and make sure everything looks right
		_, timeErr := time.Parse(time.RFC3339, lastStatus)
		if timeErr != nil {
			err := APIError{
				Code: 400,
				Data: map[string]string{
					"message": "lastStatus is not valid",
				},
			}
			return false, &err
		}
	}

	code, _, err := makeCall("POST", "deviceStatus", deviceToken, map[string]string{
		"mostRecentlyActiveDate": lastStatus,
	})
	if err != nil {
		return false, err
	}
	if code != 200 {
		return false, err
	}

	return true, nil
}
