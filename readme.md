## 使用
- git clone https://github.com/mkm1539806431/douyin.git
- 修改config/config.toml文件中的mysql, redis等配置信息
- cd douyin
- go mod tidy
- go run main.go
- 在浏览器地址栏输入 http://localhost:8080/hello 可以得到返回结果 {"Code":0,"Data":"成功"} 说明成功运行
## 各个包说明



|     包名     |     用途     |
|:----------:|:----------:|
|   common   | 常量、数据库等初始化 |
|   config   |    配置信息    |
| controller |    api     |
|    dao     |   数据库交互    |
| middleware |    中间件     |
|   model    |     模型     |
|  service   |    业务方法    |
|   utils    |    工具包     |

