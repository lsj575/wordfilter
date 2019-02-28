# 敏感词检测API

基于词典的敏感词检测系统

敏感词用Trie树存储

Go Web框架：iris [官网](https://studyiris.com/)

## install

```shell
go get github.com/lsj575/wordfilter

cd $GOPATH/src/github.com/lsj575/wordfilter

// 启动, 默认绑定9712端口
go run main.go 9712
```

## API

### 敏感词

#### 1.检测敏感词

输入一段字符串，将返回检测到的敏感词和用*号代替敏感词后的文本，若设置了白名单，则对白名单词不做替换

> GET: /check?content=

> POST: /check
>
> HEADER: sign 一段加密过的字符串，详情查看安全一节
>
> Param: content（string）带检测字符串

- 示例（以POST为例）

  - Request: `http://localhost:9712/check`

  - HEADER:

    - sign:c2Vuc2l0aXZlJTdDJTdDMTU1MTI4NDQ5MiU3QyU3Q3Rva2Vu

  - Params:

    - content:苍井空（あおい そら），日本AV演员兼电视、电影演员

  - Response:

    ```json
    {
        "err_code": 0,  // 0 表示没错误， 1 代表发生错误
        "msg": "OK",    // 信息
        "data": {
            "ret": 1,   // 1代表检测出敏感词
            "result": "苍井空，****演员",  //*号代替后的文本
            "find": [     // 查找到的敏感词
                "日本AV"
            ]
        }
    }
    ```

#### 2.添加敏感词

支持一次性添加多个敏感词，中间用英文逗号分隔

> POST:/sensitiveword
>
> PARAMS: words （string）要添加的敏感词
>
> Authorization: username:root password:root

- 示例

  - Request：`http://localhost:9712/sensitiveword`

  - Params:

    - words: 三陪

  - Response：

    ```json
    {
        "err_code": 0,
        "msg": "OK",
        "data": null
    }
    ```

    - 如果已经存在

      ```json
      {
          "err_code": 1,
          "msg": "陪睡 is exist",
          "data": null
      }
      ```

### 白名单

#### 1.添加白名单词

支持一次性添加多个白名单词，中间用英文逗号分隔

> POST:/whiteword
>
> PARAMS: words （string）要添加的敏感词
>
> Authorization: username:root password:root

- 示例

  - Request：`http://localhost:9712/sensitiveword`

  - Params:

    - words: 路口

  - Response：

    ```json
    {
        "err_code": 0,
        "msg": "OK",
        "data": null
    }
    ```

    - 如果已经存在

      ```json
      {
          "err_code": 1,
          "msg": "路口 is exist",
          "data": null
      }
      ```

### 安全

#### sign

- base64加密字符串
- 字符串：`sensitive||(当前时间戳)||token`

