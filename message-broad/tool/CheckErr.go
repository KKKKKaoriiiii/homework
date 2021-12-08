package tool

import "fmt"

// CheckErr 检查错误
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
