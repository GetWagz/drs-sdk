# Amazon Dash Replenishment Service SDK for Go

[![GoDoc](https://godoc.org/github.com/getwagz/drs-sdk?status.svg)](https://godoc.org/github.com/getwagz/drs-sdk)
[![Maintainability](https://api.codeclimate.com/v1/badges/fa368057b21ff923ca50/maintainability)](https://codeclimate.com/github/getwagz/drs-sdk/maintainability)
[![Build Status](https://travis-ci.org/getwagz/drs-sdk.svg?branch=master)](https://travis-ci.org/getwagz/drs-sdk)
[![Test Coverage](https://api.codeclimate.com/v1/badges/fa368057b21ff923ca50/test_coverage)](https://codeclimate.com/github/getwagz/drs-sdk/test_coverage)

This library serves as a simple SDK for the Amazon Dash Replenishment Service. While trying to integrate with DRS, no official SDK existed. Given the small footprint of the API, a quick SDK was started until an official SDK is released.

*NOTE* This SDK exists as a temporary solution until an official SDK is released from Amazon or a community-standard SDK is developed. This SDK supports V2 of the DRS API.

## Installing

You can simply install the package:

`go get github.com/getwagz/drs-sdk`

Or if you are using `dep`:

`dep ensure -add github.com/getwagz/drs-sdk`

## Usage

First, there are some optional environment variables (with *hopefully* sane defaults):

`DRS_SDK_ENV` is the environment the SDK is running in; defaults to `development` which turns on logging

`DRS_SDK_ROOT_URL` is the root URL for the DRS service, defaults to `https://dash-replenishment-service-na.amazon.com/`

Most calls require at a minimum the `deviceToken` which is the access token for the device obtained from Login With Amazon / DRS. The scope needed is `dash:replenish`. Refreshing this token is not the responsibility of this library.

## Testing

Testing can be tricky, as you don't want to make actual network calls with user data. Therefore, in the `endpoints.go` file where we setup the supported endpoints, we also have a `MockGood` field which holds JSON of what the call should return based upon the V2 docs found [here](https://developer.amazon.com/docs/dash/replenishment-service.html).

When testing, if the auth token is set to `TEST`, the mock data will be used instead.

Since we are mocking the calls, some of the code is tougher to test, so code coverage will likely never be 100%.

To run the tests, run

`go test`

For coverage in HTML format, run

`go test -coverprofile=coverage.out && go tool cover -html=coverage.out`

The coverage is notably lower than ideal, which may cause concerns. However, most of the uncovered calls would be calls directly to Amazon, which we cannot easily mock in success conditions, that are malformed. Feel free to check the results of the coverage report to see what exactly isn't covered and make a determination if that is acceptable to you. This library is currently being used in production.

## Contributing

Pull Requests are welcome! See our `CONTRIBUTING.md` file for more information.

## Third-party Libraries

The following libraries are used in this project. We thank the creators and maintainers for making our lives easier!

[Resty](https://github.com/go-resty/resty)

[Logrus](https://github.com/sirupsen/logrus)

[Testify](https://github.com/stretchr/testify)

[Mapstructure](https://github.com/mitchellh/mapstructure)

## Endpoints Implemented

Subscription Info [API Docs](https://developer.amazon.com/docs/dash/getsubscriptioninfo-endpoint.html)

Deregistration [API Docs](https://developer.amazon.com/docs/dash/deregistration-endpoint.html)

Device Status [API Docs](https://developer.amazon.com/docs/dash/devicestatus-endpoint.html)

Slot Status [API Docs](https://developer.amazon.com/docs/dash/slotstatus-endpoint.html)

Get Order Info [API Docs](https://developer.amazon.com/docs/dash/getorderinfo-endpoint.html)

Replenish [API Docs](https://developer.amazon.com/docs/dash/replenish-endpoint.html)

Cancel Test Order [API Docs](https://developer.amazon.com/docs/dash/canceltestorder-endpoint.html)