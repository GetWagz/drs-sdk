package drs

// DeregisterUser sends a request to DRS requesting that the user is deregistered from the service
func DeregisterUser(userToken string) (bool, *APIError) {
	if userToken == "" {
		err := APIError{
			Code: 400,
			Data: map[string]string{
				"message": "userToken cannot be blank",
			},
		}
		return false, &err
	}

	code, _, err := makeCall("DELETE", "deregister", userToken, map[string]string{})
	if err != nil {
		return false, err
	}
	if code != 200 {
		return false, err
	}

	return true, nil
}
