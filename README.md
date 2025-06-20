# go-zero-im 学习记录

## 数据库与缓存一致性的问题

- 一致性的级别
  - 弱一致性：网站评论、社交平台
    - 先更新缓存，再更新数据库
    - 先更新数据库，再更新缓存
    - 先删除缓存、再更新数据库
    - 先更新数据库，再删除缓存
    - 【更新缓存】
      - 优点：如果每次数据变化都难被及时更新，那么查询数据时不容易出现不命中的情况
      - 缺点：
        - 如果数据的计算复杂，频繁的更新会造成服务器性能的消耗比较大
        - 如果数据并不是被频繁使用，那么频繁的更新也只是浪费服务器性能，对业务没多大的帮助
      - 适用于数据使用较为频繁，且数据的计算不那么复杂的场景
    - 【删除缓存】
  - 最终一致性：多阶点数据复制、日志统计
  - 强一致性：金融交易、在线支付


>延迟双删

- 先删除缓存
- 更新数据库
- 线程等待N秒
- 再删除缓存

存在的问题：
- 线程需要在更新数据库后，还要休眠N秒，再次淘汰缓存，等所有的操作都执行完，这一个更新操作才真正完成，降低了更新操作的吞吐量
  - 解决办法：异步淘汰策略
- 如果第二次缓存淘汰失败，这不一致性依旧会存在
  - 解决办法：重试机制，需要设置重试次数

---

先更新数据后删除缓存

## 最终一致性

- 队列方式：更新数据库的同时，再异步通过队列更新缓存
- 日志监听：通过对数据库的日志监听更新缓存

## 好有关系表设计

### 冗余方式

各自存一次用户好有关系

| id | user_id | friend_id |
|----|---------|-----------|
| 1  | A       | B         |
| 1  | B       | A         |
-----------------------------

- 优点: 设计简单明了,查询简约快速
- 缺点: 冗余

### 计算唯一的 key

| id | user_1_id | user_2_id | user_key |
|----|-----------|-----------|----------|
| 1  | A         | B         | AldBld   |

- 优点: 好友关系只会存储一条数据,会减少空间
- 缺点: 查询复杂度会高一些,需要用 union all 处理

