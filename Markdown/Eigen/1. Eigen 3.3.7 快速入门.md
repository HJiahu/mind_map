# Eigen 3.3.7 入门教程

学完当前教程后可以参考 [The Matrix class][0] 进行进一步学习。   

## 如何安装 Eigen
因为 Eigen 是 header-only 的，所以直接下载 Eigen 头文件即可使用。

## 一个简单的示例

```c++
#include <iostream>
#include <Eigen/Dense>
using Eigen::MatrixXd;
int main()
{
  MatrixXd m(2,2);
  m(0,0) = 3;
  m(1,0) = 2.5;
  m(0,1) = -1;
  m(1,1) = m(1,0) + m(0,1);
  std::cout << m << std::endl;
}
```
下面先介绍编译方式再解释代码

## 编译并执行
Eigen 不依赖其他三方库，只要把 Eigen 库放到编译器可以发现的路径即可。使用GCC你可以使用`-I`指明Eigen头文件的路径，你可以使用下面的指令编译上面的代码：  
`g++ -I /path/to/eigen/ my_program.cpp -o my_program`  
在linux或者Mac OS X中，你可以将 Eigen 头文件复制到 `usr/local/include` 目录下，这样你就可以使用下面的指令编译上面的代码：  
`g++ my_program.cpp -o my_program`  
编译并执行上面的代码，你将获得如下输出：

``
3   -1
2.5 1.5
``

## 解释第一个程序
Eigen 头文件定义了很多类型，对于简单的使用来说只使用 MatrixXd 类就可以了。MatrixXd 可以表示任意维度的矩阵（MatrixXd中的X表示任意），而且矩阵的元素都是 double 类型的（MatrixXd中的d表示double）。[quick reference guide][1] 中简要介绍了 Eigen 提供的用于描述矩阵的类。

Eigen/Dense头文件定义了MatrixXd类中可用的所有方法和相关类型，所有方法均定义在Eigen名字空间中。

main 函数中的第一行定义了一个2行2列的MatrixXd矩阵，表达式`m(0,0) = 3`将矩阵左上角的元素置为3。在 Eigen 中你可以使用小括号来获得对应元素的引用。在计算机科学中第一个元素的下标为 0，然而在传统的数学表达方法中第一个元素的下标是 1。

## 示例2,：矩阵和向量
下面的例子包含了矩阵和向量，我们先讲第二段代码

```c++
// 代码片段1：运行时设置矩阵的维度
#include <iostream>
#include <Eigen/Dense>
using namespace Eigen;
using namespace std;
int main()
{
  MatrixXd m = MatrixXd::Random(3,3);
  m = (m + MatrixXd::Constant(3,3,1.2)) * 50;
  cout << "m =" << endl << m << endl;
  VectorXd v(3);
  v << 1, 2, 3;
  cout << "m * v =" << endl << m * v << endl;
}

// 代码片段2：编译时确定矩阵的维度
#include <iostream>
#include <Eigen/Dense>
using namespace Eigen;
using namespace std;
int main()
{
  Matrix3d m = Matrix3d::Random();
  m = (m + Matrix3d::Constant(1.2)) * 50;
  cout << "m =" << endl << m << endl;
  Vector3d v(1,2,3);
  
  cout << "m * v =" << endl << m * v << endl;
}
```

第二段代码编译后执行的输出为（因为矩阵是随机初始化的故输出可能不同）：  

```
m =                    
10.1251 90.8741 45.0291
66.3585 68.5009 99.5962
29.3304 57.9873  92.284
m * v =                
326.961                
502.149                
422.157                               
```

## 第二个示例的解释
第二段代码main中声明了一个3×3的矩阵并使用 Random() 进行了随机初始化（初始值在-1和1之间）。第二行代码将m中所有元素的值映射到10和110之间。Constant(3,3,1.2)生成了一个所有元素值为1.2的3×3的矩阵。第三行引入了一个新的类型：VectorXd，这个类表示了一个任意维度**列向量**。第四行代码使用[逗号初始化][2]生成了一个3×1的向量。

最后一行代码执行矩阵和向量的点乘。

上面示例中的两段代码，其功能是相同的。第一段代码使用的向量和矩阵其维度是在运行时决定的，第二段代码中矩阵和向量的形状在编译时就已知了，第二段代码中矩阵和向量的维度是固定的。

使用固定尺寸的矩阵有一定的优势，例如编译器可以生成更高效的可执行文件。使用固定尺寸的类型可以进行更加严格的编译时检查，例如编译器可以检查两个向量是否可以相乘。一个比较好的经验是在矩阵维度小于或等于 4×4 时使用固定尺寸的矩阵（例如：Matrix3f、Matrix2i、Matrix4d等内置类型）。

## 进一步的学习
阅读[long turorial][3]是比较好的选择。



## 其他教程

下面内容源自[CSCI2240][3]  

### 使用模板创建固定维度的矩阵

```c++
Matrix<short, 5, 5> m1;
Matrix<float, 20, 75> m2;
```

### 使用模板创建未知维度的矩阵

```c++
// MatrixXf, MatrixXd
```

### 迭代矩阵
Eigen 在内存中按列顺序保存矩阵，所以在迭代矩阵的时候按列的顺序进行迭代，效率会比按行进行迭代要快（逗号初始化的时候按行），如下伪代码：

```
for i = 1:4 {
	for j = 1:4 {
		B(j , i ) = 0.0;
	}
}
```

### 内置矩阵辅助函数

```c++
// Set each coefficient to a uniform random value in the range [ -1 , 1]
A = Matrix3f :: Random () ;
// Set B to the identity matrix
B = Matrix4d :: Identity () ;
// Set all elements to zero
A = Matrix3f :: Zero () ;
// Set all elements to ones
A = Matrix3f :: Ones () ;
// Set all elements to a constant value
B = Matrix4d :: Constant (4.5) ;
// Scalar multiplication , and subtraction
// What do you expect the output to be?
cout << M2 - Matrix4f :: Ones () * 2.2 << endl ;

```

### 将矩阵看做数组
一些对矩阵的运算需要操作每一个矩阵元素，使用 `array()` 方法可以让 Eigen 将矩阵看做数组，随后的操作都将是 element-wise（以元素为单位进行处理）。`array()`不是**in-place**操作，即 array() 将返回一个和原矩阵完全相同的新矩阵对象。为了效率，Eigen 没有为Matrix类提供element-wise操作，做这些操作需要先将矩阵转化为Array对象，所以自己写代码时要先确定需要哪种对象来完成操作。

```c++
// Square each element of the matrix
cout << M1 . array () . square () << endl ;
// Multiply two matrices element - wise
cout << M1 . array () * Matrix4f :: Identity () . array () << endl ;
// All relational operators can be applied element - wise
cout << M1 . array () <= M2 . array () << endl << endl ;
cout << M1 . array () > M2 . array () << endl ;
```

### 向量

```c++
// Utility functions
Vector3f v1 = Vector3f :: Ones () ;
Vector3f v2 = Vector3f :: Zero () ;
Vector4d v3 = Vector4d :: Random () ;
Vector4d v4 = Vector4d :: Constant (1.8) ;
cout << v1 * v2 . transpose () << endl ;
cout << v1 . dot ( v2 ) << endl << endl ;
cout << v1 . normalized () << endl << endl ;
cout << v1 . cross ( v2 ) << endl ;
cout << v1 . array () * v2 . array () << endl << endl ;
cout << v1 . array () . sin () << endl ;
```



[0]:http://eigen.tuxfamily.org/dox/group__TutorialMatrixClass.html
[1]:http://eigen.tuxfamily.org/dox/group__QuickRefPage.html#QuickRef_Types
[2]:http://eigen.tuxfamily.org/dox/group__TutorialAdvancedInitialization.html
[3]:<https://dritchie.github.io/csci2240/assignments/eigen_tutorial.pdf>




