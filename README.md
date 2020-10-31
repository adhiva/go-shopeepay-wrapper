go-shopeepay-wrapper
===========

A package for shopeepay wrapper client.

#### Installation
Make sure that Go is installed on your computer.
Type the following command in your terminal:

	go get github.com/adhiva/go-shopeepay-wrapper

After it the package is ready to use.


#### Import package in your project
Add following line in your `*.go` file:
```go
import "github.com/adhiva/go-shopeepay-wrapper"
```
If you are unhappy to use long `go-shopeepay-wrapper`, you can do something like this:
```go
import (
  shopeepay "github.com/adhiva/go-shopeepay-wrapper"
)
```

#### List of wrapper shopeepay client:
```go
func GenerateQRCode(req *RequestCharge) (ResponseCharge, error)
func GenerateDeeplink(req *RequestCharge) (ResponseCharge, error)
func CheckPaymentStatus(req *RequestGeneral) (ResponseCharge, error)
func InvalidateQRCode(req *RequestCharge) (ResponseCharge, error)
```

#### Examples
###### GenerateQRCode
```go
func init() {
	shopeePayClient = shopeepay.NewClient()
	shopeePayClient.ClientID = config.Config.ShopeePay.ClientID
	shopeePayClient.ClientSecret = config.Config.ShopeePay.ClientSecret
	shopeePayClient.MerchantExtID = config.Config.ShopeePay.MerchantExtID
	shopeePayClient.StoreExtID = config.Config.ShopeePay.StoreExtID
	shopeePayClient.APIEnvRegion = shopeepay.ID // Use this accordance your region, shopeepay is available in Indonesia, Singapore, Philippines, and Malaysia

	if config.Config.ShopeePay.ENV == "production" {
		shopeePayClient.APIEnvType = shopeepay.Production
	} else {
		shopeePayClient.APIEnvType = shopeepay.Sandbox
	}

	coreGateway = shopeepay.CoreGateway{
		Client: shopeePayClient,
	}
}

func ShopeePayCharge(request *shopeepay.RequestCharge) (*shopeepay.ResponseCharge, error) {
	var (
		resp shopeepay.ResponseCharge
		err  error
	)

	resp, err = coreGateway.GenerateQRCode(request)

	if err != nil {
		return nil, err
    }
    
	return &resp, nil
}
```

#### Support
If you do have a contribution to the package, feel free to create a Pull Request or an Issue.

#### What to contribute
If you don't know what to do, there are some features and functions that need to be done
- [ ] Create a Transaction List
- [ ] Create a Transaction Void
- [ ] Create a Refund Transaction

#### Advice
Feel free to create what you want, but keep in mind when you implement new features:
- Code must be clear and readable, names of variables/constants clearly describes what they are doing
- Wrapper they are created must be documented and described in source file and added to README.md to the list of wrapper shopeepay client

### Contributors 
- [Adhitya Giva Muhammad](https://github.com/adhiva)