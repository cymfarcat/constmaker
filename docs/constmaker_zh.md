Language: [English](https://github.com/cymfarcat/constmaker/blob/master/docs/constmaker.md) | [中文简体](https://github.com/cymfarcat/constmaker/blob/master/docs/constmaker_zh.md)

# ConstMaker

本文档采用[![CC BY-NC-ND 4.0]](https://github.com/cymfarcat/constmaker/blob/master/docs/LICENSE.md)协议。

在项目中使用常量可降低代码维护成本，减少拼写错误导致的运行异常，可以增强代码的可维护性、可读性、安全性。不仅如此，在不同的项目，不同的开发语言，服务器/客户端开发过程中，可能需要使用值相同的常量。

[ConstMaker‌](https://github.com/cymfarcat/constmaker) 就是用来产生常量的工具，通过常量描述文件，使用 constmaker 命令行工具可以产生多种语言的常量文件，做到了一次编写，多处使用，且随时可以修改常量的值，方便快捷。

ConstMaker 实现了自举：ConstMaker 使用的常量，是通过常量描述文件 [![src/const_define.cmt]](https://github.com/cymfarcat/constmaker/blob/master/src/const_define.cmt)，由 ConstMaker 生成的。

ConstMaker 当前支持的开发语言有：C/C++/C#/Dart/Go/Java/JavaScript//Kotlin/ObjC/Pascal/Perl/Php/Python/Qml/Ruby/Rust/Swift/TypeScript, 还包括普通文件 Markdown/Json/Xml/Text。

## 常量描述文件

ConstMaker 通过后缀为 cmt 的文本文件来描述常量，这样就屏蔽了不同开发语言的常量定义细节。

1. 常量由标识符、类型和值组成，单行定义，换行符结束；
2. enum 定义整数类型的集合常量；
3. namespace 包括由常量、enum 组成的复合常量，namespace 可以嵌套定义；
4. enum 和 namespace 本质上是起分组作用，其名称并不是常量，但可以通过参数--enum-to-prefix 或--namespace-to-prefix 转化为常量前缀；
5. 通过添加可选的 option，用来扩展常量定义功能;
6. 注释；

<a id="const-file-anchor"></a>
常量定义示例

```
/**
 *this is a test
 */
table [action: ""] /// hello word
allow_test = "at"
pi = 3.14159
percent: u8 = 100

enum Colors [prefix: "K", suffix: "Def", property: "", action: "bit-flag-hex|gen-str|sort-asc"] {
  RED
  GREEN
  BLUE
}

namespace Test [prefix: "", suffix: "ID"] {
    table
    allow_test = "at" [prefix: "", suffix: "STR", action: ""]
    pi = 3.14159
    percent: u8 = 100

    enum Colors {
        RED
        GREEN
        BLUE
    }
}
```

### 常量类型

常量需要明确其类型。ConstMaker 支持的类型和大多数开发语言一致：

> | 类型 | 说明   |
> | ---- | ------ |
> | bool | bool   |
> | u8   | uint8  |
> | i8   | int8   |
> | u16  | uint16 |
> | i16  | int16  |
> | u32  | uint32 |
> | i32  | int32  |
> | u64  | uint64 |
> | i64  | int64  |
> | f32  | float  |
> | f64  | double |
> | str  | string |

使用短字符便于减少输入，ConstMaker 会将类型映射到开发语言相对应的类型。

### 常量定义

常量由标识符、类型和值组成，其中类型和值都可以省略：

1. 单行定义一个常量，换行符结束；
2. 标识符：由字母（A-Z/a-z）、数字（0-9）、下划线（\_），且首字符必须为字母或下划线；
3. **ConstMaker 允许 hyphens (-)‌ 字符存在**， 这对于类似"-webkit-text-zoom"的定义很有用，参见[examples/html];
4. 类型省略：ConstMaker 会猜测其类型；
5. 类型或值省略：这就是字符串常量，其值为标识符。

   > | 定义              | 产生的 C++头文件                    |
   > | ----------------- | ----------------------------------- |
   > | table             | const char\* TABLE = "table";       |
   > | allow_test = "at" | const char\* ALLOW_TEST = "at";     |
   > | pi = 3.14159      | const float PI = 3.14159;           |
   > | percent: u8 = 100 | const uint8_t PERCENT = 100;        |
   > | progress: u8      | const char\* PROGRESS = "progress"; |

### enum 定义

1. enum 定义和大多数开发语言一样，是一组常量的集合，其类型为整数，可以设定初值;
2. enum 不能定义相同标识符的子常量。

> | 定义                    | 产生的 C++头文件               |
> | ----------------------- | ------------------------------ |
> | enum Colors : u16 {     | enum class COLORS : uint16_t { |
> | &nbsp;&nbsp;RED = 0x100 | &nbsp;&nbsp;RED = 256,         |
> | &nbsp;&nbsp;GREEN       | &nbsp;&nbsp;GREEN = 257,       |
> | &nbsp;&nbsp;BLUE        | &nbsp;&nbsp;BLUE = 258,        |
> | }                       | };                             |

### nameapce 定义

nameapce 和 C++的 nameapce 一样，可以包括由常量、enum 组成的复合常量，namespace 可以嵌套定义，ConstMaker 并没有限制嵌套层数。

[参考常量定义示例](#const-file-anchor)

1. **描述文件自身就是一个未命名的 nameapce。**
2. 常量的作用域由 nameapce 来决定，即使在不同 nameapce 里定义相同的标识符，也是不同的常量；
3. 在同一个作用域里不能定义相同标识符的子常量。

### 扩展 option

ConstMaker 通过添加可选的 option，用来扩展常量定义功能，方便生成复杂的常量文件。

1.  格式
    option 是放在[]内的一组选项，选项之间以,分隔，选项使用 名称:值 的格式。
    当前支持 prefix/suffix/property/action 关键字的选项名称，比如：

    > table [prefix: "android", suffix: "ver", property: "freeze|use-define", action: "sha512|md5"]

2.  prefix/suffix option

    - 设置常量名称的前缀/后缀，便于减少输入；
    - **对 enum 和 namespace**, 只会将前缀/后缀设置给每一个子常量；

    > | 定义                                            | 产生的 C++头文件                                  |
    > | ----------------------------------------------- | ------------------------------------------------- |
    > | table [prefix: "android"，suffix: "ver"]        | const char\* android_TABLE_ver = "table";         |
    > |                                                 |                                                   |
    > | enum Colors [prefix: "Circle"] {                | enum class COLORS : uint8{                        |
    > | &nbsp;&nbsp;RED = 0x100                         | &nbsp;&nbsp;Circle_RED = 256,                     |
    > | &nbsp;&nbsp;GREEN                               | &nbsp;&nbsp;Circle_GREEN = 257,                   |
    > | }                                               | };                                                |
    > | namespace Node [prefix: "html", suffix: "ID"] { | namespace NODE {                                  |
    > | &nbsp;&nbsp;Block                               | &nbsp;&nbsp;const char\* html_BLOCK_ID = "Block"; |
    > | &nbsp;&nbsp;Table                               | &nbsp;&nbsp;const char\* html_TABLE_ID = "Table"; |
    > | }                                               | };                                                |

3.  property option

    - 设置常量的属性，针对常量名称，当前支持的属性有（多个属性使用|分隔）：
    - **对 enum 和 namespace**, 同时将属性设置给每一个子常量；

    > | 名称              | 功能                                     |
    > | ----------------- | ---------------------------------------- |
    > | freeze            | 只读属性，比如 QML 可设置 readonly       |
    > | macro-define      | 使用宏定义，比如 C/C++可以使用宏定义常量 |
    > | upper-ident       | 大写标识符                               |
    > | lower-ident       | 小写标识符                               |
    > | upper-ident-camel | 大驼峰标识符                             |
    > | lower-ident-camel | 小驼峰标识符                             |

4.  action option
    设置常量动作，针对常量值执行特定操作，包括：

    - 辅助快捷生成常量，比如 bit-flag/gen-str/gen-id 和排序功能；

    - 提高常量安全性，比如对字符串常量值进行 MD5，这样就可以隐藏原始常量值，达到加密目的，比如：

      > table [action: "md5|sha512"]

      > 生成常量定义：
      > const char\* TABLE = "SG1hByraFVpkT0f1571MLNNCIUL2eih5tqqkpan69A5ka+eOP81DHbScRvxEj6PXpFCfDZxlE99rQRachRYqQA=="; // hello word, value=table, action="md5|sha512"

      - 生成的字符串常量，使用注释保留了原始明文，以及使用的函数；
      - 动作函数计算顺序是前一个动作结果作为后一个动作的输入；
      - 如果最后的动作没有 base64 或 hex，自动使用 base64 编码的常量值。

    - 当前支持的动作有（多个动作使用|分隔）：

    > | 名称          | 功能                                  | 适用范围      |
    > | ------------- | ------------------------------------- | ------------- |
    > | bit-flag      | 生成 8421 编码的整型常量-向左移位显示 | enum+namsapce |
    > | bit-flag-hex  | 生成 8421 编码的整型常量-16 进制显示  | enum+namsapce |
    > | gen-str       | 生成带\_STR 的字符串常量              | enum          |
    > | gen-id        | 生成带\_ID 的整型常量                 | namsapce      |
    > | sort-asc      | 按字母表升序排序                      | enum+namsapce |
    > | sort-desc     | 按字母表逆序排序                      | enum+namsapce |
    > | sort-asc-lex  | 按字典序升序排序                      | enum+namsapce |
    > | sort-desc-lex | 按字典序升序排序                      | enum+namsapce |
    > | md5           | md5 函数                              | 字符常量值    |
    > | sha256        | sha256 函数                           | 字符常量值    |
    > | sha384        | sha384 函数                           | 字符常量值    |
    > | sha512        | sha512 函数                           | 字符常量值    |
    > | base64        | base64 函数                           | 字符常量值    |
    > | hex           | hex 函数                              | 字符常量值    |

5.  enum option
    enum 的 option 控制 enum 子常量，包括 bit-flag/gen-str/排序，比如:

    > enum Colors : u16 [prefix: "", property: "", action: "bit-flag|gen-str|sort-asc"] {
    > &nbsp;&nbsp;RED = 0x100
    > &nbsp;&nbsp;GREEN
    > &nbsp;&nbsp;BLUE
    > }

    > 生成的常量，标识符已经按字母表排序：
    > enum class COLORS : uint16_t {
    > &nbsp;&nbsp;BLUE = 1 << 1,
    > &nbsp;&nbsp;GREEN = 1 << 2,
    > &nbsp;&nbsp;RED = 1 << 3,
    > };
    > const char* COLORS_BLUE_STR = "BLUE";
    > const char* COLORS_GREEN_STR = "GREEN";
    > const char\* COLORS_RED_STR = "RED";

6.  namespace option
    namespace 的 option 控制 namespace 子常量，包括 bit-flag/gen-id/排序，比如:

    > namespace ITEM [prefix: "", action: "bit-flag|gen-id|sort-asc-lex"] {
    > &nbsp;&nbsp;table
    > &nbsp;&nbsp;allow_test = "at"
    > &nbsp;&nbsp;\_percent = 100
    > &nbsp;&nbsp;pi = 3.14159
    > }

    > 生成的常量，标识符已经按字典序排序：
    > namespace ITEM {
    > &nbsp;&nbsp;const char* ALLOW_TEST = "at";
    > &nbsp;&nbsp;const float PI = 3.14159;
    > &nbsp;&nbsp;const char* TABLE = "table";
    > &nbsp;&nbsp;const int8_t \_PERCENT = 100;
    > &nbsp;&nbsp;const uint8_t ALLOW_TEST_ID = 1 << 1;
    > &nbsp;&nbsp;const uint8_t PI_ID = 1 << 2;
    > &nbsp;&nbsp;const uint8_t TABLE_ID = 1 << 3;
    > &nbsp;&nbsp;const uint8_t \_PERCENT_ID = 1 << 4;
    > }

### 注释

1. 注释仅用于常量描述文件，不会出现在生成的文件里:

   > |        |        |
   > | ------ | ------ |
   > | /\*\*/ | 块注释 |
   > | //     | 行注释 |

   > /\* this is a test
   > \*/
   > table [action: ""] // hello word

   > 生成的常量：
   > const char\* TABLE = "table";

2. 注释需要出现在生成的文件里：

   > |          |        |
   > | -------- | ------ |
   > | /\*\*\*/ | 块注释 |
   > | ///      | 行注释 |

   > /\*\* this is a test
   > \*/
   > table [action: ""] /// hello word

   > 生成的常量：
   > /\*\* this is a test
   > \*/
   > const char\* TABLE = "table"; // hello word

## SQLite常量

    通过解析SQLite脚本文件，提取表名及字段名结构，生成以字母字符为前缀的随机字符串常量，实现代码逻辑混淆。

1. 支持标准的SQLite脚本文件，请参考[![tests/sqlite]]目录；

2. 当前支持生成dart/go/cpp/rust语言的源文件，方便集成到项目中使用，也可以生成新的sql文件；

3. 对不同的开发语言，每一个SQLite Statement‌语句，都需要一个标识符：

    - 由ConstMaker根据Statement‌自动生成，比如：

      create_table_users_0048，select_users_0049，select_posts_users_0056

      末尾数字是全局增量，用来防止标识符重复；

      ***注意***：如果脚本文件的Statement‌有增加/删除，该数字会发生变化，因此需要修改引用该标识符的代码。

    - 对Statement添加注释，来提供该语句的标识符，这样可以保证代码中都使用相同的标识符，当前支持两种注释：

      - 短名称：--- #[stmt(short = "post_id")]
        使用该名称来替换上述的末尾数字：
        
        ``` sql
        --- #[stmt(short = "post_id")]
        UPDATE posts SET views = views + 1 WHERE post_id = 1 LIMIT 1; -- Safe update
        ```
        生成的Dart定义为：
        ``` dart
        const String update_limited_posts_post_id = "UPDATE ps3jo SET wlejia = wlejia + 1 WHERE ievwi = 1 LIMIT 1;";
        ```

      - 全名称：--- #[stmt(full = "alter_user_bio")]
        使用该名称来用作标识符：
        
        ``` sql
        --- #[stmt(full = "alter_user_bio")]
        ALTER TABLE users ADD COLUMN bio TEXT DEFAULT ''; -- Add new column
        ```
        生成的Dart定义为：
        ``` dart
        const String alter_user_bio = "ALTER TABLE u2mm2m0 ADD COLUMN kn31xf3b TEXT DEFAULT '';";
        ```

    - Statement仅使用相邻的注释来生成标识符。

4. 命令行
    - 名称最小值(--min-name)：默认值为4；
    - 名称最大值(--max-name)：默认值为8；
    - 字段白名单(--field-wl)：使用";"分割的字符串，不使用随机数的原始字段名称；

5. 注释
    - 注释仅用于SQL文件，不会出现在生成的文件里:
    > |        |        |
    > | ------ | ------ |
    > | /\*\*/ | 块注释 |
    > | --     | 行注释 |

    - 注释需要出现在生成的文件里：
    > |          |        |
    > | -------- | ------ |
    > | /\*\*\*/ | 块注释 |
    > | ---      | 行注释 |

    ***注意***：不要在Statement里添加块注释(/\*\*\*/)和行注释(---)，会造成解析异常；

6. 一些建议：
    - 同一个数据库中，不同的表字段名称应保持唯一性；
    - 不同的表定义在不同的脚本文件里，避免字段名冲突；

## 命令行参考

ConstMaker 的命令行提供了丰富的命令，用来生成常量文件，当前支持的命令行参数，如下所示：

> | 短参数 | 长参数                | 功能                                                                            |
> | ------ | --------------------- | ------------------------------------------------------------------------------- |
> | -i     | --input string        | input file to const define.                                                     |
> | -o     | --output string       | output path to generated files.                                                 |
> |        | --file-ext string     | output file name extension.                                                     |
> |        | --file-name string    | output file name.                                                               |
> |        | --check-value         | check for duplicate constant values within the current scope.                   |
> | -v     | --verbose             | print verbose info.                                                             |
> |        | --prefix string       | prefix for every const ident name.                                              |
> |        | --suffix string       | suffix for every const ident name.                                              |
> |        | --root-name string    | root name, if empty substitute with file name.                                  |
> |        | --root-namespace      | forece root node into namspace.                                                 |
> |        | --root-ns             | forece root node into namspace.                                                 |
> |        | --enum-to-prefix      | use enum name as prefix.                                                        |
> |        | --namespace-to-prefix | use namespace as prefix.                                                        |
> |        | --ns-to-prefix        | use namespace as prefix.                                                        |
> |        | --freeze              | force use readonly to declare constants.                                        |
> |        | --nested-define       | force namespace use nested define like tree, or else use flat define like list. |
> |        | --macro-define        | force use macro define to declare constants.                                    |
> |        | --upper-ident         | force ident to upper.                                                           |
> |        | --lower-ident         | force ident to lower.                                                           |
> |        | --upper-ident-camel   | force ident to upper camel case.                                                |
> |        | --lower-ident-camel   | force ident to lower camel case.                                                |
> |        | --bitflag-all string  | bitflag generate All string. (default "ALL")                                    |
> |        | --bitflag-none string | bitflag generate None string. (default "NONE")                                  |
> |        | --bitflag-none-all    | bitflag generate None and All option. (default true)                            |
> | -t     | --tab int             | tab width. (default 4)                                                          |
> | -a     | --all                 | generate all supported language files.                                          |
> | -c     | --cpp                 | generate C/CPP head files.                                                      |
> |        | --std-cpp11           | std c++11. (default true)                                                       |
> |        | --c#                  | generate C# files.                                                              |
> | -d     | --dart                | generate Dart files.                                                            |
> | -g     | --go                  | generate Go files.                                                              |
> |        | --go-package string   | generate Go package.                                                            |
> | -j     | --java                | generate Java files.                                                            |
> |        | --java-package string | generate Java package.                                                          |
> |        | --javascript          | generate JavaScript files.                                                      |
> |        | --json                | generate Json files.                                                            |
> |        | --kotlin              | generate Kotlin files.                                                          |
> |        | --markdown            | generate Markdown files.                                                        |
> |        | --objc                | generate Objective-C files.                                                     |
> |        | --pascal              | generate Pascal files.                                                          |
> |        | --perl                | generate Perl files.                                                            |
> |        | --php                 | generate Php files.                                                             |
> | -p     | --python              | generate Python files.                                                          |
> |        | --qml                 | generate QML files.                                                             |
> |        | --qml-singleton       | generate QML singleton.                                                         |
> |        | --ruby                | generate Ruby files.                                                            |
> | -r     | --rust                | generate Rust files.                                                            |
> | -s     | --swift               | generate Swift files.                                                           |
> |        | --text                | generate Text files.                                                            |
> |        | --typescript          | generate TypeScript files.                                                      |
> |        | --xml                 | generate XML files.                                                             |
> |        | --sqlite              | use sqlite consts.                                                              |
> |        | --sql                 | generate sql files.                                                             |
> |        | --use-prefix          | force name use prefix.                                                          |
> |        | --min-name            | min name length.                                                                |
> |        | --max-name            | max name length.                                                                |
> |        | --table               | prefix for every table name.                                                    |
> |        | --index               | prefix for every index name.                                                    |
> |        | --view                | prefix for every trigger name.                                                  |
> |        | --trigger             | prefix for every trigger name.                                                  |
> |        | --vtable              | prefix for every virtual table name.                                            |
> |        | --field-wl            | semicolon-separated whitelist field.                                            |
