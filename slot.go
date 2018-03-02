package drs

import (
	"time"
)

/*SlotStatus represents the current status of the slot of a device and needs to be reported to Amazon at least on a daily basis.

ExpectedReplenishmentDate:	The expected product replenishment date in ISO 8601 format

RemainingQuantityInUnit:	Remaining quantity of the container (in the unit of measurement provided during DeviceCapabilitiesGroup creation)

OriginalQuantityInUnit:	Total quantity of product the container had when it was full (in the unit of measurement that you provided during DeviceCapabilitiesGroup creation)

TotalQuantityOnHand:	Total quantity of product on hand, but not loaded into the device (if known)

LastUseDate:	The last time that product was consumed from a given slot in ISO 8601 format
*/
type SlotStatus struct {
	ExpectedReplenishmentDate string  `json:"expectedReplenishmentDate"`
	RemainingQuantityInUnit   float64 `json:"remainingQuantityInUnit"`
	OriginalQuantityInUnit    float64 `json:"originalQuantityInUnit"`
	TotalQuantityOnHand       float64 `json:"totalQuantityOnHand"`
	LastUseDate               string  `json:"lastUseDate"`
}

/*ReportSlotStatus reports a slot status to Amazon. See the SlotStatus struct for information about the parameters
 */
func ReportSlotStatus(deviceToken string, slotID string, status *SlotStatus) (bool, *APIError) {
	if deviceToken == "" || slotID == "" {
		err := APIError{
			Code: 400,
			Data: map[string]string{
				"message": "deviceToken cannot be blank",
			},
		}
		return false, &err
	}

	expParsed, timeErr := time.Parse(time.RFC3339, status.ExpectedReplenishmentDate)
	if timeErr != nil {
		err := APIError{
			Code: 400,
			Data: map[string]string{
				"message": "ExpectedReplenishmentDate is not valid",
			},
		}
		return false, &err
	}
	status.ExpectedReplenishmentDate = expParsed.Format(time.RFC3339)

	lastParsed, timeErr := time.Parse(time.RFC3339, status.LastUseDate)
	if timeErr != nil {
		err := APIError{
			Code: 400,
			Data: map[string]string{
				"message": "ExpectedReplenishmentDate is not valid",
			},
		}
		return false, &err
	}
	status.LastUseDate = lastParsed.Format(time.RFC3339)

	code, _, err := makeCall("slotStatus", []interface{}{slotID}, deviceToken, status)
	if err != nil {
		return false, err
	}
	if code != 200 {
		return false, err
	}

	return true, nil
}
