package test

import (
	"fmt"
	"oxpay/OxApi"
	"testing"
)

const (
	UserName = "123"
	Password = "345"
)

func TestLogin(t *testing.T) {

	err := OxApi.TerminalLogin(UserName, Password)
	fmt.Println(err)

}
