# GoLang生产者、消费者验证

src/chan

## 基本功能
- 生产者， 不断产生一个msg， “hello world”
- 消费者， 不断读取msg， “hello world"
- 消息探针，不断判断消息的深度， 判断是否消费者阻塞；

定义来一个缓存buffer，目的是防止消费者取数据太慢，导致生产者阻塞，通过一个缓存通道来监控， 定义来一个通道大小为100；

