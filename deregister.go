package drs

import "net/http"

// DeregisterDevice sends a request to DRS requesting that the device is deregistered from the service
func DeregisterDevice(deviceToken string) (bool, error) {
	if deviceToken == "" {
		return false, &APIError{
			Code: http.StatusBadRequest,
			Data: map[string]string{
				"message": "deviceToken cannot be blank",
			},
		}
	}

	code, _, err := makeCall("deregister", nil, deviceToken, map[string]string{})
	if err != nil || code != http.StatusOK {
		return false, err
	}

	return true, nil
}
