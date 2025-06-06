Language: [English](https://github.com/cymfarcat/constmaker/blob/master/README.md) | [中文简体](https://github.com/cymfarcat/constmaker/blob/master/README_zh.md)

# ConstMaker

[ConstMaker‌](https://github.com/cymfarcat/constmaker) is a utility designed to generate constants across multiple languages. By writing constant description files once, it conveniently generates constant files for nearly 20 programming languages including C/C++/C#/Dart/Go/Java/Rust, enabling centralized maintenance. This significantly improves code maintainability, readability, and security across projects.

## Quick Start

‌ConstMaker‌ uses ‌constant descriptor files‌ to define constants [![Refer to Dev Docs]](https://github.com/cymfarcat/constmaker/blob/master/docs/constmaker.md), as shown in the following example:

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

## Generate target languages

1. Run constmaker --cpp --upper-ident-camel in the command line to generate C++ header files:

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
```

2. Run constmaker --dart --upper-ident-camel in the command line to generate Dart files:

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

3. Run constmaker --java --upper-ident-camel in the command line to generate Java files:

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

## Examples

‌Please refer to the [![examples]] directory.

## Command-line arguments

ConstMaker provides extensive command-line parameters; run constmaker --help for more information.

## Release Plan

1. Functional fixes‌
2. ‌Constant references within the same descriptor file‌
3. ‌Constant references across different descriptor files‌
4. ‌Retrieve constant STR by ID / Retrieve constant ID by STR‌
5. Add new language support
6. ......

## Feedback and Suggestions

‌Please report issues or propose improvements through Issues.

## License

Licensed under the Apache License, Version 2.0 (see [LICENSE](https://github.com/cymfarcat/constmaker/blob/master/LICENSE) file).
