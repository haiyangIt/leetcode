# 介绍 #
Headless Service 是一类特殊的Service。
其Cluster IP 是None. 因此dns中会记录所有的这个service对应的Pod的ip地址。

将负载均衡委托给外部。

使用dns访问时，范围所有的pod的proxy地址。

