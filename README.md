# Goal Go 工具类

## Ghttp client

### 使用方法

#### 发送Get请求
```go
    httpClient=Ghttp.http{}
    httpClient.Get("https://baidu.com")
    httpClient.Execute()
```

#### 发送Post请求

```go
    httpClient=Ghttp.http{}
    httpClient.Post("https://baidu.com","a=1&b=1")
    httpClient.Execute()
```

```go
    httpClient=Ghttp.http{}
    param:=url.Values{
    		"name":{"values"},
    	}
    httpClient.Post("https://baidu.com",param)
    httpClient.Execute()
```

#### 发送Post Json请求

```go
    httpClient=Ghttp.http{}
    jsonData:=make(map[string]interface{})
    jsonData["key"]="value"
    jsonData["json2"]=make(map[string]interface{})
    jsonData["json2"]["data"]="1234"
    httpClient.Post("https://baidu.com",jsonData)
    httpClient.Execute()
```

#### 获取字符串返回值

```go
    responseText:=httpclient.Text()
    log.Println(responseText)
```
#### 获取Byte返回值

```go
    responseByte:=httpclient.Byte()
    log.Println(responseByte)
```

#### 获取StatusCode

```go
    code:=httpclient.StatusCode()
    log.Println(code)
```

## Gconvert 类型转化
```go
    
	var i = "1234567"
	var f = "123456.2345"
	var ff = 12345.6789
	var ii = 1234567


	log.Println(ii, Gconvert.Int2String(ii))
	log.Println(ff, Gconvert.Int2String(ff))

	log.Println(i, Gconvert.Str2Int(i))
	log.Println(f, Gconvert.Str2Int(f))

	log.Println(f, Gconvert.Str2Float(f))
	log.Println(f, Gconvert.Str2Float64(f))

	log.Println(i, Gconvert.Str2Float(i))
	log.Println(i, Gconvert.Str2Float64(i))

	log.Println(convert.Str2Url("1"))
	log.Println("encode base64 ",Gconvert.B64Encode("12312312"))
	log.Println("decode base64 ",Gconvert.B64Decode("324 d"))

	log.Println("urlencode ",Gconvert.UrlEncode("324=1;sd;'123 d"))
	log.Println("urldecode  ",Gconvert.UrlDecode("%25%27%22"))
	log.Println("rawurl ",Gconvert.RawDecode("%25%27%22"))
	log.Println("raw encode  ",Gconvert.RawEncode("324=1;sd;'123 d"))
```
输出
```go

=== RUN   TestConvert
2021/03/30 17:38:00 1234567 1234567
2021/03/30 17:38:00 12345.6789 12345.6789
2021/03/30 17:38:00 1234567 1234567
2021/03/30 17:38:00 [!] Str2Int Error:  strconv.Atoi: parsing "123456.2345": invalid syntax
2021/03/30 17:38:00 123456.2345 0
2021/03/30 17:38:00 123456.2345 123456.234
2021/03/30 17:38:00 123456.2345 123456.2345
2021/03/30 17:38:00 1234567 1.234567e+06
2021/03/30 17:38:00 1234567 1.234567e+06
2021/03/30 17:38:00 1
2021/03/30 17:38:00 encode base64  MTIzMTIzMTI=
2021/03/30 17:38:00 [!] b64decode Error:  illegal base64 data at input byte 3
2021/03/30 17:38:00 decode base64  
2021/03/30 17:38:00 urlencode  324%3D1%3Bsd%3B%27123%20d
2021/03/30 17:38:00 urldecode   %'"
2021/03/30 17:38:00 rawurl  %'"
2021/03/30 17:38:00 raw encode   324%3D1%3Bsd%3B%27123%20d
--- PASS: TestConvert (0.00s)
```# Goal


## Gnet 网络相关

### TCP端口检测


```go
// 参数分别为 ip,端口，超时
// 返回bool 
Gnet.TcpPortStatus("127.0.0.1", 80, 30) 

```

Ping主机存活

```go
// 返回bool  需要root权限
Gnet.PingHost("127.0.0.1")
```


CIDR生成IP列表

```go
// 返回[]string
r := Gnet.GetIPList("192.168.1.1/24")
```