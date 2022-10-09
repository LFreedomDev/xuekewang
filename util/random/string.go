package random

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func RandomString(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		temp = string(rune(RandInt(65, 90)))
		result.WriteString(temp)
		i++

	}
	return result.String()
}

func RandomCaptchaIntString() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}
