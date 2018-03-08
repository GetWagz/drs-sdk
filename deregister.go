package drs

import "net/http"

// DeregisterDevice sends a request to DRS requesting that the device is deregistered from the service
func DeregisterDevice(deviceToken string) (bool, *APIError) {
	if deviceToken == "" {
		err := APIError{
			Code: http.StatusBadRequest,
			Data: map[string]string{
				"message": "deviceToken cannot be blank",
			},
		}
		return false, &err
	}

	code, _, err := makeCall("deregister", nil, deviceToken, map[string]string{})
	if err != nil || code != http.StatusOK {
		return false, err
	}

	return true, nil
}
