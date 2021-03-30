# Goal Go 工具类

## http client

### 使用方法

#### 发送Get请求
```go
    httpClient=http.http{}
    httpClient.Get("https://baidu.com")
    httpClient.Execute()
```

#### 发送Post请求

```go
    httpClient=http.http{}
    httpClient.Post("https://baidu.com","a=1&b=1")
    httpClient.Execute()
```

```go
    httpClient=http.http{}
    param:=url.Values{
    		"name":{"values"},
    	}
    httpClient.Post("https://baidu.com",param)
    httpClient.Execute()
```

#### 发送Post Json请求

```go
    httpClient=http.http{}
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

## convert 类型转化
```go
    
	var i = "1234567"
	var f = "123456.2345"
	var ff = 12345.6789
	var ii = 1234567


	log.Println(ii, convert.Int2String(ii))
	log.Println(ff, convert.Int2String(ff))

	log.Println(i, convert.Str2Int(i))
	log.Println(f, convert.Str2Int(f))

	log.Println(f, convert.Str2Float(f))
	log.Println(f, convert.Str2Float64(f))

	log.Println(i, convert.Str2Float(i))
	log.Println(i, convert.Str2Float64(i))

	log.Println(convert.Str2Url("1"))
	log.Println("encode base64 ",convert.B64Encode("12312312"))
	log.Println("decode base64 ",convert.B64Decode("324 d"))

	log.Println("urlencode ",convert.UrlEncode("324=1;sd;'123 d"))
	log.Println("urldecode  ",convert.UrlDecode("%25%27%22"))
	log.Println("rawurl ",convert.RawDecode("%25%27%22"))
	log.Println("raw encode  ",convert.RawEncode("324=1;sd;'123 d"))
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
