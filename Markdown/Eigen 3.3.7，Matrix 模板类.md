# Eigen 3.3.7，Matrix 模板类

原文地址：[`http://eigen.tuxfamily.org/dox/group__TutorialMatrixClass.html`][0]   

在 Eigen 中，所有的矩阵和向量都是 Matrix 模板类。向量是特殊的矩阵，有着一行或者一列数据。  

## Matrix 的前三个模板参数

Matrix 有六个模板参数，这里我们只介绍前三个。后三个模板参数有默认值，我们将在其他小结中做介绍。  

前三个模板参数如下所示：  
`Matrix<typename Scalar, int RowsAtCompileTime, int ColsAtCompileTime>` 

* Scalar 是矩阵中元素的类型。如果你想构建一个由浮点数组成的矩阵，将当前参数设置为float即可
* RowsAtCompileTime和ColsAtCompileTime用来在编译时确定矩阵的尺寸

Eigen 提供了一些常用的矩阵类型，比如 Matrix4f，在Eigen中其定义如下：  
`typedef Matrix<float, 4, 4> Matrix4f;`

## 向量
Eigen 中向量是特殊的矩阵，有一列或者一行数据。有一列数据的向量被称之为列向量（Eigen中的向量默认为列向量），有一行数据的向量被称为行向量。

举个例子，内置的**列向量**Vector3f类型在Eigen中的定义如下：   
`typedef Matrix<float, 3, 1> Vector3f;`      
Eigen同时提供了内置的行向量：  
`typedef Matrix<int, 1, 2> RowVector2i;`

## 动态矩阵
Eigen提供了编译时尺寸不明确的矩阵类。RowsAtCompileTime 和 ColsAtCompileTime可以被设置为`Dynamic`以标识矩阵的尺寸在编译时是不明确的，只能在运行时确定。在Eigen中编译时可确定尺寸的矩阵被称为固定尺寸矩阵（后文称之为定维矩阵），在运行时才能确定大小的矩阵被称为动态尺寸矩阵（后文称之为动态矩阵）。Eigen内置的MatrixXd，一个元素类型为double的动态矩阵的定义如下：  
`typedef Matrix<double, Dynamic, Dynamic> MatrixXd;`  
类似的，Eigen定义了VectorXi：  
`typedef Matrix<int, Dynamic, 1> VectorXi;`  
你可以定义一个维度未知而另一个维度已知的矩阵，如下：  
`Matrix<float, 3, Dynamic>`

## 构造函数
默认构造函数总是可用的，默认构造函数不会分配内存亦不会初始化矩阵元素：  

```c++
Matrix3f a; // 3*3 matrix, uninitialized
MatrixXf b; // 0-by-0 currently
```

Eigen提供接受指定矩阵尺寸参数的构造函数。对于矩阵，行数总是第一个参数。对于向量，只需要提供向量参数的个数即可。这些构造函数分配了矩阵或者向量所需要的内存空间但**没有初始化这些内存**。

```c++
MatrixXf a(10,15);
VectorXf b(30);
```

为了给定维矩阵和动态矩阵提供统一的API，对定维矩阵使用上面的构造函数也是合法的，例如：`Matrix3f a(3,3)`，虽然这样没有什么意义，但对于Eigen而言这不算错误。

Eigen 提供了向量元素初始化的构造函数，不过这类内置构造函数只支持向量尺寸小于等于4的情形：

```c++
Vector2d a(5.0, 6.0);
Vector3d b(5.0, 6.0, 7.0);
Vector4d c(5.0, 6.0, 7.0, 8.0);
```

## 矩阵元素引用
Eigen中主要的矩阵元素引用方式是使用重载后的小括号。对于矩阵，行索引总是第一个参数。对于向量而言，只需要一个索引值就可以了。所有索引值都从0开始。

只有一个索引值的元素引用语法 `m(index)`并不仅限于向量，也适用于一般的矩阵，意味着读取一行（列）元素。然而这个特性十分依赖Eigen的存储顺序。所有的Eigen矩阵默认按列存储在内存中，不过你可以手动修改存储顺序，可参考[Storage orders][1]。

运算符 `[]` 也被重载用于引用向量的元素，但是因为C++语言的特性，`[]`不能用于矩阵元素的引用，因为在C++中`[i,j]`和`[j]`的计算值是相同的，都是`[j]`。

## 逗号初始化
矩阵和向量都可以使用逗号初始化语法：

```c++
Matrix3f m;
m << 1, 2, 3,
     4, 5, 6,
     7, 8, 9;
std::cout << m;
```
[其他初始化方式][2]。

## 维度调整
Matrix 的尺寸信息可以通过 rows()、cols() 和 size() 获取，分别返回矩阵的行数、列数和所有元素的个数。矩阵尺寸调整可以通过方法 `resize()`。

如果矩阵的实际尺寸没有发生变化，那么resize() 方法将不进行任何实际操作。如果使用resize()调整了矩阵的尺寸，那么矩阵元素的值可能发生变化。如果你需要一个保守的不改变矩阵元素值的resize方法，可以参考[conservativeResize()][3]。

为了实现API的统一，所有调整矩阵尺寸的方法都适用于固定尺寸的矩阵，不过大部分时候这些函数都会报错或者不做任何实际操作。对于固定尺寸的矩阵 `Matrix4d m`，做这样的调整`m.resize(4,4)`是不会报错的，因为m并没有改变尺寸。

## 赋值时调整尺寸
Eigen使用等号运算符执行拷贝赋值，Eigen将**自动调整**等号左侧矩阵的尺寸与右侧矩阵相同，从而实现合法的赋值操作。不过如果左侧的矩阵是固定尺寸的且等会两侧矩阵尺寸不同，赋值将会失败。如果你不想Eigen自动调整矩阵尺寸，可以禁止这个特性。

## 定维 vs 动态
什么时候使用定维矩阵？什么时候使用动态矩阵？比较好的答案是：当矩阵尺寸比较小的时候使用定维矩阵，当矩阵尺寸比较大或者你不能使用定维矩阵时使用动态矩阵。对于定维矩阵而言，其执行效率要高于动态矩阵。定维矩阵，Eigen将不会为其分配动态内存。在Eigen中定维矩阵在内存中的实现方式是普通的数组，`Matrix4f m`的实现方式和`float m[16]`是非常类似的，而`MatrixXf mx`的实现方式和`float mx*=new float[rows*cols]`是类似的。

当矩阵元素个数大于32或者更多时，固定尺寸的性能优势将不再明显。创建一个非常大的固定尺寸矩阵很可能造成**栈溢出**，因为Eigen在栈中创建定维矩阵。根据不同的编译与运行环境，Eigen可能使用更激进的向量化手段（例如使用SIMD指令），这些特性可以参考 Vectorization。

## 含默认值的模板参数
文章的开始提到Matrix一共有6个模板参数，到现在为止我们只讨论了其中3个，其余三个有默认参数。下面是Matrix的完整声明：  

```c++
Matrix<typename Scalar,
       int RowsAtCompileTime,
       int ColsAtCompileTime,
       int Options = 0,
       int MaxRowsAtCompileTime = RowsAtCompileTime,
       int MaxColsAtCompileTime = ColsAtCompileTime>
```

* Options ，不同比特位设置不同属性。这里我们只讨论一个比特位：RowMajor。设置这个位，当前矩阵将按照行顺序来保存矩阵，Eigen默认是按列顺序进行保存。
* MaxRowsAtCompileTime 和 MaxColsAtCompileTime， 用于设置矩阵编译时最大尺寸。设置这两个参数最主要的原因是为了避免Eigen动态分配内存。

## 内置类型
Eigen预定义了以下 Matrix 类型：
* MatrixNt ，等价于 Matrix<type, N, N>，示例：MatrixXi=Matrix<int, Dynamic, Dynamic>
* VectorNt，等价于Matrix<type, N, 1>，示例：Vector2f=Marix<float, 2, 1>
* RowVectorNt，等价Matrix<type, 1, N>，示例：RowVector3d=Matrix<double, 1, 3>

说明：
* N 的取值可以是 2、3、4、或者 X，X表示动态
* t 可以是i（int），f（float），d（double），cf（complex<float>），或者cd（complex<double>）。内置类型除了这几种还可以是用户自定义类型


[0]:http://eigen.tuxfamily.org/dox/group__TutorialMatrixClass.html
[1]:http://eigen.tuxfamily.org/dox/group__TopicStorageOrders.html
[2]:http://eigen.tuxfamily.org/dox/group__TutorialAdvancedInitialization.html
[3]:http://eigen.tuxfamily.org/dox/classEigen_1_1PlainObjectBase.html#a712c25be1652e5a64a00f28c8ed11462
