# 指定额外的搜索目录，使用冒号分隔，当前目录优先级最高
VPATH = ../../src:../headers         
# vpath <pattern> <directories> 是更灵活的搜索目录指定方式
vpath %.h ../headers

# 嵌套执行 makefile
subsystem: 
 cd subdir && $(MAKE)
 # $(MAKE) -C subdir 等价于上一句

# make 中的变量就是C语言中的宏定义，所以要注意 空格 
objects_1 = main.o print.o # 等号可以使用后面定义的变量
objects_1 += hello.o # 如果 objects_1 未定义，则 += 退化为 = 
objects_2 := $(wildcard *.o) # := 只能使用前面已经定义好的变量
objects_3 ?= bar # 如果 FOO 在前面定义过则这条语句什么都不做，否则 FOO 的值是 bar

# 变量替换
foo := a.o b.o c.o d.c e.f g.a 
bar := $(foo:.o=.c)   # a.c b.c c.c d.c e.f g.a
bar := $(foo:%.o=%.c) # a.c b.c c.c d.c e.f g.a

helloworld : $(objects_1) 
 gcc -o helloworld $(objects)
 # 在指令前加 @ 符号则当前指令不会显示在控制台，但执行结果会显示
 @echo compile helloworld 
 cd ../; pwd # 使用上一条指令的执行结果
 cd ../  
 pwd # 与上条指令的执行结果无关

# 静态模式，目标集为 objects_1 中以 .o 为后缀的字符串，依赖集是目标集后缀修改为 .c 的字符串
$(objects_1): %.o: %.c
 $(CC) -c $(CFLAGS) $< -o $@
$(objects_2) : print.h # 都依赖 print.h ，最好不要使用这种特性
mian.o  : mian.c # 自动推导 gcc -c main.c 
# make 自动关联.c文件并自动推导生成 foo.o 和 bar.o 的语句
# 注意这里没有推导头文件的依赖 且不适用于 C++  
foo : foo.o bar.o 
 cc –o foo foo.o bar.o $(CFLAGS) $(LDFLAGS)  

# 函数
PWD := $(shell pwd) # shell 函数
TIME:= $(shell date "+%Y-%m-%d %H:%M:%S") # shell 函数
OBJ += $(patsubst %.cpp, %.o, $(wildcard ./logic/*.cpp))

# 定义模式规则
%.o: %.cpp # 把所有的[.cpp]文件都编译成[.o]文件
 $(CC) $(CFLAG) -g -c -O3 $< -o $@ $(INCLUDE) 
 @echo $@

# 伪目标，要运行伪目标只能以 make clean 的方式
.PHONY : clean
clean :
 # 减号表示出错也不要退出，继续后面的指令
 -rm helloworld $(OBJ) 
# 利用伪目标一次生成多个目标，make all
all : prog1 prog2 
.PHONY : all
prog1 : prog2.o 
 commands
prog2 : prog1.o 
 commands

# 第五章第八节，自动生成依赖性