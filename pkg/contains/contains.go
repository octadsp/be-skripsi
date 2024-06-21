package contains

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func OrderStatuses() []string {
	return []string{
		"WAITING FOR ORDER CONFIRMATION",
		"WAITING FOR PAYMENT",
		"WAITING FOR PAYMENT CONFIRMATION",
		"ACCEPTED",
		"REJECTED",
		"ON DELIVERY",
		"WAITING FOR CUSTOMER CONFIRMATION",
		"DONE",
	}
}

func OrderPaymentStatuses() []string {
	return []string{
		"WAITING FOR PAYMENT CONFIRMATION",
		"ACCEPTED",
		"REJECTED",
	}
}
