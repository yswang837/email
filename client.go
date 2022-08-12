package email

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/yswang837/email/utils"
	"math/rand"
	"net/smtp"
	"time"
)

func SendEmail(to string) (string, error) {
	e := email.NewEmail()
	e.From = utils.EmailFrom
	e.To = []string{to}
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	str := "<h3> 验证码：" + code + "</h3> <p>仅" + utils.ActiveTime + "分钟内有效,请勿转发他人!</p>"
	e.HTML = []byte(str)
	if err := e.Send(utils.Addr+utils.Port, smtp.PlainAuth("", utils.Username, utils.Password, utils.Addr)); err != nil {
		return "", err
	}
	return code, nil
}
