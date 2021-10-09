# sensitive-words-filter-service

[基于sensitive-words-math敏感词匹配的开箱即用的服务](https://github.com/dongweifly/sensitive-words-match)

支持单纯词匹配，组合词，正则的匹配方式。词库从文件中加载，请求接口采用Http API；

data目录的敏感词可以满足IM，评论，商品描述等场景最基础的内容合规要求，可以指定自己的敏感词；

## Usage

**COMPILE**

```bash
git clone https://github.com/dongweifly/sensitive-words-filter-service.git 
cd sensitive-words-filter-service
go mod tidy && go build
```

**START**

```bash
./sensitive-words-match -p 8088 -d ./data
```

启动参数说明： 

-p 指定http服务的端口；

-d 敏感词存储的路径，路径下所有的文件会自动遍历加载到内存，不支持二级目录；

<!-- ![image](https://user-images.githubusercontent.com/90187291/133043656-3a75fdc2-5193-438d-937e-b37f235662c1.png)
 -->
 
**HTTP API**

curl Request demo
```bash
curl --location --request POST 'http://127.0.0.1:8088/words/filter' --header 'Content-Type: application/json' --data-raw '{
    "text": "GOODO  fxxk,"
}'
```

包含敏感词的请求返回
```json
{
    "code": 200,
    "msg": "success",
    "data": {
        "suggestion": "block",
        "sensitiveWords": [
            "fuck"
        ],
        "desensitization": "GOODO  ****,。"
    }
}
```
## RoadMap

TODO
- [ ] 词库文件发生变更后能自动加载
- [ ] 支持从数据库加载词库
- [ ] 支持客户端自定义需要检测哪个词库
- [ ] 更加丰富的API接口

## Contact
[欢迎提交issue](https://github.com/dongweifly/sensitive-words-filter-service/issues) 

dongwei.fly@gmail.com

