
============ Wed Mar 20 13:56:33 CST 2019
时序图 Sequence
```mermaid
sequenceDiagram;
    participant 服务端;
    participant 客户端;
    participant web页面;
    服务端->>客户端:1.请求
    客户端->>web页面:2.检查
    客户端->>服务端:3.握手
    客户端->>web页面:4.调整
    web页面->>web页面:5.检查
    客户端->>服务端:6.完成
```
