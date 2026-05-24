package utils

func UpdateUserPaymentFailFirebase(userID string, isPaymentFail bool) error {
	updateData := map[string]interface{}{
		"is_payment_fail": isPaymentFail,
	}
	return UpdateUserFirebase(userID, updateData)
}
