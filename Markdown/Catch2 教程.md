# Catch 使用教程（入门，官方文档翻译）

`译者注：当前文档并不是官方文档的直译。在翻译的过程中我删除部分原文中的内容并添加了一些自己的理解，可能有偏差，请见谅`



1. 获得 Catch
2. 如何使用？
3. 编写测试用例 
4. 测试用例和测试区段
5. BDD-Style
6. 小结
7. 参数类型化测试
8. 后续学习与使用

## 获得 Catch

获得Catch最简单的方式是下载最新的 [single header version](https://raw.githubusercontent.com/catchorg/Catch2/master/single_include/catch2/catch.hpp)。这个头文件由若干其他独立的头文件合并而成。

你也可以使用其他方法获得Catch，例如使用CMake来构建编译版Catch，这可以提高项目的编译速度。

完整的Catch包含测试、说明文档等内容，你可以从GitHub下载完整的Catch。Catch官方链接为：[http://catch-lib.net](http://catch-lib.net) ，此链接将重定向到GitHub。

## 如何使用 Catch?

Catch是header-only的，故你只需要将Catch的头文件放到编译器可以发现的路径既可。  

下面的教程默认你的编译器可以发现并使用 Catch。

_如果你使用Catch的预编译形式，即已经编译并生成了Catch链接库（.lib 或者 .a 文件），你的Catch头文件包含形式应该形如：`#include <catch2/catch.hpp>` 。_

## 编写测试用例

让我们从一个简单的示例开始(examples/010-TestCase.cpp)。假设你已经写了一个用于计算阶乘的函数，现在准备测试它。（TDD的基本原则是先写测试代码，为了方便学习，这里先忽略这个原则）

```c++
unsigned int Factorial( unsigned int number ) {
    return number <= 1 ? number : Factorial(number-1)*number;
}
```

为了尽量简单，我们把所有的代码都放到一个源文件中。

```c++
#define CATCH_CONFIG_MAIN  // 当前宏强制Catch在当前编译单元中创建 main()，这个宏只能出现在一个CPP文件中，因为一个项目只能有一个有效的main函数
#include "catch.hpp"

unsigned int Factorial( unsigned int number ) {
    return number <= 1 ? number : Factorial(number-1)*number;
}

TEST_CASE( "Factorials are computed", "[factorial]" ) {
    REQUIRE( Factorial(1) == 1 );
    REQUIRE( Factorial(2) == 2 );
    REQUIRE( Factorial(3) == 6 );
    REQUIRE( Factorial(10) == 3628800 );
}
```

编译结束后将生成一个可以接受运行时参数的可执行文件，具体可用参数请参考command-line.md。如果以不带参数的方式执行可执行文件，所有测试用例都将被执行。详细的测试报告将输出到终端，测试报告包含失败的测试用例、失败的测试用例个数、成功的测试用例个数等信息。

执行上面代码生成的可执行文件，所有测试用例都将通过。真的没有错误码？不是的，上面的阶乘函数是有错误的，我写的第一版教程中就有这个Bug，感谢CTMacUser帮我指出了这个错误。

这个错误是什么呢？0的阶乘是多少？——0的阶乘是1而不是0，这就是上面阶乘函数的错误之处。参考：
[0的阶乘是1](http://mathforum.org/library/drmath/view/57128.html)

让我们把上面的规则写入到测试用例中：

```c++
TEST_CASE( "Factorials are computed", "[factorial]" ) {
    REQUIRE( Factorial(0) == 1 );
    REQUIRE( Factorial(1) == 1 );
    REQUIRE( Factorial(2) == 2 );
    REQUIRE( Factorial(3) == 6 );
    REQUIRE( Factorial(10) == 3628800 );
}
```

现在测试失败了，Catch输出：

```
Example.cpp:9: FAILED:
  REQUIRE( Factorial(0) == 1 )
with expansion:
  0 == 1
```

Catch的测试报告会输出期望值和Factorial(0)计算出的错误值0，这样我们就可以很方便的找到错误。

让我们修正阶乘函数:

```c++
unsigned int Factorial( unsigned int number ) {
  return number > 1 ? Factorial(number-1)*number : 1;
}
```

现在所有的测试用例都通过了。

当然了上面的阶乘函数依旧有不少问题，例如当number很大时计算的结果将溢出，不过我们暂不管这些。

### 我们做了什么?

虽然上面的测试比较简单，但已经足够展示如何使用Catch了。在进一步学习前，我们先解释一下上面那段代码。

1. 我们定义了一个宏，并包含了Catch的头文件，然后编译这个源文件并生成了一个接受运行时参数的可执行文件。为了可执行，定义了宏`#define CATCH_CONFIG_MAIN`，强制Catch引入预定义main函数，你也可以编写自己的main函数（参考：own-main.md）。

2. 我们在宏`TEST_CASE`中编写测试用例。这个宏可以包含一个或者两个参数，其中一个参数是没有固定格式的测试名，另一个参数则包含一个或多个标签（下文介绍）。测试名必须唯一。参考command-line.md以获得更多有关执行可执行文件的信息。

3. 测试名和标签都是字符串。

4. 我们仅使用宏`REQUIRE`来编写测试断言。Catch没有使用分立的测试函数表示不同的断言（例如REQUIRE_TRUE、REQUIRE_FALSE、REQUIRE_EQUAL、REQUIRE_LESS等），而是直接使用C++表达式的真值结果。此外Catch使用模板表达式捕获测试表达式的左侧和右侧（例如 `exp_a == exp_b`，Catch将捕获exp_a和exp_b的计算结果），从而在测试报告中显示两侧的计算结果。

## 测试用例和测试区段（Test case and section）

大部分测试框架都有某种基于类的机制。例如，在很多框架（例如JUnit）的`setup()`阶段可以创建一个在其他用例中使用的测试对象（可以是需要测试的对象，也可以是Mock对象），在`teardown()`阶段销毁这些对象，从而避免在每一个测试用例中创建与销毁测试对象（或mock对象）。

对于Catch而言，使用上面传统的测试方式有一定的缺陷，例如对于同一批测试用例你只能创建同一个测试对象，这样的话测试粒度就比较大。（译者注：其他缺陷可以参考原文）

Catch 使用全新的方式解决了上面的问题，如下：

```c++
TEST_CASE( "vectors can be sized and resized", "[vector]" ) {

    std::vector<int> v( 5 );

    REQUIRE( v.size() == 5 );
    REQUIRE( v.capacity() >= 5 );

    SECTION( "resizing bigger changes size and capacity" ) {
        v.resize( 10 );

        REQUIRE( v.size() == 10 );
        REQUIRE( v.capacity() >= 10 );
    }
    SECTION( "resizing smaller changes size but not capacity" ) {
        v.resize( 0 );

        REQUIRE( v.size() == 0 );
        REQUIRE( v.capacity() >= 5 );
    }
    SECTION( "reserving bigger changes capacity but not size" ) {
        v.reserve( 10 );

        REQUIRE( v.size() == 5 );
        REQUIRE( v.capacity() >= 10 );
    }
    SECTION( "reserving smaller does not change size or capacity" ) {
        v.reserve( 0 );

        REQUIRE( v.size() == 5 );
        REQUIRE( v.capacity() >= 5 );
    }
}
```

对于每一个`SECTION`，`TEST_CASE`都将重新从当前`TEST_CASE`的起始部分开始执行并忽略其他`SECTION`。  （译者注：这段原文简单解释了原因，Catch使用了if语句并把section看做子节点，每次执行TEST_CASE时Catch先执行起始部分的非`SECTION`代码，然后选择一个子节点并执行）。

到目前为止，Catch使用上述方式已经实现了大部分测试框架基于类（setup&teardown）的测试机制。

`SECTION`可以嵌套任意深度，每一个`SECTION`子节点都只会被执行一次，大量嵌套的`SECTION`会形成一棵“树”，父节点执行失败将不再执行对应的子节点：

```c++
    SECTION( "reserving bigger changes capacity but not size" ) {
        v.reserve( 10 );

        REQUIRE( v.size() == 5 );
        REQUIRE( v.capacity() >= 10 );

        SECTION( "reserving smaller again does not change capacity" ) {
            v.reserve( 7 );

            REQUIRE( v.capacity() >= 10 );
        }
    }
```

## BDD-Style

Catch可以使用BDD-Style形式的测试，具体请参考：test-cases-and-sections.md，下面是一个简单的例子：

```c++
SCENARIO( "vectors can be sized and resized", "[vector]" ) {

    GIVEN( "A vector with some items" ) {
        std::vector<int> v( 5 );

        REQUIRE( v.size() == 5 );
        REQUIRE( v.capacity() >= 5 );

        WHEN( "the size is increased" ) {
            v.resize( 10 );

            THEN( "the size and capacity change" ) {
                REQUIRE( v.size() == 10 );
                REQUIRE( v.capacity() >= 10 );
            }
        }
        WHEN( "the size is reduced" ) {
            v.resize( 0 );

            THEN( "the size changes but not capacity" ) {
                REQUIRE( v.size() == 0 );
                REQUIRE( v.capacity() >= 5 );
            }
        }
        WHEN( "more capacity is reserved" ) {
            v.reserve( 10 );

            THEN( "the capacity changes but not the size" ) {
                REQUIRE( v.size() == 5 );
                REQUIRE( v.capacity() >= 10 );
            }
        }
        WHEN( "less capacity is reserved" ) {
            v.reserve( 0 );

            THEN( "neither size nor capacity are changed" ) {
                REQUIRE( v.size() == 5 );
                REQUIRE( v.capacity() >= 5 );
            }
        }
    }
}
```

运行上面的测试用例将输出以下内容:

```
Scenario: vectors can be sized and resized
     Given: A vector with some items
      When: more capacity is reserved
      Then: the capacity changes but not the size
```

## 小结

为了保证教程的简洁性我们把所有代码放在了一个文件中，在实际项目中这并不是好的方式。

比较好的方式是将下面这段代码写在一个独立的源文件中，其他测试文件仅包含Catch头文件和测试代码。不要在其他测试文件中重复包含下面的`#define`语句。

```c++
#define CATCH_CONFIG_MAIN
#include "catch.hpp"
```

**不要在头文件中写测试代码！**


## 类型参数化测试

Catch支持类型参数化测试，宏`TEMPLATE_TEST_CASE`和`TEMPLATE_PRODUCT_TEST_CASE`的行为和`TEST_CASE`类似，但测试用例会在不同类型下执行。下面代码中`TestType`的取值依次为`int`、`float`、`std::string`、`Bar`，所有测试用例都将在这些类型下执行一遍。

```c++
struct Bar {}; 

TEMPLATE_TEST_CASE("Templated test","",int,float, std::string, Bar)
{
    std::vector<TestType> v( 5 );
    REQUIRE( v.size() == 5 );
    REQUIRE( v.capacity() >= 5 );
}
```
更多信息请参考：test-cases-and-sections.md中type-parametrised-test-cases小节

## 后续学习与使用

当前文档简要介绍了Catch，也指出了Catch和其他测试框架的一些区别。了解这些知识后你已经可以编写一些实际的测试用例了。

当然还有很多东西需要学习，但你只需要在用到那些新特性的时候再学。你可以在 Readme.md 中找到Catch的所有特性。
