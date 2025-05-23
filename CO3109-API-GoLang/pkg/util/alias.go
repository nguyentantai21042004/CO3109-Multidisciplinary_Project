package util

import (
	"regexp"
	"strings"
)

func RemoveSignVietnamese(data string) string {
	data = regexp.MustCompile(`[àáạảãâầấậẩẫăằắặẳẵ]`).ReplaceAllString(data, "a")
	data = regexp.MustCompile(`[ÀÁẠẢÃĂẰẮẶẲẴÂẦẤẬẨẪ]`).ReplaceAllString(data, "A")
	data = regexp.MustCompile(`[èéẹẻẽêềếệểễ]`).ReplaceAllString(data, "e")
	data = regexp.MustCompile(`[ÈÉẸẺẼÊỀẾỆỂỄ]`).ReplaceAllString(data, "E")
	data = regexp.MustCompile(`[òóọỏõôồốộổỗơờớợởỡ]`).ReplaceAllString(data, "o")
	data = regexp.MustCompile(`[ÒÓỌỎÕÔỒỐỘỔỖƠỜỚỢỞỠ]`).ReplaceAllString(data, "O")
	data = regexp.MustCompile(`[ìíịỉĩ]`).ReplaceAllString(data, "i")
	data = regexp.MustCompile(`[ÌÍỊỈĨ]`).ReplaceAllString(data, "I")
	data = regexp.MustCompile(`[ùúụủũưừứựửữ]`).ReplaceAllString(data, "u")
	data = regexp.MustCompile(`[ƯỪỨỰỬỮÙÚỤỦŨ]`).ReplaceAllString(data, "U")
	data = regexp.MustCompile(`[ỳýỵỷỹ]`).ReplaceAllString(data, "y")
	data = regexp.MustCompile(`[ỲÝỴỶỸ]`).ReplaceAllString(data, "Y")
	data = regexp.MustCompile(`[Đ]`).ReplaceAllString(data, "D")
	data = regexp.MustCompile(`[đ]`).ReplaceAllString(data, "d")
	return data
}

// build alias , remove dấu
func BuildAlias(data string) string {
	data = RemoveSignVietnamese(data)
	if !AlphabetString(data) {
		data = RemoveSpecialCharacter(data)
	}
	data = strings.ReplaceAll(data, "-", " ")
	data = strings.ToLower(data)
	data = strings.TrimSpace(data)
	data = strings.ReplaceAll(data, " ", "-")
	data = regexp.MustCompile(`[^a-z0-9-]`).ReplaceAllString(data, "")
	data = strings.ReplaceAll(data, "--", "-")

	return data
}

func BuildCode(data string) string {
	data = RemoveSignVietnamese(data)
	if !AlphabetString(data) {
		data = RemoveSpecialCharacter(data)
	}
	data = regexp.MustCompile(` {2,}`).ReplaceAllString(data, " ")
	data = strings.ReplaceAll(data, "-", " ")
	data = strings.ToLower(data)
	data = strings.TrimSpace(data)
	data = strings.ReplaceAll(data, " ", "_")
	data = regexp.MustCompile(`[^a-z0-9_]`).ReplaceAllString(data, "")
	data = strings.ReplaceAll(data, "--", "_")
	data = strings.ToUpper(data)

	return data
}

func IsValidCode(name string, code string) bool {
	return code == BuildCode(name)
}
