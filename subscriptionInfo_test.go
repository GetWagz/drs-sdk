package drs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGettingTheUserSubscription(t *testing.T) {
	ConfigSetup()
	_, err := GetSubscriptionInfo("")
	assert.NotNil(t, err)
	info, err := GetSubscriptionInfo("TEST")
	assert.Nil(t, err)
	//refer to the mock data in the client
	assert.True(t, info.Slots["slot1"].Subscribed)
	assert.Equal(t, "string", info.Slots["slot1"].ProductInfoList[0].ASIN)
	assert.Equal(t, 1, info.Slots["slot1"].ProductInfoList[0].Quantity)
	assert.Equal(t, "count", info.Slots["slot1"].ProductInfoList[0].Unit)

}
