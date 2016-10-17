package user

func checkUserName(s string) bool {
	if s == nil || len(s) == 0 {
		return false
	}
	// TODO: pattern for user name
	return true
}

func checkEmail(s string) bool {
	if s == nil || len(s) == 0 {
		return false
	}
	// TODO: pattern for email
	return true
}
