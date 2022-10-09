package xuekewang

import (
	"time"
	"xuekewang/internal"
)

func NewSdkClient(appId, appSecret string, timeout time.Duration) *internal.SdkClient {
	return internal.NewSdkClient(appId, appSecret, timeout)
}
