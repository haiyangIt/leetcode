## RC，RS，Deployment的区别 ##

### RC ###
Replication Controller: 用来部署，升级Pod，通过RC，可以控制Pod的数量，保证有效运行的Pod的个数，通过Selector来筛选要控制的Pod的。通过Template来定义Pod

### RS ###
Replication Set: 新一代的RC，拥有所有的RC的功能，Selector功能更加强大。

### Deployment ###
Deployment: 拥有更加强大的升级、回滚功能。

两种：Recreate和RollingUpdate

MaxSurge：滚动升级时，会先启动的pod个数
MaxUnavailable： 升级时，允许的最大Unavailable的个数。