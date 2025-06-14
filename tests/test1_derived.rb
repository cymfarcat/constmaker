# This file is generated by ConstMaker, DON'T MODIFY.

# 
# this is a test
# 
Table = "table" # hello word
AllowTest = "at"
Percent = -100
# pi 
Pi = 3.14159

# 
# colors
# 
module Colors
    KNoneDef = 0
    KBlueDef = (1 << 0)
    KGreenDef = (1 << 1)
    KRedDef = (1 << 2)
    KAllDef = 0x7
    ColorsKNoneDefStr = "KNoneDef"
    ColorsKBlueDefStr = "KBlueDef"
    ColorsKGreenDefStr = "KGreenDef"
    ColorsKRedDefStr = "KRedDef"
    ColorsKAllDefStr = "KAllDef"
end

module Node
    HtmlBlockId = "Block"
    HtmlTableId = "Table"
end

module Item
    AllowTest = "at"
    Pi = 3.14159
    Table = "table"
    Percent = 100
    NoneId = 0
    AllowTestId = (1 << 0)
    PiId = (1 << 1)
    TableId = (1 << 2)
    PercentId = (1 << 3)
    AllId = 0xf

    module Colors
        Blue = 0
        Green = 1 # green
        Red = 2
    end
end

module ItemNode
    Border = "border"

    module Colors
        Red = 0
        Green = 1 # green
        Blue = 2
        ColorsRedStr = "Red"
        ColorsGreenStr = "Green"
        ColorsBlueStr = "Blue"
    end
end

module ItemNodeQuick
    Item = "item"
end

module ItemNodeQuickNode
end

module Css3
    XWebkitAirplay = "x-webkit-airplay"
    WebkitTextZoom = "-webkit-text-zoom"
    WebkitBackdropFilter = "-webkit-backdrop-filter"
    WebkitBorderHorizontalSpacing = "-webkit-border-horizontal-spacing"
    WebkitBorderVerticalSpacing = "-webkit-border-vertical-spacing"
    WebkitBoxAlign = "-webkit-box-align"
    NoneId = 0
    XWebkitAirplayId = (1 << 0)
    WebkitTextZoomId = (1 << 1)
    WebkitBackdropFilterId = (1 << 2)
    WebkitBorderHorizontalSpacingId = (1 << 3)
    WebkitBorderVerticalSpacingId = (1 << 4)
    WebkitBoxAlignId = (1 << 5)
    AllId = 0x3f
end