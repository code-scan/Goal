# Goal Go Red-Team 工具类


> 文档可能更新不及时，具体使用可以看`tests`中的文件


基于该项目我写了：


- [x] [求索 协同渗透测试平台](http://aq.mk/index.php/archives/94.html)
- [x] GoRat 跨平台C2
- [x] PaperCache [不安全信息聚合阅读](https://unsafe.sh)
- [x] [蹉跎 黑客工具箱](http://aq.mk/index.php/archives/111.html)
- [x] [NexusHack](https://github.com/code-scan/NexusHack)
- [x] [咪咕音乐下载器](https://github.com/code-scan/MiGuDownloander)
- [x] [WpGo wordpress批量快速爆破工具](https://github.com/code-scan/WpGo)
- [x] [DirGo 目录扫描小工具](https://github.com/code-scan/DirGo)
- ...


- #### Ghttp    Http客户端
    - [x] Get/Post/Head..
    - [x] SetHeader
    - [x] Set/GetCookie
    - [x] Save To File
    - [x] Use Http/Socks5 Proxy
  
    
- #### Gconvert 常用类型转化方法
    - [x] HostName/Domain To Ip
    - [x] Icmp Host Alive Check
    - [x] Tcp Port Status Check
    - [x] CIDR To Addres List
    
- #### Gsensor  信息获取探针
    - [x] Fofa
    - [x] SecurityTrails
    - [x] shodan
    - [x] Bufferover
    - [x] Crt.sh
    - [x] ZoomEye
    - [x] 爱企查
    - [x] ksubdomain
    - [x] WappalyzerGo

- #### Gproxy   反向socks5代理
    - [x] Reverse Socks5 Proxy
    
- #### Gfile    常用文件操作
    - [x] Unzip 
    - [x] CheckExist
    - [x] ReadFileToString
    - [x] WriteStringToFile
    - [x] GetFileList

- ### GLogin 常用服务
    - [x] mysql
    - [x] postgres
    - [x] mssql
    - [x] ftp
    - [x] pop3
    - [x] smtp


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

#### 开启Session 自动记录cookie

```go
    httpClient := Ghttp.New()
    httpClient.Session()
    httpClient.New("GET","https://www.baidu.com")
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

#### 设置代理
```go
	httpClient.SetProxy("http://127.0.0.1:6152")
	httpClient.SetProxy("socks5://ss:ss@127.0.0.1:6153")
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

## Gsensor 第三方API
### Fofa
 - 端口获取
 - 子域名获取
 - 同服获取
 
 具体代码可看tests/sensor_test.go

### SecurityTrails
 - A记录历史解析
 - 子域名获取
 - 同服获取
 
  具体代码可看tests/sensor_test.go

### Shodan
 - 端口获取
 
### Bufferover
 - 子域名获取
  具体代码可看tests/sensor_test.go

### Crt.sh
 - 子域名获取

  具体代码可看tests/sensor_test.go

## Gproxy 反向Socks5代理

```go
	// 服务端运行
	go Gproxy.ClientWait("8888") // 监听客户端端口，等待客户端链接
	go Gproxy.ServerWait("8889") // 服务端端口，用户链接用于代理
	// 客户端运行
	Gproxy.RunProxy("127.0.0.1:8888")
```
