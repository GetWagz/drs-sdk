package drs

import "net/http"

type endpoint struct {
	Path     string
	Method   string
	Headers  []endpointHeader
	MockGood string
}

type endpointHeader struct {
	Header string
	Value  string
}

// endpoints holds all of the endpoints the SDK currently supports, including
// the mockdata needed for tests
var endpoints = map[string]endpoint{
	"cancelTestOrder": endpoint{
		Path:   "testOrders/slots/%s",
		Method: http.MethodDelete,
		Headers: []endpointHeader{
			endpointHeader{
				Header: "x-amzn-accept-type",
				Value:  "x-amzn-accept-type: com.amazon.dash.replenishment.DrsCancelTestOrdersResult@1.0",
			},
			endpointHeader{
				Header: "x-amzn-type-version",
				Value:  "x-amzn-type-version: com.amazon.dash.replenishment.DrsCancelTestOrdersInput@1.0",
			},
		},
		MockGood: `{
			"slotOrderStatuses": [
				{
					"orderStatus": "NO_ORDER_IN_PROGRESS",
					"slotId": "slot1"
				},
				{
					"orderStatus": "NO_ORDER_IN_PROGRESS",
					"slotId": "slot2"
				}
			]
		}`,
	},
	"deregister": endpoint{
		Path:   "registration",
		Method: http.MethodDelete,
		Headers: []endpointHeader{
			endpointHeader{
				Header: "x-amzn-accept-type",
				Value:  "com.amazon.dash.replenishment.DrsDeregisterResult@1.0",
			},
			endpointHeader{
				Header: "x-amzn-type-version",
				Value:  "com.amazon.dash.replenishment.DrsDeregisterInput@2.0",
			},
		},
		MockGood: "",
	},
	"deviceStatus": endpoint{
		Path:   "deviceStatus/%s",
		Method: http.MethodPost,
		Headers: []endpointHeader{
			endpointHeader{
				Header: "x-amzn-accept-type",
				Value:  "com.amazon.dash.replenishment.DrsDeviceStatusResult@1.0",
			},
			endpointHeader{
				Header: "x-amzn-type-version",
				Value:  "com.amazon.dash.replenishment.DrsDeviceStatusInput@1.0",
			},
		},
		MockGood: "",
	},
	"getOrderInfo": endpoint{
		Path:   "getOrderInfo/%s",
		Method: http.MethodGet,
		Headers: []endpointHeader{
			endpointHeader{
				Header: "x-amzn-accept-type",
				Value:  "x-amzn-accept-type: com.amazon.dash.replenishment.DrsOrderInfoResult@1.0",
			},
			endpointHeader{
				Header: "x-amzn-type-version",
				Value:  "x-amzn-type-version: com.amazon.dash.replenishment.DrsOrderInfoInput@1.0",
			},
		},
		MockGood: `{
			"orderInfoData": {
				"instanceId": "amzn1.dash.v2.o.--------",
				"orderItems": [
					{
						"asin": "-------",
						"expectedDeliveryDate": "2017-01-05T07:59:59.000Z",
						"quantity": 1,
						"slotId": "PaperTowel",
						"status": "Pending"
					 }
				 ]
			 }
		 }`,
	},
	"replenishSlot": endpoint{
		Path:   "replenish/%s",
		Method: http.MethodPost,
		Headers: []endpointHeader{
			endpointHeader{
				Header: "x-amzn-accept-type",
				Value:  "com.amazon.dash.replenishment.DrsReplenishResult@1.0",
			},
			endpointHeader{
				Header: "x-amzn-type-version",
				Value:  "com.amazon.dash.replenishment.DrsReplenishInput@1.0",
			},
		},
		MockGood: `{
			"eventInstanceId" : "SOME_EVENT_INSTANCE",
			"detailCode" : "STANDARD_ORDER_PLACED"
			}`,
	},
	"slotStatus": endpoint{
		Path:   "slotStatus/%s",
		Method: http.MethodPost,
		Headers: []endpointHeader{
			endpointHeader{
				Header: "x-amzn-accept-type",
				Value:  "x-amzn-accept-type: com.amazon.dash.replenishment.DrsSlotStatusResult@1.0",
			},
			endpointHeader{
				Header: "x-amzn-type-version",
				Value:  "x-amzn-type-version: com.amazon.dash.replenishment.DrsSlotStatusInput@1.0",
			},
		},
		MockGood: "",
	},
	"subscriptionInfo": endpoint{
		Path:   "/subscriptionInfo",
		Method: http.MethodGet,
		Headers: []endpointHeader{
			endpointHeader{
				Header: "x-amzn-accept-type",
				Value:  "com.amazon.dash.replenishment.DrsSubscriptionInfoResult@2.0",
			},
			endpointHeader{
				Header: "x-amzn-type-version",
				Value:  "com.amazon.dash.replenishment.DrsSubscriptionInfoInput@1.0",
			},
		},
		MockGood: `{
			"slotsSubscriptionStatus": {
				"slot1": {
					"productInfoList": [{
						"asin": "string",
						"quantity": 1,
						"unit": "count"
					}],
					"subscribed": true
				}
			}
		 }`,
	},
}
