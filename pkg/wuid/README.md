# wuid

这个工具会记录创建的id位置存储方式，可以是redis，也可以用mysql，
下面是mysql的表结构

```sql
CREATE TABLE `wuid` (
                        `h` int(10) NOT NULL AUTO_INCREMENT,
                        `x` tinyint(4) NOT NULL DEFAULT '0',
                        PRIMARY KEY (`x`),
                        UNIQUE KEY `h` (`h`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;
```