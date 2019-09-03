# 协程简介

协程不是线程，很多时候协程被称为“轻量级线程”、“微线程”等。简单来说可以认为协程是线程里不同的函数，这些函数之间可以相互**快速**切换。

协程适用于 IO 密集型的任务，即 IO 时间长于 CPU 计算时间的任务，例如单点登录系统（SSO，不知道也没事，记得这是个高 IO 低 CPU 计算服务既可）。

以Worker 线程池，不使用消息队列的形式来实现 SSO 时，因为认证过程计算量很小（不考虑复杂加密和其他CPU密集型处理）所以整个请求耗时最大的部分一般在网络处理过程。一台机器上合理的线程数是有限的且线程的上下文切换也有一定的消耗，减少线程的切换次数与线程等待网络请求的时间可以提高系统的吞吐率。

这里使用几个简单的公式来描述系统的吞吐率。假设系统中线程上下文切换的时间间隔为 $\eta$（一个线程占用 CPU 时间片$\eta$ 后会休息一会，此时其他线程会占用 CPU 时间片）；假设系统切换线程的消耗为 $\xi$ （操作系统保存线程上下文，系统使用一定的策略调度线程，OS用户态与内核态的切换比较耗时）；假设每个 SSO 认证请求消耗 CPU 时间为 $\tau$，不考虑网络耗时时，线程在一个时间片 $\eta $ 内可以处理的请求个数为：$\frac {\eta}{\tau}$ ①。假设每个请求的网络耗时为 $\mu$，那么线程在一个时间片 $\eta $ 内可以处理的请求个数为：$\frac {\eta}{\tau + \mu}$。很明显，把网络请求时间从处理过程中去除，必然提高系统吞吐量。

从另一个角度来讲，CPU 不等待网络请求就意味着 CPU 一直不停的工作，CPU 又没有人那种长时间工作会降低效率的劣势 :cry:，效率肯定高。

为了降低网络请求对线程处理效率的影响，一个常见方案是使用消息队列。内网中网络耗时比公网要小很多，所以外网的访问可以缓存在消息队列中，worker 线程每次从消息队列中获取完整的请求数据。消息队列机器和worker机器可以使用不同的硬件配置以降低运维消耗：消息队列机器可以配置大内存，高速网络接口，CPU可以相对差一些；worker 机器可以提供比队列机器弱的网络能力，但一般要赋予更强的计算能力。worker 线程任劳任怨，只要机器之间搭配合理，请求队列中就会一直有数据供 worker 线程，不让其闲下来。所有机器都满负荷运算（暂不考虑冗余），可以最大化利用系统资源。

能不能连消息队列也去掉？如果请求数据直接在内存中，那获取请求数据的耗时几乎可以忽略，协程可以实现这个目的。

要是把线程直接绑定到指定的 CPU 核心，那么线程切换的耗时也可以省略了。

考虑下面这种处理方式：线程在处理一个服务时出现了网络请求，随后线程把网络请求交给操作系统（例如：epol）自己去处理其他服务，当网络请求结束时系统告诉线程“你有个服务可以继续处理了，你有空的时候处理下吧，数据我已经放到你指定的内存空间里了”。这个过程中操作系统充当了消息队列的角色。

从上面的描述中我们可以总结下协程的几个特点：

1. 协程可以自动让出CPU时间片。这里不是当前线程让出 CPU 时间片，而是线程内的某个协程让出时间片供其他协程运行
2. 协程可以恢复 CPU 上下文，当另一个协程可以继续的时候，其需要其 CPU 上下文环境
3. 协程有个管理者，其可以选择一个协程来运行
4. 运行中的协程将占有当前线程的所有计算资源

## ucontext

下面关于 ucontext 的介绍源自：[http://pubs.opengroup.org/onlinepubs/7908799/xsh/ucontext.h.html][2]

为了方便后面的叙述，这里先介绍一下 ucontext 这个 c 语言库。linux 系统一般都有这个库，这个库主要用于操控当前线程下的 CPU 上下文。

**说明：ucontext 只操作与当前线程相关的 CPU 上下文，所以下文中涉及 ucontext 的上下文均指当前线程的上下文。一般CPU 有多个核心，一个线程只能使用其中一个，所以 ucontext 只涉及一个与当前线程相关的 CPU 核心**

ucontext.h 头文件中定义了 `ucontext_t` 这个结构体，这个结构体中至少包含以下成员：

```c
ucontext_t *uc_link     // 当前上下文（协程）返回时会使用这个变量指向的上下文对象重置 CPU 上下文，如果为 NULL 则当前线程退出
sigset_t    uc_sigmask  // 当前上下文（协程）活跃时此变量中的信号都会被阻塞
stack_t     uc_stack    // 当前上下文（协程）所使用的栈
mcontext_t  uc_mcontext // 实际保存 CPU 上下文的变量，这个变量与平台&机器相关
```

同时，ucontext.h 头文件中定义了四个函数，下面分别介绍：

```c
int  getcontext(ucontext_t *); // 获得当前 CPU 上下文
int  setcontext(const ucontext_t *);// 重置当前 CPU 上下文
void makecontext(ucontext_t *, (void *)(), int, ...); // 修改上下文信息
int  swapcontext(ucontext_t *, const ucontext_t *);
```

### getcontext & setcontext

```c
#include <ucontext.h>
int getcontext(ucontext_t *ucp);
int setcontext(ucontext_t *ucp);
```

getcontext 函数使用当前 CPU 上下文初始化 ucp 所指向的结构体，初始化的内容包括 CPU 寄存器、信号mask和当前线程所使用的栈空间。

setcontext 使用 ucp 所指向的上下文信息重置当前 CPU 的上下文。setcontext 中的 ucp 一般由 getcontext 或 makecontext 创建，或者作为一个变量传递给信号处理函数，便于信号处理完成后的恢复。如果 ucp 指向的对象由 getcontext 创建，那么 setcontext 成功执行后线程的行为和执行 getcontext 后，继续执行原函数是一致的；如果 ucp 指向的对象由 makecontext 创建，那么 setcontext 成功执行后的行为和直接调用传递给makecontext 的函数是一致的。如果 ucp 所指向的 `ucontext_t` 对象中的 `uc_link` 为 0，则说明当前 context 为主 context，当前 context 返回时当前线程退出，**否则当前 context 退出后使用 uc_link 指向的对象重置 CPU 上下文并继续执行**。

**返回值**：getcontext 成功返回 0，失败返回 -1。**注意**，如果 setcontext 执行成功，那么调用 setcontext 的函数将不会返回，因为当前 CPU 的上下文已经交给其他函数或者过程了，当前函数完全放弃了 对 CPU 的“所有权”。

**应用**：当信号处理函数需要执行的时候，当前线程的上下文需要保存起来，随后进入信号处理阶段。可移植的程序最好不要读取与修改 `ucontext_t` 中的`uc_mcontext`，因为不同平台下`uc_mcontext`的实现是不同的。

### makecontext & swapcontext

```c
#include <ucontext.h>
void makecontext(ucontext_t *ucp, (void *func)(), int argc, ...);
int swapcontext(ucontext_t *oucp, const ucontext_t *ucp);
```

makecontext 修改由 getcontext 创建的上下文 ucp。如果 ucp 指向的上下文由 swapcontext 或 setcontext 恢复，那么当前线程将执行传递给 makecontext 的函数 `func(...)`。

执行 makecontext 后需要为新上下文分配一个栈空间，如果不创建，那么新函数`func`执行时会使用旧上下文的栈，而这个栈可能已经不存在了。argc 必须和 func 中整型参数的个数相等。

swapcontext 将当前上下文信息保存到 oucp 中并使用 ucp 重置 CPU 上下文。

**返回值**：swapcontext 成功则返回 0，失败返回 -1 并置 errno。

如果 ucp 所指向的上下文没有足够的栈空间以执行余下的过程，swapcontext 将返回 -1。

### 进一步学习

有很多协程库的实现是基于 ucontext 的，我们可以在学习这些库的时候顺便学习一下 ucontext 库的使用方法，下面介绍基于 c&ucontext 的协程库 coroutine 。

## [coroutine][3]，简单的 C 协程库

[coroutine][3] 是基于 ucontext 的一个 C 语言协程库实现，因为比较简单，所以先学习下这个库。coroutine 这个库适用场景有限，作者的目的是实现一个任务队列，这些任务执行完成之后即销毁。











[1]:https://jwt.io/
[2]:http://pubs.opengroup.org/onlinepubs/7908799/xsh/ucontext.h.html
[3]:https://github.com/cloudwu/coroutine/