package drs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCancelTestOrder(t *testing.T) {
	ConfigSetup()
	_, err := CancelTestOrder("", "")
	assert.NotNil(t, err)
	_, err = CancelTestOrder("TEST", "")
	assert.NotNil(t, err)
	info, err := CancelTestOrder("TEST", "TEST")
	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, 2, len(info.SlotOrderStatuses))
}

func TestGetOrderInfo(t *testing.T) {
	ConfigSetup()
	_, err := GetOrderInfo("", "")
	assert.NotNil(t, err)
	_, err = GetOrderInfo("TEST", "")
	assert.NotNil(t, err)
	info, err := GetOrderInfo("TEST", "TEST")
	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, 1, len(info.OrderItems))
	assert.Equal(t, info.InstanceID, "amzn1.dash.v2.o.--------")
	assert.Equal(t, info.OrderItems[0].ASIN, "-------")
	assert.Equal(t, info.OrderItems[0].ExpectedDeliveryDate, "2017-01-05T07:59:59.000Z")
	assert.Equal(t, info.OrderItems[0].Quantity, 1)
	assert.Equal(t, info.OrderItems[0].SlotID, "PaperTowel")
	assert.Equal(t, info.OrderItems[0].Status, "Pending")
}
