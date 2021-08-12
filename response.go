package payload

type Response struct {
	WalletPayoutResponse
}

type WalletPayoutResponse struct {
	Balance int
}