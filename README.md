# Amazon Dash Replenishment Service SDK for Go

This library serves as a simple SDK for the Amazon Dash Replenishment Service. While trying to integrate with DRS, no official SDK existed. Given the small footprint of the API, a quick SDK was started until an official SDK is released.

*NOTE* This SDK exists as a temporary solution until an official SDK is released from Amazon or a community-standard SDK is developed. This SDK supports V2 of the DRS API.

## Installing

You can simply install the package:

`go get github.com/kevineaton/drs-sdk`

Or if you are using `dep`:

`dep ensure -add github.com/kevineaton/drs-sdk`

## Usage

First, there are some optional environment variables (with *hopefully* sane defaults):

`DRS_SDK_ENV` is the environment the SDK is running in; defaults to `development` which turns on logging

`DRS_SDK_ROOT_URL` is the root URL for the DRS service, defaults to `https://dash-replenishment-service-na.amazon.com/`

## Testing

Testing can be tricky, as you don't want to make actual network calls with user data. Therefore, in the `client.go` file where we setup the supported endpoints, we also have a `MockGood` field which holds JSON of what the call should return based upon the V2 docs found [here](https://developer.amazon.com/docs/dash/replenishment-service.html).

When testing, if the user auth token is set to `TEST`, the mock data will be used instead.

Since we are mocking the calls, some of the code is tougher to test, so code coverage will likely never be 100%.

To run the tests, run

`go test`

For coverage in HTML format, run

`go test -coverprofile=coverage.out && go tool cover -html=coverage.out`

## Contributing

Pull Requests are welcome! See our `CONTRIBUTING.md` file for more information.

## Roadmap

[X] Subscription Info [API Docs](https://developer.amazon.com/docs/dash/getsubscriptioninfo-endpoint.html)

[X] Deregistration [API Docs](https://developer.amazon.com/docs/dash/deregistration-endpoint.html)

[ ] Slot Status [API Docs](https://developer.amazon.com/docs/dash/slotstatus-endpoint.html)

[ ] Device Status [API Docs](https://developer.amazon.com/docs/dash/devicestatus-endpoint.html)

[ ] Get Order Info [API Docs](https://developer.amazon.com/docs/dash/getorderinfo-endpoint.html)

[ ] Replenish [API Docs](https://developer.amazon.com/docs/dash/replenish-endpoint.html)

[ ] Cancel Test Order [API Docs](https://developer.amazon.com/docs/dash/canceltestorder-endpoint.html)