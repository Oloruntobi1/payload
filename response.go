package payload

type Response struct {
	WalletPayoutResponse
}

type WalletPayoutResponse struct {
	Balance int `json:"balance" example:"1"`
}

type BanksResponse struct {
	RequestSuccessful bool `json:"requestSuccessful" example:"true"`
    ResponseMessage string `json:"responseMessage" example:"success"`
    ResponseCode string `json:"responseCode" example:"0"`
	ResponseBody []Banks `json:"responseBody"`
}

type Banks struct {
	Name string `json:"name" example:"Access bank"`
	Code string `json:"code" example:"044"`
	UssdTemplate string `json:"ussdTemplate" example:"*901*Amount*AccountNumber#"`
	BaseUssdCode string `json:"baseUssdCode" example:"901#"`
	TransferUssdTemplate string `json:"transferUssdTemplate" example:"*901*AccountNumber#"`
}


// git commit -a -m "my new version"
// git push
// git tag v0.1.4
// git push -q origin v0.1.4

