package user

type request struct {
	Username      string `json:"username"`
	WalletAddress string `json:"walletAddress"`
}
