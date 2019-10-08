# linux 管理

说明：

当前文档以 ubuntu 18 作为操作系统进行讲解

## 安全问题

### 配置SSH

禁止SSH的root登录权限，可以使用普通用户登录SSH然后再切换为root。SSH的配置文件为：`/etc/ssh/sshd_config`，注意，不是 `ssh_config`

```
# 禁止 root 直接登录，默认为yes
PermitRootLogin no
# 修改Port，默认为 22
Port 10008
```

重启ssh：`service ssh restart`











































