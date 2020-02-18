基于 Golang 实现的 DNS 查询工具。主要目标是在无标准 nslookup 的设备上执行 nslookup 命令

### 使用方法

```shell

bin/nslookgo baidu.com
bin/nslookgo baidu.com 8.8.8.8

```

### 跨平台编译方法

```shell

sh package.sh

```

- 安卓设备可以使用 Linux_arm 的版本
- OpenWrt 可以使用 Linux_mipsle 的版本