# 联想 Y7000 安装 Ubuntu18

为了方便使用自己这两天花了点时间在联想Y7000中安装了Ubuntu双系统，下面记录一下安装过程遇到的一些问题和解决办法。

## 系统安装

官网上下载最新的 Ubuntu18 系统，为了使用优盘安装 Ubuntu，需要先进入 bios 关闭 security boot 选项

安装过程中请勾选：安装第三方驱动选项

我的电脑上有两块 SSD，一块SSD用于安装win10，一块用户安装ubuntu。因为不想win10和ubuntu有任何的关联，所以在装ubuntu时我从主板上拔掉了装有win10系统的SSD

### 系统启动异常

现象：安装 ubuntu 后，重启可以点亮屏幕但无法进入桌面

原因：一般是显卡驱动的问题，因为 Y7000 相对较新，Ubuntu官方带的开源 NVIDIA 显卡驱动还不支持

解决办法（安装闭源驱动）：

* 重启，当桌面显示 Lenovo 图标的时候点击 Esc 键（如果失败，可多尝试几次）
* 选择 recovery 模式，选择 root ，进入命令行
* 编辑 `/etc/modprobe.d/blacklist.conf`，在最后一行添加，`blacklist nouveau`
* 执行：`sudo update-initramfs -u`
* 现在可以重启并进入桌面
* 进入终端，添加闭源驱动`sudo add-apt-repository ppa:graphics-drivers/ppa`
	 更新与安装NVIDIA官方驱动：执行 `sudo apt-get update` 和 `sudo apt-get install nvidia-driver-430，注意这里可以选择430，也可以选择其他可用的版本	`

###  wifi 无法联网

* 指令 `rfkill list all` 可以查看 Y7000 无线设备的状态，执行指令后发现无线网卡被关闭了，需要打开
* 在终端执行 `blacklist ideapad_laptop`，随后可以联网
* 创建 `/etc/modprobe.d/ideapad.conf`并添加一行：`blacklist ideapad_laptop`，系统启动时将自动执行这条指令并启用无线网卡

## 常用软件安装

### 搜狗输入法

使用下列指令安装sogou，然后重启系统，点击`ctrl+space`可以切换输入法

```bash
sudo apt-get remove ibus
sudo apt-get purge ibus # 清除ibus配置
sudo  apt-get remove indicator-keyboard
sudo apt install fcitx-table-wbpy fcitx-config-gtk
im-config -n fcitx
wget http://cdn2.ime.sogou.com/dl/index/1524572264/sogoupinyin_2.2.0.0108_amd64.deb?st=ryCwKkvb-0zXvtBlhw5q4Q&e=1529739124&fn=sogoupinyin_2.2.0.0108_amd64.deb
sudo dpkg -i sogoupinyin_2.2.0.0108_amd64.deb # 可能下载的文件名和当前名称不同
sudo apt-get install -f # 修复损坏或者缺失的包
```

### 百度网盘

从[官网][1]下载linux（deb格式）版的安装软件，使用指令`sudo dpkg -i baidu...deb` 形式进行安装。安装百度网盘是因为在我的网络环境下下载一些其他软件非常的慢，使用网盘的离线下载再从网盘软件中下载，会快很多

### vs code

```bash
wget -q https://packages.microsoft.com/keys/microsoft.asc -O- | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main"
sudo apt update
sudo apt install code
```

### chrome

```bash
wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | sudo apt-key add -
echo 'deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main' | sudo tee /etc/apt/sources.list.d/google-chrome.list
sudo apt-get update 
sudo apt-get install google-chrome-stable
```

### MB168B 驱动

可以从[官网][2]上下载linux 64bit版的驱动，不过需要先安装DKMS：`sudo apt-get install dkms`

### TexStudio
先下载texlive iso镜像，然后安装 tex 环境：`sudo ./install-tl`，输入 `I` 进行安装

```bash
sudo add-apt-repository ppa:sunderme/texstudio
sudo apt-get update
sudo apt-get install texstudio
```

### 其他常用软件

* typora
* seafile
* FoxitReader PDF



[1]:http://pan.baidu.com/download
[2]:https://www.asus.com/us/Monitors/MB169BPlus/HelpDesk_Download/