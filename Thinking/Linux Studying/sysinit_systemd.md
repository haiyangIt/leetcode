# Sysvinit #
最开始的一种linux启动方式。

   也就是启动PID为1的进程，这个进程是<font color='red'>根进程</font>。其他进程都是由这个进程创建。

配合/etc/init.d/使用，放各种启动停止脚本。以及/etc/rc[X].d，放各种后台进程服务。

## 服务的最佳实践 ##
1. 进程启动时，关闭所有的文件描述符（除了0,1,2），然后重置所有的信号量。
2. 然后fork()子进程，在子进程中setsid()，使子进程成为进程组组长，同时，父进程退出。
3. 然后再次调用fork()创建孙子进程，然后子进程退出。
4. 在孙子进程中，将标准输入输出等链接到/dev/null上，创建其他...
5. 开始提供服务。


### 缺点： ###

+ **不支持热插拔。**

+ **无法对服务启动的所有子子孙孙进程进行监视**

# UpStart #

采用Event和Job的方式来管理

**优点：支持热插拔**

### 缺点：###

+ **事件管理混乱，启动顺序依赖混乱** 
+ **采用strace跟踪服务比较麻烦**

因此：

# Systemd #

## 原则 ##
+ 尽可能快的启动（并行）
+ 尽可能少的启动

## 方法 ##
+ 用C代替脚本
+ 解决三种依赖：
  
    + 文件系统： Lazy 启动，当真正需要访问时，才执行具体的文件操作。
    + Socket：缓存所依赖服务的socket的请求和数据。
    + D-Bus：缓存所依赖的服务的数据。
+ 采用cgroup来进行进程组的进程的管理