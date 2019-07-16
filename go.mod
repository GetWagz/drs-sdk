module github.com/GetWagz/drs-sdk

go 1.12

require (
	github.com/go-resty/resty v1.9.1
	github.com/mitchellh/mapstructure v1.1.1
	github.com/sirupsen/logrus v1.1.0
	github.com/stretchr/testify v1.2.2
	golang.org/x/crypto v0.0.0-20180927165925-5295e8364332 // indirect
	golang.org/x/sys v0.0.0-20180928133829-e4b3c5e90611 // indirect
	gopkg.in/resty.v1 v1.12.0 // indirect
)

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0
