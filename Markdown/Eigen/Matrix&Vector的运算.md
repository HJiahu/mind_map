# Matrix&Vector的运算
原文地址：[`http://eigen.tuxfamily.org/dox/group__TutorialMatrixArithmetic.html`][0]

本章主要对矩阵、向量和标量之间的计算做一些简要介绍

## 介绍

为了实现矩阵（向量）之间的计算，Eigen 同时提供了运算符重载（+、-、×、/ 等）和类方法（dot()、corss() 等）两大形式的工具。对于 Matrix 类，重载的运算符只支持线性代数相关算法。例如，`matrix1*matrix2`意味着矩阵之间的点乘，`vector+scalar`是不被允许的表达式。如果你需要数组操作而非线性代数计算，可参考[这里][1]。

## 加减运算
加减运算符左右的矩阵必须有相同的尺寸，也必须有相同的元素类型，Eigen不支持自动类型提升。以下为运算符示例：`a+b、a-b、-a、a+=b、a-=b`。

## 矩阵与标量的乘除
矩阵乘或者除以标量的方法也很简单，例如：`matrix/scalar、m*s、s*m、m/=s、m*=s`

## 表达式注意事项
这里描述的是 Eigen 的高阶特性，具体内容可以参考[这里][2]，本节我们稍微提一下。在 Eigen 中算术运算符并不做实际的计算，这些运算符只是返回一个**被标识**要做相关计算的对象，真实的计算发生在对整个表达式求值的时候，一般是遇见赋值等号的时候。这样做有利于编译器做优化。举个例子：

```c++
VectorXf a(50), b(50), c(50), d(50);
...
a = 3*b + 4*c + 5*d;
```
Eigen 将上面的表达式编译为一个 for 循环，在不考虑其他优化（如 SIMD）的情况下可以等价为如下代码：

```
for(int i = 0; i < 50; ++i)
  a[i] = 3*b[i] + 4*c[i] + 5*d[i];
```
_译者注：如果所有的实际计算发生在运算符出现的地方，那么上面的代码最少需要三个 for循环分别迭代b、c、d这三个矩阵，和一个for循环比起来在循环次数上后者是前者的1/3_

所以不要怕使用非常长的算术表达式，长的算术表达式便于编译器做优化。

## 转置和共轭
_译者注：因为很久没有接触线性代数了，很多概念已经很生疏，故不对本小结再做翻译，以防出错。本小结主要介绍了一些求解矩阵特殊解的方法，例如求转置的transpose()等_

## 矩阵与矩阵和矩阵与向量之间的乘法
矩阵与矩阵（向量）之间的乘法使用运算符`*`，例如：`a*b、a*=b`

默认情况下 Eigen 在做矩阵乘法的时候会生成一个临时变量来保存计算值，然后赋予等号左边的变量。如果你确定你不需要临时变量，可以使用[`noalias()`][3] 方法

## 点乘与叉乘
Eigen使用 dot() 和 cross() 实现点乘和叉乘。
...

## 基本的算术约简运算
Eigen 提供了一些约简方法，将矩阵和向量转化为一个标量值。例如元素求和方法 sum()、元素乘积方法 prod()、最大元素方法 maxCoeff()、最小元素方法 minCoeff()。求对角元素之和可以使用方法 trace() 也可以使用 a.diagonal().sum() 方法。

## 计算合法性检查
Eigen 将检查你的运算是否合法。Eigen 会在编译或者运行时检查运算的有效性。那些在编译期间无法检查到的问题会在运行时进行检查。Debug 模式下 Eigen 会使用断言来检查计算的有效性与合法性，Release模式下如果出现非法运算，程序将直接崩溃并退出。






[0]:http://eigen.tuxfamily.org/dox/group__TutorialMatrixArithmetic.html
[1]:http://eigen.tuxfamily.org/dox/group__TutorialArrayClass.html
[2]:http://eigen.tuxfamily.org/dox/TopicEigenExpressionTemplates.html
[3]:http://eigen.tuxfamily.org/dox/group__TopicAliasing.html

