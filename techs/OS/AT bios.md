# AT BIOS 参考

## BIOS 中断

### INT10

显示中断

#### 显示一个字符

* 设置模式：`AH = 0x0e`
* 设置字符：`AL = ...`
* `BH = 0`
* 设置颜色：`BL = ...`，文本模式可以设置前后背景色与闪烁