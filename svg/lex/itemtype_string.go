// Code generated by "stringer -type=itemType"; DO NOT EDIT.

package lex

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ItemError-0]
	_ = x[ItemSyntaxError-1]
	_ = x[ItemColon-2]
	_ = x[ItemDot-3]
	_ = x[ItemOpenBracket-4]
	_ = x[ItemCloseBracket-5]
	_ = x[ItemSemiColon-6]
	_ = x[ItemComma-7]
	_ = x[ItemHEXColour-8]
	_ = x[ItemRGBColour-9]
	_ = x[ItemAttribute-10]
	_ = x[ItemNumber-11]
	_ = x[ItemUnit-12]
	_ = x[ItemText-13]
	_ = x[ItemWhitespace-14]
	_ = x[ItemEOF-15]
}

const _itemType_name = "ItemErrorItemSyntaxErrorItemColonItemDotItemOpenBracketItemCloseBracketItemSemiColonItemCommaItemHEXColourItemRGBColourItemAttributeItemNumberItemUnitItemTextItemWhitespaceItemEOF"

var _itemType_index = [...]uint8{0, 9, 24, 33, 40, 55, 71, 84, 93, 106, 119, 132, 142, 150, 158, 172, 179}

func (i itemType) String() string {
	if i < 0 || i >= itemType(len(_itemType_index)-1) {
		return "itemType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _itemType_name[_itemType_index[i]:_itemType_index[i+1]]
}
