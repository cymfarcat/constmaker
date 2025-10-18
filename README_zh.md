Language: [English](https://github.com/cymfarcat/constmaker/blob/master/README.md) | [中文简体](https://github.com/cymfarcat/constmaker/blob/master/README_zh.md)

# ConstMaker

[ConstMaker‌](https://github.com/cymfarcat/constmaker) 是用来产生常量的工具，通过一次编写常量描述文件，方便快捷的产生 C/C++/C#/Dart/Go/Java/Rust 等近 20 种语言的常量文件在多处使用，增强了代码的可维护性、可读性、安全性。

## 快速入门

ConstMaker 通过**常量描述文件**来描述常量[![参考开发文档]](https://github.com/cymfarcat/constmaker/blob/master/docs/constmaker_zh.md)，如下例子：

````
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

## 生成目标语言

1. 在命令行里运行 constmaker --cpp --upper-ident-camel，产生 C++头文件：

```c++
#include <stdint.h>

/**
 *this is a test
 */
const char* Table = "table"; // hello word
const char* AllowTest = "at";
const float Pi = 3.14159;
const uint8_t Percent = 100;

enum Colors {
    KNoneDef = 0x00,
    KBlueDef = 0x01,
    KGreenDef = 0x02,
    KRedDef = 0x04,
    KAllDef = 0x07,
};
const char* ColorsKNoneDefStr = "NONE";
const char* ColorsKBlueDefStr = "BLUE";
const char* ColorsKGreenDefStr = "GREEN";
const char* ColorsKRedDefStr = "RED";
const char* ColorsKAllDefStr = "ALL";

namespace Test {
    const char* TableId = "table";
    const char* AllowTestId = "at"; // value="at", action=""
    const float PiId = 3.14159;
    const uint8_t PercentId = 100;

    enum class Colors : uint8_t {
        Red,
        Green,
        Blue,
    };
};
````

2. 在命令行里运行 constmaker --dart --upper-ident-camel，产生 Dart 文件：

```dart
/**
 *this is a test
 */
const String Table = "table"; // hello word
const String AllowTest = "at";
const double Pi = 3.14159;
const int Percent = 100;

enum Colors {
    KNoneDef(0x00),
    KBlueDef(0x01),
    KGreenDef(0x02),
    KRedDef(0x04),
    KAllDef(0x07);

    static const String KNoneDefStr = "NONE";
    static const String KBlueDefStr = "BLUE";
    static const String KGreenDefStr = "GREEN";
    static const String KRedDefStr = "RED";
    static const String KAllDefStr = "ALL";

    final int value;
    const Colors(this.value);
}

class Test {
    static const String TableId = "table";
    static const String AllowTestId = "at"; // value="at", action=""
    static const double PiId = 3.14159;
    static const int PercentId = 100;
}

enum TestColors {
    Red,
    Green,
    Blue,
}
```

3. 在命令行里运行 constmaker --java --upper-ident-camel，产生 Java 文件：

```java
package tests;

interface test1 {
    /**
     *this is a test
     */
    public static final String Table = "table"; // hello word
    public static final String AllowTest = "at";
    public static final float Pi = 3.14159f;
    public static final byte Percent = 100;

    enum Colors {
        KNoneDef(0x00),
        KBlueDef(0x01),
        KGreenDef(0x02),
        KRedDef(0x04),
        KAllDef(0x07);

        public static final String KNoneDefStr = "NONE";
        public static final String KBlueDefStr = "BLUE";
        public static final String KGreenDefStr = "GREEN";
        public static final String KRedDefStr = "RED";
        public static final String KAllDefStr = "ALL";

        final int value;
        Colors(int value) { this.value = value; }
    }

    interface Test {
        public static final String TableId = "table";
        public static final String AllowTestId = "at"; // value="at", action=""
        public static final float PiId = 3.14159f;
        public static final byte PercentId = 100;

        enum Colors {
            Red,
            Green,
            Blue
        }
    }
}
```

## SQLite常量
    通过解析SQLite脚本文件，提取表名及字段名结构，生成以字母字符为前缀的随机字符串常量，实现代码逻辑混淆，比如：

```sql
    CREATE TABLE "users" (
        "user_id" INTEGER PRIMARY KEY,      -- Auto-incrementing primary key
        [username] TEXT UNIQUE NOT NULL,    -- Unique username constraint
        [email] TEXT CHECK(email LIKE '%@%'), -- Email format validation
        'created_at' DATETIME DEFAULT CURRENT_TIMESTAMP -- Automatic timestamp
    );
```

    在命令行里运行 constmaker --sqlite --dart --lower-ident，产生 dart 文件：

```dart
const String users = "u0e8qk3";
const String user_id = "s5oqf";
const String username = "lyhi2";
const String email = "po32p0f1";
const String created_at = "g8o1044";

const String create_table_users_0048 = """CREATE TABLE u0e8qk3(
    s5oqf INTEGER PRIMARY KEY,
    lyhi2 TEXT UNIQUE NOT NULL,
    po32p0f1 TEXT CHECK(po32p0f1 LIKE '%@%'),
    g8o1044 DATETIME DEFAULT CURRENT_TIMESTAMP);""";
```

## 例子

请参考[![examples]]目录。

## 命令行参数

ConstMaker 提供了丰富的命令行参数，在命令行里运行 constmaker --help 可获得更多的信息。

## 发布计划

1. 功能修复；
2. 同一常量描述文件内的常量引用；
3. 不同常量描述文件间的常量引用；
4. 通过常量的 ID 获取对应的 STR，通过常量的 STR 获取对应的 ID；
5. 添加新语言支持；
6. ......

## 反馈与建议

请通过 Issue 报告问题或提出改进建议。

## License

Licensed under the Apache License, Version 2.0 (see [LICENSE](https://github.com/cymfarcat/constmaker/blob/master/LICENSE) file).
