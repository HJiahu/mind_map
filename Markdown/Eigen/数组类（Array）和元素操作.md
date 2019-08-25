# 数组类（Array）和元素操作

原文地址：[`http://eigen.tuxfamily.org/dox/group__TutorialArrayClass.html`][0]
## Array 类有什么用？
Array 类提供了一个一般用途的数组，用于操作元素相关算法，就像 Matrix 专门用于实现线性代数算法。更具体一点，Array 提供了对元素进行操作的方法，这些操作大部分与线性代数算法无关。比如数组中每个元素都加一个常量或者两个同维数组对应元素相乘。

## Array 类型
Array 是一个模板类，和Matrix有着相同的模板参数，前三个模板参数如下：  
`Array<typename Scalar, int RowsAtCompileTime, int ColsAtCompileTime>`

模板参数的功能与 Matrix 相同，这里就不在重复解释。

与 Matrix 相同，Eigen 也内置了一些常用的 Array 对象，不过与 Matrix 不同指出在于Array同时定义了一维和二维Array。内置 Array 类型的任意维度均不大于4：   

```c++
Array<float,Dynamic,1> ArrayXf 
Array<float,3,1> Array3f 
Array<double,Dynamic,Dynamic> ArrayXXd 
Array<double,3,3> Array33d 
```
## Array 元素的引用
与 Matrix 相同，Array 重载了小括号以实现 Array 对象中元素的读写。同样，Array重载了 `<<` 以实现 Array 的初始化或者打印（`std::cout<<array;`）。

更多有关初始化的话题请参考[这里][1]。

## 加减运算
Array 的加减运算是 element-wise 的，也就是相同尺寸 Array 对应元素之间的加减。

Array 也提供了 Array 和标量之间的加减运算（这是与 Matrix 的不同之处）：`array+scalar`，Array 每个元素都与 scalar 相加。

## 乘法运算
两个相同尺寸的 Array 对应元素之间的乘法。Matrix 中也有对应的方法：`.cwiseProduct(.)`

## 其他元素相关的运算
Array 提供了其他有用的元素相关的运算，例如，求元素绝对值的 abs()、求平方根的 sqrt()、求两个 Array 相同位置最小值的 min(.) 等，示例如下：

```c++
cout << a.abs() << endl;
cout << a.abs().sqrt() << endl;
cout << a.min(a.abs().sqrt()) << endl;
```
更多方法可参考：[Quick reference guide][2]。

## Matrix和Array之间的转换
Eigen 中 Matrix 和 Array 的职能不同，前者完成线性代数相关算法，后者完成元素相关的计算，你需要按照自己需求在二者之间进行选择。有些时候我们需要同时对数据进行线性代数相关的操作和元素相关的操作，这个时候就需要在 Matrix 和 Array 之间进行转换。

Matrix 有一个 array() 方法可将 Matrix 转化为 Array 对象。同样，Array 有一个 matrix() 方法可将 Array 对象转化为 Matrix 对象。和 Eigen 表达式优化一样，这种转换没有运行时消耗（假设你允许编译器优化）。.array() 方法和 .matrix() 方法的返回值均可以当做左值或者右值进行使用。

Eigen 不允许在同一个表达式中同时使用 Array 和 Matrix 对象，也就是说你不能直接将 Array 和 Matrix 对象相加。重载的运算符 `+`，其两侧要么全是 Array，要么全是 Matrix 对象。赋值运算符是个例外，你可以将 Matrix d对象赋予 Array 对象，同样你也可以将 Array 对象赋予 Matrix对象。

```c++
  MatrixXf m(2,2);
  MatrixXf n(2,2);
  MatrixXf result(2,2);
  m << 1,2,
       3,4;
  n << 5,6,
       7,8;
  result = m * n;
  result = m.array() * n.array();
  result = m.cwiseProduct(n);
  result = m.array() + 4;
```

其他例子，如：`(m.array() + 4).matrix()*m`，`(m.array() * n.array()).matrix() * m`。



[0]:http://eigen.tuxfamily.org/dox/group__TutorialArrayClass.html
[1]:http://eigen.tuxfamily.org/dox/group__TutorialAdvancedInitialization.html
[2]:http://eigen.tuxfamily.org/dox/group__QuickRefPage.html