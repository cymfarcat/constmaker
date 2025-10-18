Language: [English](https://github.com/cymfarcat/constmaker/blob/master/docs/constmaker.md) | [中文简体](https://github.com/cymfarcat/constmaker/blob/master/docs/constmaker_zh.md)

# ConstMaker

This documentation is licensed under [![CC BY-NC-ND 4.0]](https://github.com/cymfarcat/constmaker/blob/master/docs/LICENSE.md).

Using constants in projects reduces code maintenance costs and minimizes runtime errors caused by typos, enhancing code maintainability, readability, and security. Moreover, identical constant values may be required across different projects, various development languages, and server/client development environments.

[ConstMaker‌](https://github.com/cymfarcat/constmaker) is a tool for generating constants. Through a constant definition file, the constmaker CLI tool can produce constant files in multiple languages - enabling "write once, use everywhere" with real-time value updates.

ConstMaker bootstraps itself: The constants used by ConstMaker are generated from [![src/const_define.cmt]](https://github.com/cymfarcat/constmaker/blob/master/src/const_define.cmt) definition file via ConstMaker.

ConstMaker supported languages‌: C/C++/C#/Dart/Go/Java/JavaScript/Kotlin/ObjC/Pascal/Perl/PHP/Python/QML/Ruby/Rust/Swift/TypeScript, and document formats: Markdown/JSON/XML/Text.

## ‌Constant Descriptor File

ConstMaker uses text files with the suffix .cmt to describe constants, thereby abstracting away the details of constant definitions across different programming languages.

1. A constant consists of an identifier, type, and value, defined in a single line and terminated by a newline character;
2. enum defines a collection of integer-type constants;
3. namespace includes composite constants composed of constants and enum, and namespaces can be nested define;
4. enum and namespace essentially serve as grouping mechanisms—their names are not constants themselves, but they can be converted into constant prefixes using the parameters --enum-to-prefix or --namespace-to-prefix;
5. Extends constant definition functionality through optional extended options;
6. Comments.

<a id="const-file-anchor"></a>
Constant Definition Sample

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

### Constant Type

Constants must have explicitly defined types. The types supported by ConstMaker align with those in common development languages:

> | Type | Full   |
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

Using shorthand to reduce input, with ConstMaker mapping types to their corresponding equivalents in development languages.

### Constant Definition

A constant consists of an identifier, a type, and a value, where ‌both the type and the value can be omitted:

1. Define a constant in a single line, terminated by a newline character;
2. Identifier: Composed of letters (A-Z/a-z), digits (0-9), and underscores (\_), where the first character must be a letter or underscore;
3. **ConstMaker allows hyphens (-)‌** which is useful for definitions like "-webkit-text-zoom" see[examples/html];
4. Type omission: ConstMaker will guess the type;
5. When either type or value is omitted: It becomes a string constant with the identifier as its value.

   > | definition        | generated C++ head file             |
   > | ----------------- | ----------------------------------- |
   > | table             | const char\* TABLE = "table";       |
   > | allow_test = "at" | const char\* ALLOW_TEST = "at";     |
   > | pi = 3.14159      | const float PI = 3.14159;           |
   > | percent: u8 = 100 | const uint8_t PERCENT = 100;        |
   > | progress: u8      | const char\* PROGRESS = "progress"; |

### enum Definition

1. Enum definitions, like in most programming languages, represent a ‌set of constants‌ of integer type that can be assigned initial values;
2. Enums cannot contain child defined with the same identifier.

> | definition              | generated C++ head file        |
> | ----------------------- | ------------------------------ |
> | enum Colors : u16 {     | enum class COLORS : uint16_t { |
> | &nbsp;&nbsp;RED = 0x100 | &nbsp;&nbsp;RED = 256,         |
> | &nbsp;&nbsp;GREEN       | &nbsp;&nbsp;GREEN = 257,       |
> | &nbsp;&nbsp;BLUE        | &nbsp;&nbsp;BLUE = 258,        |
> | }                       | };                             |

### nameapce Definition

Namespaces function identically to C++ namespaces – they can contain compound constants (including constants and enums), support nested definitions, and ConstMaker imposes no limit on nesting depth.

[Reference: Constant Definition Sample](#const-file-anchor)

1. **The descriptor file itself constitutes an unnamed namespace**;
2. The scope of constants is determined by their namespace – constants defined with identical identifiers in different namespaces are distinct entities;
3. Additionally, duplicate child identifiers within the same scope are prohibited.

### Extended Option

ConstMaker extends constant definition functionality through optional extended options, thereby facilitating the generation of complex constant files.

1.  Format
    Options are specified within square brackets [] as a comma-separated collection of name:value entries.
    The currently supported option names with ‌prefix/suffix/property/action‌ keywords can be translated as follows:

    > table [prefix: "android", suffix: "ver", property: "freeze|use-define", action: "sha512|md5"]

2.  prefix/suffix option

    - Set prefix/suffix for constant names to reduce manual input;
    - **For enum and namespace**, the prefix/suffix is only applied to child constants；

    > | definition                                      | generated C++ head file                           |
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

    - Configure constant properties for constant names. Currently supported properties include (multiple properties separated by |):
    - **For enum and namespace**, properties are simultaneously applied to every child constant;

    > | Name              | Feature‌                                                                      |
    > | ----------------- | ----------------------------------------------------------------------------- |
    > | freeze            | Read-only property, such as setting readonly in QML                           |
    > | macro-define      | Using macro definitions, for example C/C++ can use macros to define constants |
    > | upper-ident       | Uppercase Identifier                                                          |
    > | lower-ident       | Lowercase Identifier                                                          |
    > | upper-ident-camel | UpperCamelCase‌ Identifier                                                    |
    > | lower-ident-camel | LowerCamelCase‌ Identifier                                                    |

4.  action option
    Define constant actions to perform specific operations on constant values, including:

    - Assist in quickly generating constants, such as bit-flag/gen-str/gen-id, and sort functionality;

    - Enhance constant security by applying measures such as MD5 hashing to string constant values. This hides the original constant values for encryption purposes, for example:

      > table [action: "md5|sha512"]

      > generated definition:
      > const char\* TABLE = "SG1hByraFVpkT0f1571MLNNCIUL2eih5tqqkpan69A5ka+eOP81DHbScRvxEj6PXpFCfDZxlE99rQRachRYqQA=="; // hello word, value=table, action="md5|sha512"

      - For generated string constants, the original plaintext and the generation function used are preserved in comments;
      - Action functions are executed sequentially, with the result of each preceding action serving as input to the next action;
      - If the final action's output lacks base64 or hex encoding, the constant value is automatically encoded in base64.

    - The currently supported actions are (multiple actions separated by |):

    > | Name          | Feature‌                                                        | Usage            |
    > | ------------- | --------------------------------------------------------------- | ---------------- |
    > | bit-flag      | Generate Left-Shift Display for 8421-Encoded Integer Constants  | enum+namsapce    |
    > | bit-flag-hex  | Generate Hexadecimal Display for 8421-Encoded Integer Constants | enum+namsapce    |
    > | gen-str       | Generate String Constants with \_STR Suffix                     | enum             |
    > | gen-id        | Generate Integer Constants with \_ID Suffix                     | namsapce         |
    > | sort-asc      | Sort in ascending alphabetical order                            | enum+namsapce    |
    > | sort-desc     | Sort in descending alphabetical order                           | enum+namsapce    |
    > | sort-asc-lex  | Sort in ascending lexicographical order                         | enum+namsapce    |
    > | sort-desc-lex | Sort in descending lexicographical order                        | enum+namsapce    |
    > | md5           | md5 function                                                    | String Constants |
    > | sha256        | sha256 function                                                 | String Constants |
    > | sha384        | sha384 function                                                 | String Constants |
    > | sha512        | sha512 function                                                 | String Constants |
    > | base64        | base64 function                                                 | String Constants |
    > | hex           | hex function                                                    | String Constants |

5.  enum option
    The enum's options control enum children, including bit-flag/gen-str/sort. For example:

    > enum Colors : u16 [prefix: "", property: "", action: "bit-flag|gen-str|sort-asc"] {
    > &nbsp;&nbsp;RED = 0x100
    > &nbsp;&nbsp;GREEN
    > &nbsp;&nbsp;BLUE
    > }

    > The generated constants: identifiers are sorted alphabetically:
    > enum class COLORS : uint16_t {
    > &nbsp;&nbsp;BLUE = 1 << 1,
    > &nbsp;&nbsp;GREEN = 1 << 2,
    > &nbsp;&nbsp;RED = 1 << 3,
    > };
    > const char* COLORS_BLUE_STR = "BLUE";
    > const char* COLORS_GREEN_STR = "GREEN";
    > const char\* COLORS_RED_STR = "RED";

6.  namespace option
    The namespace's options control namespace children, including bit-flag/gen-id/sort. For example:

    > namespace ITEM [prefix: "", action: "bit-flag|gen-id|sort-asc-lex"] {
    > &nbsp;&nbsp;table
    > &nbsp;&nbsp;allow_test = "at"
    > &nbsp;&nbsp;\_percent = 100
    > &nbsp;&nbsp;pi = 3.14159
    > }

    > Generated constants: identifiers are sorted lexicographically:
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

### Comments

1. Comments are exclusively for the constant descriptor file and will not appear in generated files:

   > |        |               |
   > | ------ | ------------- |
   > | /\*\*/ | Block Comment |
   > | //     | Line comment  |

   > /\* this is a test
   > \*/
   > table [action: ""] // hello word

   > Generated constants:
   > const char\* TABLE = "table";

2. Comments must be included in generated files:

   > |          |               |
   > | -------- | ------------- |
   > | /\*\*\*/ | Block Comment |
   > | ///      | Line comment  |

   > /\*\* this is a test
   > \*/
   > table [action: ""] /// hello word

   > Generated constants:
   > /\*\* this is a test
   > \*/
   > const char\* TABLE = "table"; // hello word

## SQLite Consts

    Parse SQLite scripts to extract table/field metadata, then generate random letter-prefixed string constants for code obfuscation.

1. Supports standard SQLite script files, please refer to the [![tests/sqlite]] directory;

2. Currently supports generating source files in Dart/Go/Cpp/Rust for easy project integration, and can also create new SQL files;

3. For each SQLite statement in different programming languages, a unique identifier is required:

    - Automatically generated by ConstMaker based on the Statement, for example:

      create_table_users_0048，select_users_0049，select_posts_users_0056

      The trailing number serves as a global increment to prevent identifier duplication.

      ***Note‌***: adding or removing statements in the script file will alter this number, remember to update all code references accordingly.

    - Annotate Statements with their identifiers in comments to maintain uniform references throughout the codebase. The system currently recognizes two comment styles:

      - Short name: --- #[stmt(short = "post_id")]
        Use this name to replace the trailing number mentioned above:
        
        ``` sql
        --- #[stmt(short = "post_id")]
        UPDATE posts SET views = views + 1 WHERE post_id = 1 LIMIT 1; -- Safe update
        ```
        Generated Dart definition:
        ``` dart
        const String update_limited_posts_post_id = "UPDATE ps3jo SET wlejia = wlejia + 1 WHERE ievwi = 1 LIMIT 1;";
        ```

      - Full name: --- #[stmt(full = "alter_user_bio")]
        Use this name as an identifier:
        
        ``` sql
        --- #[stmt(full = "alter_user_bio")]
        ALTER TABLE users ADD COLUMN bio TEXT DEFAULT ''; -- Add new column
        ```
        Generated Dart definition:
        ``` dart
        const String alter_user_bio = "ALTER TABLE u2mm2m0 ADD COLUMN kn31xf3b TEXT DEFAULT '';";
        ```

    - The statement uses only adjacent comments to generate identifiers.

4. Command-Line
    - name min(--min-name): default = 4.
    - name max(--max-name): default = 8.
    - field white list(--field-wl): semicolon-separated string using non-randomized original field names.

5. Comments
    - Comments are exclusively for the SQLite script file and will not appear in generated files:
    > |        |        |
    > | ------ | ------ |
    > | /\*\*/ | Block Comment |
    > | --     | Line comment |

    - Comments must be included in generated files:
    > |          |        |
    > | -------- | ------ |
    > | /\*\*\*/ | Block Comment |
    > | ---      | Line comment |

    ***Note‌***：Do not add block comments (/***/) or line comments (---) in the Statement, as they may cause parsing errors.

6. Some suggestions:
    - In the same database, field names should remain unique across different tables.
    - The tables are defined in different script files to prevent field name conflicts.

## Command-Line Reference

The ConstMaker command-line provides a comprehensive set of commands for generating constant files. The currently supported parameters are as follows:

> | Short Param | Long Param            | Function                                                                        |
> | ----------- | --------------------- | ------------------------------------------------------------------------------- |
> | -i          | --input string        | input file to const define.                                                     |
> | -o          | --output string       | output path to generated files.                                                 |
> |             | --file-ext string     | output file name extension.                                                     |
> |             | --file-name string    | output file name.                                                               |
> |             | --check-value         | check for duplicate constant values within the current scope.                   |
> | -v          | --verbose             | print verbose info.                                                             |
> |             | --prefix string       | prefix for every const ident name.                                              |
> |             | --suffix string       | suffix for every const ident name.                                              |
> |             | --root-name string    | root name, if empty substitute with file name.                                  |
> |             | --root-namespace      | forece root node into namspace.                                                 |
> |             | --root-ns             | forece root node into namspace.                                                 |
> |             | --enum-to-prefix      | use enum name as prefix.                                                        |
> |             | --namespace-to-prefix | use namespace as prefix.                                                        |
> |             | --ns-to-prefix        | use namespace as prefix.                                                        |
> |             | --freeze              | force use readonly to declare constants.                                        |
> |             | --nested-define       | force namespace use nested define like tree, or else use flat define like list. |
> |             | --macro-define        | force use macro define to declare constants.                                    |
> |             | --upper-ident         | force ident to upper.                                                           |
> |             | --lower-ident         | force ident to lower.                                                           |
> |             | --upper-ident-camel   | force ident to upper camel case.                                                |
> |             | --lower-ident-camel   | force ident to lower camel case.                                                |
> |             | --bitflag-all string  | bitflag generate All string. (default "ALL")                                    |
> |             | --bitflag-none string | bitflag generate None string. (default "NONE")                                  |
> |             | --bitflag-none-all    | bitflag generate None and All option. (default true)                            |
> | -t          | --tab int             | tab width. (default 4)                                                          |
> | -a          | --all                 | generate all supported language files.                                          |
> | -c          | --cpp                 | generate C/CPP head files.                                                      |
> |             | --std-cpp11           | std c++11. (default true)                                                       |
> |             | --c#                  | generate C# files.                                                              |
> | -d          | --dart                | generate Dart files.                                                            |
> | -g          | --go                  | generate Go files.                                                              |
> |             | --go-package string   | generate Go package.                                                            |
> | -j          | --java                | generate Java files.                                                            |
> |             | --java-package string | generate Java package.                                                          |
> |             | --javascript          | generate JavaScript files.                                                      |
> |             | --json                | generate Json files.                                                            |
> |             | --kotlin              | generate Kotlin files.                                                          |
> |             | --markdown            | generate Markdown files.                                                        |
> |             | --objc                | generate Objective-C files.                                                     |
> |             | --pascal              | generate Pascal files.                                                          |
> |             | --perl                | generate Perl files.                                                            |
> |             | --php                 | generate Php files.                                                             |
> | -p          | --python              | generate Python files.                                                          |
> |             | --qml                 | generate QML files.                                                             |
> |             | --qml-singleton       | generate QML singleton.                                                         |
> |             | --ruby                | generate Ruby files.                                                            |
> | -r          | --rust                | generate Rust files.                                                            |
> | -s          | --swift               | generate Swift files.                                                           |
> |             | --text                | generate Text files.                                                            |
> |             | --typescript          | generate TypeScript files.                                                      |
> |             | --xml                 | generate XML files.                                                             |
> |             | --sqlite              | use sqlite consts.                                                              |
> |             | --sql                 | generate sql files.                                                             |
> |             | --use-prefix          | force name use prefix.                                                          |
> |             | --min-name            | min name length.                                                                |
> |             | --max-name            | max name length.                                                                |
> |             | --table               | prefix for every table name.                                                    |
> |             | --index               | prefix for every index name.                                                    |
> |             | --view                | prefix for every trigger name.                                                  |
> |             | --trigger             | prefix for every trigger name.                                                  |
> |             | --vtable              | prefix for every virtual table name.                                            |
> |             | --field-wl            | semicolon-separated whitelist field.                                            |
