**使用方法**
```go
import "github.com/LFreedomDev/xuekewang/sdk"

func main(){
    cli := xuekewang.NewSdkClient("appId","appSecret",time.Second*30)
    res, err := cli.GetSubjects()
}
```