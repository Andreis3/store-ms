package util

func StringToPointer(s string) *string {
	return &s
}

func StringPointerToString(sp *string) string {
	if sp == nil {
		return ""
	}
	return *sp
}
