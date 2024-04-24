package domain

type DeliveryInfo struct {
	DeclaredSumNano uint64
	CourierFeeNano  uint64

	Description string

	From Point
	To   Point
}
