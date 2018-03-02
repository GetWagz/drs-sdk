package drs

import (
	"github.com/mitchellh/mapstructure"
)

// SubscriptionInfo holds the results of the Subscription Information call. The Slots are a map of strings to Slot data
type SubscriptionInfo struct {
	Slots map[string]Slot `json:"slotsSubscriptionStatus"`
}

// Slot represents a DRS Slot and it's subscription status
type Slot struct {
	ProductInfoList []ProductInfoListItem `json:"productInfoList"`
	Subscribed      bool                  `json:"subscribed"`
}

// ProductInfoListItem represents a single ASIN in the subscription
type ProductInfoListItem struct {
	ASIN     string `json:"asin"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

// GetSubscriptionInfo gets the subscription information from DRS for the passed in device token
func GetSubscriptionInfo(deviceToken string) (*SubscriptionInfo, *APIError) {
	if deviceToken == "" {
		err := APIError{
			Code: 400,
			Data: map[string]string{
				"message": "deviceToken cannot be blank",
			},
		}
		return nil, &err
	}

	info := SubscriptionInfo{
		Slots: map[string]Slot{},
	}

	code, body, err := makeCall("subscriptionInfo", nil, deviceToken, map[string]string{})
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, err
	}

	for key, value := range body["slotsSubscriptionStatus"].(map[string]interface{}) {
		slotName := key
		slot := Slot{}
		slotErr := mapstructure.Decode(value, &slot)
		if slotErr != nil {
			return nil, &APIError{
				Code: 400,
				Data: map[string]string{
					"message": "Could not decode response",
				},
			}
		}

		info.Slots[slotName] = slot
	}
	return &info, nil
}
