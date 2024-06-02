package util

func StringPointerToString(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}
