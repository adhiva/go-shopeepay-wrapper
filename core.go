package shopeepay

import (
	"errors"
	"fmt"
)

type CoreGateway struct {
	Client Client
}

func (gateway *CoreGateway) GenerateQRCode(req *RequestCharge) (ResponseCharge, error) {
	req.MerchantExtID = gateway.Client.MerchantExtID
	req.StoreExtID = gateway.Client.StoreExtID

	domain := fmt.Sprintf("%s%s", gateway.Client.APIEnvType.String(), gateway.Client.APIEnvRegion.String())
	response := ResponseCharge{}
	path := domain + "/v3/merchant-host/qr/create"

	err := gateway.Client.DoRequest("POST", path, req, &response)
	if err != nil {
		fmt.Println("Shopeepay Error Generate QR Code : ", err)
		errMessage := fmt.Sprintf("Shopeepay Generate QR Code Error : %s", err.Error())
		return response, errors.New(errMessage)
	}

	if response.ErrCode != 0 {
		fmt.Println("Shopeepay Error : ", response.DebugMsg)
		errMessage := fmt.Sprintf("Shopeepay Generate QR Code Error : %s", response.DebugMsg)
		return response, errors.New(errMessage)
	}

	return response, nil
}

func (gateway *CoreGateway) GenerateDeeplink(req *RequestCharge) (ResponseCharge, error) {
	req.MerchantExtID = gateway.Client.MerchantExtID
	req.StoreExtID = gateway.Client.StoreExtID

	domain := fmt.Sprintf("%s%s", gateway.Client.APIEnvType.String(), gateway.Client.APIEnvRegion.String())
	response := ResponseCharge{}
	path := domain + "/v3/merchant-host/order/create"

	err := gateway.Client.DoRequest("POST", path, req, &response)
	if err != nil {
		fmt.Println("Shopeepay Error Generate Deep Link : ", err)

		errMessage := fmt.Sprintf("Shopeepay Generate Deep Link Error : %s", err.Error())
		return response, errors.New(errMessage)
	}

	if response.ErrCode != 0 {
		fmt.Println("Shopeepay Error : ", response.DebugMsg)
		errMessage := fmt.Sprintf("Shopeepay Generate Deep Link Error : %s", response.DebugMsg)
		return response, errors.New(errMessage)
	}

	return response, nil
}

func (gateway *CoreGateway) CheckPaymentStatus(req *RequestGeneral) (ResponseCheckPaymentStatus, error) {
	req.MerchantExtID = gateway.Client.MerchantExtID
	req.StoreExtID = gateway.Client.StoreExtID

	domain := fmt.Sprintf("%s%s", gateway.Client.APIEnvType.String(), gateway.Client.APIEnvRegion.String())
	response := ResponseCheckPaymentStatus{}
	path := domain + "/v3/merchant-host/transaction/payment/check"

	err := gateway.Client.DoRequest("POST", path, req, &response)
	if err != nil {
		fmt.Println("Shopeepay Error Check Payment Status : ", err)

		errMessage := fmt.Sprintf("Shopeepay Check Payment Status Error : %s", err.Error())
		return response, errors.New(errMessage)
	}

	if response.ErrCode != 0 {
		fmt.Println("Shopeepay Error : ", response.DebugMsg)
		errMessage := fmt.Sprintf("Shopeepay Error Check Payment Status : %s", response.DebugMsg)
		return response, errors.New(errMessage)
	}

	return response, nil
}

func (gateway *CoreGateway) InvalidateQRCode(req *RequestGeneral) (ResponseCancelQRCode, error) {
	req.MerchantExtID = gateway.Client.MerchantExtID
	req.StoreExtID = gateway.Client.StoreExtID

	domain := fmt.Sprintf("%s%s", gateway.Client.APIEnvType.String(), gateway.Client.APIEnvRegion.String())
	response := ResponseCancelQRCode{}
	path := domain + "/v3/merchant-host/qr/invalidate"

	err := gateway.Client.DoRequest("POST", path, req, &response)
	if err != nil {
		fmt.Println("Shopeepay Error Invalidate QR Code : ", err)

		errMessage := fmt.Sprintf("Shopeepay Invalidate QR Code Error : %s", err.Error())
		return response, errors.New(errMessage)
	}

	if response.ErrCode != 0 {
		fmt.Println("Shopeepay Error : ", response.DebugMsg)
		errMessage := fmt.Sprintf("Shopeepay Invalidate QR Code Error : %s", response.DebugMsg)
		return response, errors.New(errMessage)
	}

	return response, nil
}
