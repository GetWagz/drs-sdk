package drs

// DeregisterDevice sends a request to DRS requesting that the device is deregistered from the service
func DeregisterDevice(deviceToken string) (bool, *APIError) {
	if deviceToken == "" {
		err := APIError{
			Code: 400,
			Data: map[string]string{
				"message": "deviceToken cannot be blank",
			},
		}
		return false, &err
	}

	code, _, err := makeCall("DELETE", "deregister", nil, deviceToken, map[string]string{})
	if err != nil {
		return false, err
	}
	if code != 200 {
		return false, err
	}

	return true, nil
}
