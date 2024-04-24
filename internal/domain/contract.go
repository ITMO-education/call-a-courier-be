package domain

type Contract struct {
	TonAddress   string `json:"tonAddress"`
	OwnerAddress string `json:"ownerAddress"`
}

type ListContractRequest struct {
	ListRequest
}
