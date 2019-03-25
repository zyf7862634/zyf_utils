```sequence
title: 创建channel流程
participant sdk as s
participant peer as p
participant orderer as o
  s -> p: 创建流程：创建channel
  p -> p: SYSCC创建流程接口
  p --> s: 返回流程ID
  s -> o: 提交交易
  o -> o: 打包区块
  o --> p: 分发区块
  p --> s: 发送BlockEvent
  Note over s,o: 投票流程
  s -> p: 投票流程：{流程ID}
  p -> p: 投票接口
  p --> s: 流程id
  s -> o: 提交交易
  o --> s: 返回结果
  Note over s,o: 发送CONFIG_UPDATE交易
  s -> p: 提交流程ID,请求CONFIG_UPDATE交易
  p --> s: 返回CONFIG_UPDATE交易
  s -> o: 提交CONFIG_UPDATE交易
  o -> o: 执行更新
  o --> s: 返回更新结果
```





