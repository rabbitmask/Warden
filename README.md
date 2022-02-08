# Warden
我可以写个网络测绘嘛～

Warden?
命名来自于魔兽史，典狱长 - 玛维·影歌

第一天～画个流程图～虽然我都不晓得ES集群是啥

```
=======================================================================
### Redis模块设计
IP    ====>   【Redis】缓存10天  ====>  如已存在，丢～  
                               ====>  如不存在，丢入爬虫模块
=======================================================================
### 爬虫模块设计
IP    ====>    端口查询与协议判断  ====>  指纹识别与框架标记    
                                ====> 丢入ES集群模块
=======================================================================
### ES集群模块设计
ES集群设计分布式存储
=======================================================================
### 资产梳理模块设计
ES集群拉取源数据   ====>  数据清洗与资产关联  ====>  重构ES数据源
=======================================================================
```

我凑，go是真的难封装，类似python的游标对象，我的redis为什么封装不了，只能傻乎乎的面向过程堆代码了
不过，go好像也没讲过自己是面向对象的～
```
Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy. The concept of “interface” in Go provides a different approach that we believe is easy to use and in some ways more general. There are also ways to embed types in other types to provide something analogous — but not identical — to subclassing. Moreover, methods in Go are more general than in C++ or Java: they can be defined for any sort of data, even built-in types such as plain, “unboxed” integers. They are not restricted to structs (classes).
Also, the lack of a type hierarchy makes “objects” in Go feel much more lightweight than in languages such as C++ or Java.
```

睡了12天，，所以，这算，，第二天？？？
ES完成初步封装，完成增查功能，目标导入 - redis - 爬虫 - 清洗 -入库ES - 数据查询 的布局大致初步成型