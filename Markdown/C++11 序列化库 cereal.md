# cereal —— C++11 序列化库

## 介绍

**cereal**是一个只包含头文件的C++序列化库，cereal支持任何类型的数据并可以将其序列化为不同形式，例如：二进制、XML或者JSON。

**cereal**的设计理念是快速、轻量级和容易扩展——cereal没有依赖第三库而且可以轻易的将其和其他代码相。

### cereal 完整支持 C++11
cereal 已经支持 C++11 标准库中的所有类型了，而且 cereal 也完全支持继承和多态。为了保持 cereal 的简洁性并不降低性能，cereal 没有像 Boost 等库那样跟踪并序列化类中所有成员变量。cereal不支持原始指针和引用对象的序列化，但智能指针是支持的。

### cereal 支持众多符合C++11标准的编译器
cereal 使用了很多C++11与编译器的特性。cereal 官方支持 g++4.7.3、clang++3.3、MSVC 2013或者更新的编译器。cereal可能支持其他类型或版本的编译器，比如ICC，但cereal不保证完全可用。在使用g++或者clang++时，cereal可以与libstdc++和libc++一起编译与执行。

### cereal 小巧且迅速
一些简单的性能测试表名，cereal一般比Boost中的序列化库（或者[其他库][1]）要快，并且如果使用二进制文件存储序列化后的对象（数据），cereal占用的空间更小，这些特点在序列化小对象时更加明显。cereal使用了C++中最快的XML和JSON解析库。cereal的代码相对于其他库如Boost而言更加容易理解且更易扩展。

### cereal 是可扩展的
cereal支持将序列化后的对象保存为XML、JSON、和二进制格式。如果需要，你可以添加自己想要的序列化文件格式和其他需要被序列化的数据类型（如自定义类，cereal默认只支持标准库中的类）。

### cereal 有单元测试
为了保证库的可用性与可靠性，我们为cereal编写了单元测试。cereal使用了[Boost 单元测试框架][2]，所以运行这些单元测试需要配置Boost库。

### cereal 很容易使用
cereal的使用是非常简单的，包含头文件并编写序列化语句既可。cereal有优秀的文档来描述其自身的概念和代码。cereal近最大可能在编译时识别并报告你代码中的错误。

### cereal 提供了与Boost类似的语法
如果你使用过Boost中的序列化库，那么你使用cereal是就会感觉到很熟悉，cereal的设计使得熟悉boost的用户容易学习。cereal在对象的方法或者非成员函数中寻找序列化函数。与Boost不同，你不需要告诉cereal你序列化对象的类型。如果你曾经使用过Boost，请查看我们的[迁移手册][3]。

下面是一个示例：

```c++
#include <cereal/types/unordered_map.hpp>
#include <cereal/types/memory.hpp>
#include <cereal/archives/binary.hpp>
#include <fstream>
    
struct MyRecord
{
  uint8_t x, y;
  float z;
  
  template <class Archive>
  void serialize( Archive & ar )
  {
    ar( x, y, z );
  }
};
    
struct SomeData
{
  int32_t id;
  std::shared_ptr<std::unordered_map<uint32_t, MyRecord>> data;
  
  template <class Archive>
  void save( Archive & ar ) const
  {
    ar( data );
  }
      
  template <class Archive>
  void load( Archive & ar )
  {
    static int32_t idGen = 0;
    id = idGen++;
    ar( data );
  }
};

int main()
{
  std::ofstream os("out.cereal", std::ios::binary);
  cereal::BinaryOutputArchive archive( os );

  SomeData myData;
  archive( myData );

  return 0;
}
```

### cereal 在 [BSD][4] 协议下发行



## 快速入门

本小结帮助你在几分钟内获得与编译运行cereal。你唯一需要准备的是一台安装了符合C++11标准的编译器，例如GCC 4.7.3、clang 3.3、MSVC 2013，或者更新的这些编译器更新的版本。低版本的编译器也许可以正常工作，但我们不提供这样的保证。

### 获得 cereal
你可以从[github][5]仓库下载最新的cereal，并将头文件放在任何你编译器可以找到的路径之下。cereal不需要预先构建与编译——cereal是header-only的。

### 向需要序列化的对象中添加成员方法
cereal需要知道对象中的哪些成员是需要序列化的。你需要在对象中实现`serialize`方法来告知cereal：

```c++
struct MyClass
{
  int x, y, z;

  // This method lets cereal know which data members to serialize
  template<class Archive>
  void serialize(Archive & archive)
  {
    // serialize things by passing them to the archive
    archive( x, y, z ); 
  }
};
```
cereal提供了其他更加灵活的序列化函数，你可以在[这里][6]看到有关序列化函数的文档。cereal也支持类版本、私有序列化方法，也支持没有默认构造函数的类。

使用cereal自带的序列化函数，你可以序列化C++中原始数据类型和[几乎所有标准库中定义的类型][7]。

### 选择一种归档（序列化）格式
cereal暂时支持三种序列化文件格式：二进制（支持portable模式）、JSON、XML。XML和JSON方便阅读但也会降低性能（时间与空间上消耗会更大）。

请从下面的头文件中选择一个你喜欢的序列化文件格式：

```c++
#include <cereal/archives/binary.hpp>
#include <cereal/archives/portable_binary.hpp>
#include <cereal/archives/xml.hpp>
#include <cereal/archives/json.hpp>
```

### 序列化你的数据
创建一个cereal**归档对象**并提供你想序列化的类。归档对象使用RAII方式管理资源，cereal保证在归档对象析构的时候将序列化数据写进具体文档中（也可能在析构之前）。创建归档对象时需要一个 `std::istream`或者`std::ostream`对象。

```c++
#include <cereal/archives/binary.hpp>
#include <sstream>

int main()
{
  std::stringstream ss; // any stream can be used

  {
    // Create an output archive
    cereal::BinaryOutputArchive oarchive(ss); 

    MyData m1, m2, m3;
    oarchive(m1, m2, m3); // Write the data to the archive
  } // archive goes out of scope, ensuring all contents are flushed

  {
    cereal::BinaryInputArchive iarchive(ss); // Create an input archive

    MyData m1, m2, m3;
    iarchive(m1, m2, m3); // Read the data from the archive
  }
}
```

***重要提示！***如果你没有阅读上面cereal使用了RAII机制的段落，请仔细阅读一遍。cereal中的一些归档对象只会在其析构的时候将数据写进输入输出流对象（iostream）中。一定要保证，特别是将对象序列化时，你的归档对象会在你使用完成后自动析构，构造序列化数据可能是不完整的。

### 具名变量
cereal支持你为序列化对象命名，在你使用方便阅读的归档格式（XML 或 JSON）时这个特性非常有用：

```c++
#include <cereal/archives/xml.hpp>
#include <fstream>

int main()
{
  {// **注意这列大括号的作用**，大括号之外析构归档对象，将归档数据写进对应文档中
    std::ofstream os("data.xml");
    cereal::XMLOutputArchive archive(os);

    MyData m1;
    int someInt;
    double d;

    archive( CEREAL_NVP(m1), // 序列化变量 m1 时使用变量的原始名称 'm1'
             someInt,        // cereal 将自动为 someInt 定义一个名字
             // 序列化变量 d 时，使用用户自定义的名字
             cereal::make_nvp("this_name_is_way_better", d) );
  }// 析构 archive，将数据写进 os 中

  {
    std::ifstream is("data.xml");
    cereal::XMLInputArchive archive(is);
    
    MyData m1;
    int someInt;
    double d;
    // 反序列化过程不需要指定变量名，也可以指定
    archive( m1, someInt, d );
  }
}
```

更多有关具名变量的信息可以阅读这些[文档][8]。

### 进一步学习
cereal可以做更多更复杂的序列化。cereal支持智能指针、多态、继承等。更多信息请参考官方文档比如，[代码文档][9]和[官方教程][10]。


## cereal 对标准库的支持

cereal当前支持C++中的绝大部分容器和类。引用 cereal 中对应的头文件以实现标准库容易和类的支持（例如：`<cereal/types/vector.hpp>`）。查阅 [doxygen docs][11]查看cereal所支持的所有标准库容器和类。 

### 模板支持

未了序列化一个标准库中的容器或者类，你只需要包含对应的头文件，就像`#include <cereal/types/xxxx.hpp>`这样，然后使用普通的语法：

```c++
// type support
#include <cereal/types/map.hpp>
#include <cereal/types/vector.hpp>
#include <cereal/types/string.hpp>
#include <cereal/types/complex.hpp>

// for doing the actual serialization
#include <cereal/archives/json.hpp>
#include <iostream>

class Stuff
{
  public:
    Stuff() = default;

    void fillData()
    {
      data = { {"real", { {1.0f, 0},
                          {2.2f, 0},
                          {3.3f, 0} }},
               {"imaginary", { {0, -1.0f},
                               {0, -2.9932f},
                               {0, -3.5f} }} };
    }

  private:
    std::map<std::string, std::vector<std::complex<float>>> data;
    
    friend class cereal::access;

    template <class Archive>
    void serialize( Archive & ar )
    {
      ar( CEREAL_NVP(data) );
    }
};

int main()
{
  cereal::JSONOutputArchive output(std::cout); // stream to cout

  Stuff myStuff;
  myStuff.fillData();

  output( cereal::make_nvp("best data ever", myStuff) );
}
```
上面的代码会生成下面的JSON文件：

```c++
{
    "best data ever": {
        "data": [
            {
                "key": "imaginary",
                "value": [
                    {
                        "real": 0,
                        "imag": -1
                    },
                    {
                        "real": 0,
                        "imag": -2.9932
                    },
                    {
                        "real": 0,
                        "imag": -3.5
                    }
                ]
            },
            {
                "key": "real",
                "value": [
                    {
                        "real": 1,
                        "imag": 0
                    },
                    {
                        "real": 2.2,
                        "imag": 0
                    },
                    {
                        "real": 3.3,
                        "imag": 0
                    }
                ]
            }
        ]
    }
}
```

如果你在序列化标准库类型时编译器报错，比如cereal无法找到合适的序列化方法，你可能需要检查一下是否包含了对应的cereal头文件。

更多有关归档对象和序列化函数的信息可以查阅[序列化函数][12]和[归档对象][13]这两小节。



## 从Boost迁移到cereal
略











[1]:http://uscilab.github.io/cereal/assets/coverage/index.html
[2]:https://www.boost.org/doc/libs/1_53_0/libs/test/doc/html/utf.html
[3]:http://uscilab.github.io/cereal/transition_from_boost.html
[4]:https://opensource.org/licenses/BSD-3-Clause
[5]:https://github.com/USCiLab/cereal
[6]:http://uscilab.github.io/cereal/serialization_functions.html
[7]:http://uscilab.github.io/cereal/stl_support.html
[8]:http://uscilab.github.io/cereal/assets/doxygen/group__Utility.html
[9]:http://uscilab.github.io/cereal/assets/doxygen/index.html
[10]:http://uscilab.github.io/cereal/index.html
[11]:http://uscilab.github.io/cereal/assets/doxygen/group__STLSupport.html
[12]:http://uscilab.github.io/cereal/serialization_functions.html
[13]:http://uscilab.github.io/cereal/serialization_archives.html