# linux下常用工具

## 基础命令

* `adduser -d /home/hjiahu hjiahu`，添加用户，随后需要使用`passwd hjiahu`设置密码

## 远程控制

### SSH

`ssh -o StrictHostKeyChecking= no root@localhost`，不显示是否保存公钥记录的提示  

`ssh -f student@127.0.0.1 find / &> ~/find1. log `，后台自动执行   

CentOS7 下的重启：`systemctl restart sshd.service`

#### 建议

* 取消root直接登录的权限：`PermitRootLogin no`，***一定要记得提前创建非管理员账号***
* 禁止nossh群组与testssh用户使用sshd，这个一般需要主动配置
* 设置SSH的协议版本为2：`Protocol 2`
* 修改 SSH 端口：`Port 12345`
* /etc/hosts.allow 及/etc/hosts.deny
* iptables 封包过滤防火墙

### SFTP

模拟FTP    

常用命令：  

* 控制远程主机：cd、ls、dir
* 控制本地主机：lcd、lpwd，l 表示 local
* 上传指令：put local_file_dir remote_dir
* 下载指令：get remote_dir local_dir

### SCP

异地直接复制命令   

* 本地到远端：scp local_file user@ip:remote_dir
* 远端到本地：scp user@ip:remote_file local_dir

### rsync

Linux 系统备份工具