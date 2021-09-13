# sensitive-word-match

一个简单的基于DFA算法和Http的敏感词过滤服务；

## Usage

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

http api /words/filter

client post 请求：
```json
{
  "text": "xdcvdfdsf fuck"
}
```

正常请求返回：
```json
{
    "code":200,
    "msg":"success",
    "data":{
        "suggestion":"pass",
        "desensitization":"Hello world, 世界你好"
    }
}
```
包含敏感词的请求返回
```json
{
    "code": 200,
    "msg": "success",
    "data": {
        "suggestion": "block",
        "sensitiveWords": [
            "fuck",
            "他妈的"
        ],
        "desensitization": "打扫房间****,***。。。。"
    }
}
```

curl Request demo
```bash
curl --location --request POST 'http://127.0.0.1:8088/words/filter' --header 'Content-Type: application/json' --data-raw '{
    "text": "打扫房间fuck,他妈的。。。。"
}'
```

执行显示
![image](https://user-images.githubusercontent.com/90187291/133043656-3a75fdc2-5193-438d-937e-b37f235662c1.png)
