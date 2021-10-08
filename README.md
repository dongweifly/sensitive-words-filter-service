# sensitive-words-filter-service

一个简单的基于DFA算法和Http的敏感词过滤服务；[基于此repo的开箱即用的service](https://github.com/dongweifly/sensitive-words-match)

TODO
- [ ] 词库文件发生变更后能自动加载
- [ ] 支持从数据库加载词库
- [ ] 支持客户端自定义需要检测哪个词库
- [ ] 更加丰富的API接口

## Usage

**START**

```bash
go mod tidy && go build
```

启动
```bash
./sensitive-words-match -p 8088 -d ./data/
```


启动参数说明： 
-p 指定http服务的端口；
-d 敏感词存储的路径，路径下所有的文件会自动遍历加载到内存，不支持二级目录；

![image](https://user-images.githubusercontent.com/90187291/133043656-3a75fdc2-5193-438d-937e-b37f235662c1.png)

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


## Contact
dongwei.fly@gmail.com


