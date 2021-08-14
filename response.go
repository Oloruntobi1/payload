package payload

type Response struct {
	WalletPayoutResponse
}

type WalletPayoutResponse struct {
	Balance int `json:"id" example:"1"`
}

// git commit -a -m "my new version"
// git push
// git tag v0.1.2
// git push -q origin v0.1.2

