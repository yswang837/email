package email

import (
	"fmt"
	"github.com/yswang837/email/utils"
	"testing"
)

func TestSendEmail(t *testing.T) {

	code, err := SendEmail(utils.Username)
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", code)
}
