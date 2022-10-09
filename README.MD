**使用方法**
```go
import "git.lfreedom.com/LFreedomDev/xuekewang"

func main(){
    cli := xuekewang.NewSdkClient("appId","appSecret",time.Second*30)
    res, err := cli.GetSubjects()
}
```