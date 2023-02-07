package cat

import "os"

func Logotype() (string, error) {
	text, err := os.ReadFile("../files/logo.txt")
	if err != nil {
		return "", nil
	}
	return string(text), nil
}
