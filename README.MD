**使用方法**
```go
import "github.com/LFreedomDev/xuekewang/sdk"

func main(){
    cli := sdk.NewSdkClient("appId","appSecret",time.Second*30)
    res, err := cli.GetSubjects()
}
```