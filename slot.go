package drs

import (
	"net/http"
	"time"
)

// SlotStatus represents the current status of the slot of a device and needs
// to be reported to Amazon at least on a daily basis.
type SlotStatus struct {
	// ExpectedReplenishmentDate is the expected product replenishment date in
	// ISO 8601 format
	ExpectedReplenishmentDate string `json:"expectedReplenishmentDate"`
	// RemainingQuantityInUnit is the emaining quantity of the container
	// (in the unit of measurement provided during DeviceCapabilitiesGroup creation)
	RemainingQuantityInUnit float64 `json:"remainingQuantityInUnit"`
	// OriginalQuantityInUnit is the total quantity of product the container
	// had when it was full (in the unit of measurement that you provided
	// during DeviceCapabilitiesGroup creation)
	OriginalQuantityInUnit float64 `json:"originalQuantityInUnit"`
	// TotalQuantityOnHand is the total quantity of product on hand, but not
	// loaded into the device (if known)
	TotalQuantityOnHand float64 `json:"totalQuantityOnHand"`
	//	LastUseDate is the last time that product was consumed from a given
	// slot in ISO 8601 format
	LastUseDate string `json:"lastUseDate"`
}

// ReportSlotStatus reports a slot status to Amazon. See the SlotStatus struct
// for information about the parameters
func ReportSlotStatus(deviceToken string, slotID string, status *SlotStatus) (bool, error) {
	if deviceToken == "" || slotID == "" {
		return false, &APIError{
			Code: http.StatusBadRequest,
			Data: map[string]string{
				"message": "deviceToken cannot be blank",
			},
		}
	}

	expParsed, timeErr := time.Parse(time.RFC3339, status.ExpectedReplenishmentDate)
	if timeErr != nil {
		return false, &APIError{
			Code: http.StatusBadRequest,
			Data: map[string]string{
				"message": "ExpectedReplenishmentDate is not valid",
			},
		}
	}
	status.ExpectedReplenishmentDate = expParsed.Format(time.RFC3339)

	lastParsed, timeErr := time.Parse(time.RFC3339, status.LastUseDate)
	if timeErr != nil {
		return false, &APIError{
			Code: http.StatusBadRequest,
			Data: map[string]string{
				"message": "ExpectedReplenishmentDate is not valid",
			},
		}
	}
	status.LastUseDate = lastParsed.Format(time.RFC3339)

	code, _, err := makeCall("slotStatus", []interface{}{slotID}, deviceToken, status)
	if err != nil || code != http.StatusOK {
		return false, err
	}

	return true, nil
}
