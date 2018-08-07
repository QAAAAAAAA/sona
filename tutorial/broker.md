### broker原理

broker是sona的数据管理中心，负责：
1、维护与`agent`的长连接，方便各节点订阅、拉取最新配置，以及实时推送最新数据
2、与mongoDB关联，管理数据向mongoDB中的写入，并作为mongoDB的缓存层

具体而言：

1、`agent`订阅or拉取某服务配置，`broker`直接从其缓存中拿出配置回复；
如果不在缓存中，则从mongo中读取配置，写到缓存，并回复；

2、管理员试图新增、清空、修改某个服务的配置，`broker`将直接写入mongo，如果写入成功，顺便回写缓存

3、每个服务配置会在一定时间（100s）后~~在缓存中删除~~被特定的Goroutine重拉，以便更新为最新配置

>若在缓存中删除，那Mongo挂了会导致缓存渐渐变空：删除之后的重拉一直失败

### 各请求流程

#### agent订阅某服务配置
broker端若不存在，返回订阅失败；否则订阅成功，回复配置内容和版本号，记录订阅关系

#### agent重拉某服务配置
broker端若存在，返回配置内容和版本号，否则返回空内容

#### admin新增服务配置
broker端若存在，回复失败；否则写入mongo+回写缓存，回复成功

#### admin修改、清空服务配置
请求要携带版本号v，broker端若存在且版本号一致，则修改mongo+回写缓存，回复成功，顺便将此服务配置推送给订阅者`agent`们；否则失败

### QA

Q1、既然数据更新会回写到缓存，为什么还要周期性重拉最新配置？

A1：对于 主broker 来说，确实重拉的意义不大，这里主要是为了主备切换。
当 备broker 切换为 主broker后，其为`agent`提供的数据可能不是最新的，
而通过过期缓存策略，使得成为主的备broker在过一段时间后会达到与mongo的数据一致性

Q2、如果写入mongo产生超时错误，无法确定数据是否真的写入，怎么处理？

A2：视为操作失败，管理员一般会重试

Q3：如果同时两个人在修改同一个服务配置，谁的结果覆盖谁的？

A3：为消除歧义，`broker`控制任意时刻仅可有一个人在编辑同一个服务配置

且数据的修改是需要提供版本号v的，用于告知`broker`：我希望在v版本基础上修改数据
如果mongo中此数据版本号不一致，则不允许修改（有点CAS的意思）