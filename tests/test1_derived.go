// This file is generated by ConstMaker, DON'T MODIFY.

package tests

const (
    /**
     *this is a test
     */
    Table string = "table" // hello word
    AllowTest string = "at"
    Percent int8 = -100

    /** pi */
    Pi float32 = 3.14159
)

/**
 * colors
 */
// enum Colors
const (
    ColorsKNoneDef uint16 = 0
    ColorsKBlueDef uint16 = (1 << 0)
    ColorsKGreenDef uint16 = (1 << 1)
    ColorsKRedDef uint16 = (1 << 2)
    ColorsKAllDef uint16 = 0x7
    ColorsKNoneDefStr string = "KNoneDef";
    ColorsKBlueDefStr string = "KBlueDef";
    ColorsKGreenDefStr string = "KGreenDef";
    ColorsKRedDefStr string = "KRedDef";
    ColorsKAllDefStr string = "KAllDef";
)

// namespace Node
const (
    NodeHtmlBlockId string = "Block"
    NodeHtmlTableId string = "Table"
)

// namespace Item
const (
    ItemAllowTest string = "at"
    ItemPi float32 = 3.14159
    ItemTable string = "table"
    ItemPercent int8 = 100
    ItemNoneId uint8 = 0;
    ItemAllowTestId uint8 = (1 << 0);
    ItemPiId uint8 = (1 << 1);
    ItemTableId uint8 = (1 << 2);
    ItemPercentId uint8 = (1 << 3);
    ItemAllId uint8 = 0xf;
)

// enum ItemColors
const (
    ItemColorsBlue uint8 = 0
    ItemColorsGreen uint8 = 1 // green
    ItemColorsRed uint8 = 2
)

// namespace ItemNode
const (
    ItemNodeBorder string = "border"
)

// enum ItemNodeColors
const (
    ItemNodeColorsRed uint8 = 0
    ItemNodeColorsGreen uint8 = 1 // green
    ItemNodeColorsBlue uint8 = 2
    ItemNodeColorsRedStr string = "Red";
    ItemNodeColorsGreenStr string = "Green";
    ItemNodeColorsBlueStr string = "Blue";
)

// namespace ItemNodeQuick
const (
    ItemNodeQuickItem string = "item"
)

// namespace ItemNodeQuickNode

// namespace Css3
const (
    Css3XWebkitAirplay string = "x-webkit-airplay"
    Css3WebkitTextZoom string = "-webkit-text-zoom"
    Css3WebkitBackdropFilter string = "-webkit-backdrop-filter"
    Css3WebkitBorderHorizontalSpacing string = "-webkit-border-horizontal-spacing"
    Css3WebkitBorderVerticalSpacing string = "-webkit-border-vertical-spacing"
    Css3WebkitBoxAlign string = "-webkit-box-align"
    Css3NoneId uint8 = 0;
    Css3XWebkitAirplayId uint8 = (1 << 0);
    Css3WebkitTextZoomId uint8 = (1 << 1);
    Css3WebkitBackdropFilterId uint8 = (1 << 2);
    Css3WebkitBorderHorizontalSpacingId uint8 = (1 << 3);
    Css3WebkitBorderVerticalSpacingId uint8 = (1 << 4);
    Css3WebkitBoxAlignId uint8 = (1 << 5);
    Css3AllId uint8 = 0x3f;
)