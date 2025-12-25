// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"

	src "github.com/cymfarcat/constmaker/src"
)

const (
	TypeObj       uint8 = 1
	TypeEnum      uint8 = 2
	TypeEnumValue uint8 = 3
	TypeNamespace uint8 = 4
	TypeProperty  uint8 = 5
)

// ////////////////////////////////////////////////////////////////////////////
type ConstObject struct {
	Prefix        string
	Suffix        string
	Ident         string `json:"id" xml:"id,attr"`
	Typo          string `json:"type" xml:"type,attr"`
	Value         string `json:"value" xml:"value,attr"`
	value         string
	CommentDoc    string
	CommentTriple string         `json:"comment" xml:"comment,attr"`
	Children      []*ConstObject `json:"consts,omitempty" xml:"const"`
	genObjs       []*ConstObject

	parent  *ConstObject
	objType uint8
	options map[string]*ConstObject

	EnumToPrefix      bool
	NamespaceToPrefix bool
	freeze            bool
	MacroDefine       bool
	IdentMapStr       bool
	StrMapIdent       bool
	StrMapQName       bool
	UpperIdent        bool
	LowerIdent        bool
	UpperIdentCamel   bool
	LowerIdentCamel   bool

	enumSetValue bool
}

func NewConstObject(parent *ConstObject, objType uint8) *ConstObject {
	obj := &ConstObject{}
	obj.initialize(parent, objType)
	return obj
}

func (obj *ConstObject) cleanup() {}

func (obj *ConstObject) initialize(parent *ConstObject, objType uint8) {
	obj.Children = make([]*ConstObject, 0)
	obj.genObjs = make([]*ConstObject, 0)
	obj.options = make(map[string]*ConstObject)
	obj.parent = parent
	obj.objType = objType
}

// check current scope identifier is redefined
func (obj *ConstObject) identRedefined(object *ConstObject, line int, column int) bool {
	for _, constObj := range obj.Children {
		if constObj.Ident == object.Ident {
			fmt.Fprintf(os.Stderr, "Identifier '%s' is redefined at %d : %d.\n", object.Ident, line, column+1)
			return true
		}
	}
	return false
}

func (obj *ConstObject) addObject(object *ConstObject, line int, column int) {
	if object.value == "space-all" {
		return
	}
	if !obj.identRedefined(object, line, column) {
		obj.Children = append(obj.Children, object)
	}
}

func (obj *ConstObject) addOption(option *ConstObject) {
	obj.options[option.Ident] = option
}

func (obj *ConstObject) sortChildren(sortAsc bool, sortDesc bool, sortAscLex bool, sortDescLex bool) {
	if sortAsc {
		sort.Slice(obj.Children, func(i, j int) bool {
			return obj.Children[i].Ident < obj.Children[j].Ident
		})
	}

	if sortDesc {
		sort.Slice(obj.Children, func(i, j int) bool {
			return obj.Children[i].Ident > obj.Children[j].Ident
		})
	}

	if sortAscLex {
		sort.Slice(obj.Children, func(i, j int) bool {
			a, b := obj.Children[i].Ident, obj.Children[j].Ident
			isAlphaA := len(a) > 0 && unicode.IsLetter(rune(a[0]))
			isAlphaB := len(b) > 0 && unicode.IsLetter(rune(b[0]))

			switch {
			case isAlphaA && !isAlphaB:
				return true
			case !isAlphaA && isAlphaB:
				return false
			default:
				return a < b
			}
		})
	}

	if sortDescLex {
		sort.Slice(obj.Children, func(i, j int) bool {
			a, b := obj.Children[i].Ident, obj.Children[j].Ident
			isAlphaA := len(a) > 0 && unicode.IsLetter(rune(a[0]))
			isAlphaB := len(b) > 0 && unicode.IsLetter(rune(b[0]))

			switch {
			case isAlphaA && !isAlphaB:
				return true
			case !isAlphaA && isAlphaB:
				return false
			default:
				return a > b
			}
		})
	}
}

func (obj *ConstObject) CheckValue() {
	valueMap := make(map[string]string)

	for _, constObj := range obj.Children {

		ident, exists := valueMap[constObj.Value]
		if exists {
			fmt.Fprintf(os.Stderr, "Warning: identifier '%s' and '%s' has duplicate value: %s.\n", constObj.Ident, ident, constObj.Value)
		}

		valueMap[constObj.Value] = constObj.Ident
	}
}

func (obj *ConstObject) inheritedProperty(parent *ConstObject) {
	obj.Prefix = parent.Prefix
	obj.Suffix = parent.Suffix
	obj.freeze = parent.freeze
	obj.MacroDefine = parent.MacroDefine
	obj.IdentMapStr = parent.IdentMapStr
	obj.StrMapIdent = parent.StrMapIdent
	obj.StrMapQName = parent.StrMapQName
	obj.UpperIdent = parent.UpperIdent
	obj.LowerIdent = parent.LowerIdent
	obj.UpperIdentCamel = parent.UpperIdentCamel
	obj.LowerIdentCamel = parent.LowerIdentCamel
}

/*
 * apply property only for children
 */
func (obj *ConstObject) applyProperty() {
	if len(obj.Children) == 0 {
		return
	}

	// children use parent's prefix
	if len(obj.Prefix) > 0 {
		for _, constObj := range obj.Children {
			constObj.Prefix = obj.Prefix
		}
	}

	// children use parent's suffix
	if len(obj.Suffix) > 0 {
		for _, constObj := range obj.Children {
			constObj.Suffix = obj.Suffix
		}
	}

	// if obj freeze, children need freeze
	if obj.freeze {
		for _, constObj := range obj.Children {
			constObj.freeze = obj.freeze
		}
	}

	// if obj useDefine, children need useDefine
	if obj.MacroDefine {
		for _, constObj := range obj.Children {
			constObj.MacroDefine = obj.MacroDefine
		}
	}

	if obj.IdentMapStr {
		for _, constObj := range obj.Children {
			constObj.IdentMapStr = obj.IdentMapStr
		}
	}

	if obj.StrMapIdent {
		for _, constObj := range obj.Children {
			constObj.StrMapIdent = obj.StrMapIdent
		}
	}

	if obj.StrMapQName {
		for _, constObj := range obj.Children {
			constObj.StrMapQName = obj.StrMapQName
		}
	}

	// if obj upperIdent, children need upperIdent
	if obj.UpperIdent {
		for _, constObj := range obj.Children {
			constObj.UpperIdent = obj.UpperIdent
		}
	}

	// if obj lowerIdent, children need lowerIdent
	if obj.LowerIdent {
		for _, constObj := range obj.Children {
			constObj.LowerIdent = obj.LowerIdent
		}
	}

	// if obj cameCaseIdent, children need cameCaseIdent
	if obj.UpperIdentCamel {
		for _, constObj := range obj.Children {
			constObj.UpperIdentCamel = obj.UpperIdentCamel
		}
	}

	// if obj snakeCaseIdent, children need snakeCaseIdent
	if obj.LowerIdentCamel {
		for _, constObj := range obj.Children {
			constObj.LowerIdentCamel = obj.LowerIdentCamel
		}
	}
}

func (obj *ConstObject) clearProperty() {
	obj.Prefix = ""
	obj.Suffix = ""
	// obj.freeze = false
	// obj.MacroDefine = false
	// obj.IdentMapStr = false
	// obj.StrMapIdent = false
	// obj.StrMapQName = false
	// obj.upperIdent = false
	// obj.lowerIdent = false
	// obj.upperIdentCame = false
	// obj.lowerIdentCame = false
}

func (obj *ConstObject) detectType(str string) {
	if len(str) == 0 {
		obj.Typo = src.Typo_Str
		return
	} else if intVal, err := strconv.ParseInt(str, 0, 64); err == nil { // first ParseInt
		obj.Typo = intType(intVal)
		obj.Value = strconv.FormatInt(intVal, 10)
	} else if uintVal, err := strconv.ParseUint(str, 0, 64); err == nil { // then ParseUint
		obj.Typo = uintType(uintVal)
		obj.Value = strconv.FormatUint(uintVal, 10)
	} else if floatVal, err := strconv.ParseFloat(str, 32); err == nil { // then ParseFloat32
		if !strings.ContainsAny(str, "eE.") && floatVal == math.Trunc(floatVal) {
			obj.detectType(fmt.Sprintf("%.0f", floatVal))
		} else {
			obj.Typo = src.Typo_F32
			obj.Value = strconv.FormatFloat(floatVal, 'f', -1, 32)
		}
	} else if floatVal, err := strconv.ParseFloat(str, 64); err == nil { // then ParseFloat64
		if !strings.ContainsAny(str, "eE.") && floatVal == math.Trunc(floatVal) {
			obj.detectType(fmt.Sprintf("%.0f", floatVal))
		} else {
			obj.Typo = src.Typo_F64
			obj.Value = strconv.FormatFloat(floatVal, 'f', -1, 64)
		}
	} else {
		obj.Typo = src.Typo_Str
		obj.Value = str
	}
}

func (obj *ConstObject) genValue(option *src.Options) {
	// apply action option
	bitFlag := false
	bitFlagHex := false
	toUpper := false
	toLower := false
	genEntity := false
	genStr := false
	genId := false

	sortAsc := false
	sortDesc := false
	sortAscLex := false
	sortDescLex := false

	// apply prefix option
	prefix, ok := obj.options[src.Option_Prefix]
	if ok && len(prefix.Value) > 0 {
		obj.Prefix = src.UnquoteStr(prefix.Value)
	}

	// apply suffix option
	suffix, ok := obj.options[src.Option_Suffix]
	if ok && len(suffix.Value) > 0 {
		obj.Suffix = src.UnquoteStr(suffix.Value)
	}

	action, ok := obj.options[src.Option_Action]
	if ok && len(action.Value) > 0 {
		str := strings.ToLower(src.UnquoteStr(action.Value))

		if strings.Contains(str, src.Option_BitFlag) {
			bitFlag = true
		}

		if strings.Contains(str, src.Option_BitFlagHex) {
			bitFlagHex = true
		}

		if strings.Contains(str, src.Option_ToUpper) {
			toUpper = true
		}

		if strings.Contains(str, src.Option_ToLower) {
			toLower = true
		}

		if strings.Contains(str, src.Option_GenEntity) {
			genEntity = true
		}

		if strings.Contains(str, src.Option_GenStr) {
			genStr = true
		}

		if strings.Contains(str, src.Option_GenId) {
			genId = true
		}

		if strings.Contains(str, src.Option_SortAscLex) {
			sortAscLex = true
		} else if strings.Contains(str, src.Option_SortAsc) {
			sortAsc = true
		}

		if strings.Contains(str, src.Option_SortDescLex) {
			sortDescLex = true
		} else if strings.Contains(str, src.Option_SortDesc) {
			sortDesc = true
		}
	}

	// first genEntity
	if genEntity {
		for _, constObj := range obj.Children {
			constObj.Value = strconv.Quote(src.UnicodeToString(constObj.Value))
			constObj.value = obj.Prefix + constObj.value + obj.Suffix
		}

		// clear Prefix/Suffix
		obj.Prefix = ""
		obj.Suffix = ""
	}

	// first change toUpper/toLower
	if toUpper || toLower {
		for _, constObj := range obj.Children {
			if toLower {
				constObj.Ident = strings.ToLower(constObj.Ident)
				constObj.value = strings.ToLower(constObj.value)
			} else if toUpper {
				constObj.Ident = strings.ToUpper(constObj.Ident)
				constObj.value = strings.ToUpper(constObj.value)
			}
		}
	}

	// correct indent
	correctIdent(option, obj)

	// apply property option
	obj.EnumToPrefix = false
	obj.NamespaceToPrefix = false
	obj.freeze = false
	obj.MacroDefine = false
	obj.IdentMapStr = false
	obj.StrMapIdent = false
	obj.StrMapQName = false
	obj.UpperIdent = false
	obj.LowerIdent = false
	obj.UpperIdentCamel = false
	obj.LowerIdentCamel = false

	property, ok := obj.options[src.Option_Property]
	if ok && len(property.Value) > 0 {
		str := strings.ToLower(src.UnquoteStr(property.Value))

		if strings.Contains(str, src.Option_EnumToPrefix) {
			obj.EnumToPrefix = true
		}

		if strings.Contains(str, src.Option_NamespaceToPrefix) || strings.Contains(str, src.Option_NSToPrefix) {
			obj.NamespaceToPrefix = true
		}

		if strings.Contains(str, src.Option_Freeze) {
			obj.freeze = true
		}

		if strings.Contains(str, src.Option_MacroDefine) {
			obj.MacroDefine = true
		}

		if strings.Contains(str, src.Option_IdentMapStr) {
			obj.IdentMapStr = true
		}

		if strings.Contains(str, src.Option_StrMapIdent) {
			obj.StrMapIdent = true
		}

		if strings.Contains(str, src.Option_strMapQName) {
			obj.StrMapQName = true
		}

		if strings.Contains(str, src.Option_UpperIdent) {
			obj.UpperIdent = true
		}

		if strings.Contains(str, src.Option_LowerIdent) {
			obj.LowerIdent = true
		}

		if strings.Contains(str, src.Option_UpperIdentCamel) {
			obj.UpperIdentCamel = true
		}

		if strings.Contains(str, src.Option_LowerIdentCamel) {
			obj.LowerIdentCamel = true
		}
	}

	// pre-check bitFlag||bitFlagHex
	if (bitFlag || bitFlagHex) && len(obj.Children) > int(src.BitFlagMax) {
		fmt.Fprintf(os.Stderr, "\"%s\" sub-items exceeds the maximum value of the type, can't exec bitflag.\n", obj.Ident)
		bitFlag = false
		bitFlagHex = false
	}

	// calc max bitflag
	bitFlagCount := bitFlagCount(adjustLen(len(obj.Children), option.BitFlagNoneAll))

	if obj.objType == TypeEnum {
		// do sort
		obj.sortChildren(sortAsc, sortDesc, sortAscLex, sortDescLex)

		// apply property to all children
		obj.applyProperty()

		if !genEntity {
			// first check type
			typo := obj.Typo

			// check enum if set value
			enumOrderValue := ""
			obj.enumSetValue = false
			for _, constObj := range obj.Children {
				if len(constObj.Value) != 0 {
					obj.enumSetValue = true

					// if has sort, keep last order value and clear constObj.Value
					if sortAsc || sortDesc || sortAscLex || sortDescLex {
						enumOrderValue = constObj.Value
						constObj.Value = ""
					} else {
						break
					}
				}
			}

			// if has sort, need change first order value
			if (sortAsc || sortDesc || sortAscLex || sortDescLex) && len(enumOrderValue) > 0 && len(obj.Children) > 0 {
				obj.Children[0].Value = enumOrderValue
			}

			// gen enum value
			if bitFlag || bitFlagHex {
				obj.enumSetValue = true

				if len(typo) == 0 {
					typo = bitFlagType(uint64(len(obj.Children)), option.BitFlagNoneAll)
				}

				allValue := 0
				for idx, constObj := range obj.Children {
					constObj.Typo = typo
					constObj.Value = "1 << " + strconv.FormatInt(int64(idx), 10)
					if bitFlagHex && bitFlagCount != 0 {
						constObj.Value = fmt.Sprintf("0x%0*X", bitFlagCount, 1<<idx)
					}
					constObj.value = constObj.Value
					allValue |= 1 << idx
				}

				if option.BitFlagNoneAll {
					// add None object
					none := NewConstObject(obj, TypeEnumValue)
					none.inheritedProperty(obj)
					none.Ident = option.BitFlagNone
					none.Typo = typo
					none.Value = "0"
					if bitFlagHex && bitFlagCount != 0 {
						none.Value = fmt.Sprintf("0x%0*X", bitFlagCount, 0)
					}
					none.value = option.BitFlagNone
					obj.Children = append([]*ConstObject{none}, obj.Children...)

					// add All object
					all := NewConstObject(obj, TypeEnumValue)
					all.inheritedProperty(obj)
					all.Ident = option.BitFlagAll
					all.Typo = typo
					all.Value = "0x" + strconv.FormatInt(int64(allValue), 16)
					if bitFlagHex && bitFlagCount != 0 {
						all.Value = fmt.Sprintf("0x%0*X", bitFlagCount, allValue)
					}
					all.value = option.BitFlagAll
					obj.Children = append(obj.Children, all)
				}
			} else {
				if len(typo) == 0 {
					typo = uintType(uint64(len(obj.Children)))
				}

				var lastValue int64 = 0
				hasNegativeNum := false

				if obj.enumSetValue {
					// enum generate value
					for _, constObj := range obj.Children {
						constObj.Typo = typo

						if len(constObj.Value) == 0 {
							constObj.Value = strconv.FormatInt(int64(lastValue), 10)
						} else {
							intVal, err := strconv.ParseInt(constObj.Value, 0, 64)
							if err == nil {
								lastValue = intVal

								//check is negative
								if intVal < 0 {
									hasNegativeNum = true
								}
							}
							constObj.Value = strconv.FormatInt(int64(lastValue), 10)
						}
						lastValue++
					}
				} else {
					// enum default value
					for _, constObj := range obj.Children {
						constObj.Value = strconv.FormatInt(int64(lastValue), 10)
						lastValue++
					}
				}

				// reset enum typo
				if hasNegativeNum {
					typo = intType(lastValue)
				} else {
					typo = uintType(uint64(lastValue))
				}

				obj.Typo = typo
				for _, constObj := range obj.Children {
					constObj.Typo = typo
				}
			}
		}

		// gen str
		if genStr && len(obj.Children) > 0 {
			for _, constObj := range obj.Children {
				genObj := NewConstObject(obj, TypeObj)
				genObj.inheritedProperty(obj)
				genObj.Ident = constObj.Ident + src.UnderScore + src.ConstSTR
				if len(genObj.Suffix) > 0 {
					// always use constObj.Ident
					// genObj.Ident = constObj.Ident + src.UnderScore + genObj.Suffix + src.UnderScore + src.ConstSTR
					genObj.Suffix = ""
				}
				genObj.Typo = src.Typo_Str

				// always use constObj.value
				genObj.Value = strconv.Quote(constObj.value) // strconv.Quote(genIdentName(option, "", constObj))

				obj.genObjs = append(obj.genObjs, genObj)
			}
		}

		// check duplicate value
		if option.CheckValue {
			obj.CheckValue()
		}

		// clear self property
		obj.clearProperty()
	} else if obj.objType == TypeEnumValue {
	} else if obj.objType == TypeNamespace {
		// do sort
		obj.sortChildren(sortAsc, sortDesc, sortAscLex, sortDescLex)

		// apply property to all children
		obj.applyProperty()

		// gen id
		if genId && len(obj.Children) > 0 {
			typo := bitFlagType(uint64(len(obj.Children)), option.BitFlagNoneAll)

			allValue := 0
			for idx, constObj := range obj.Children {
				genObj := NewConstObject(obj, TypeObj)
				genObj.inheritedProperty(obj)
				genObj.Ident = constObj.Ident + src.UnderScore + src.ConstID
				if len(genObj.Suffix) > 0 {
					genObj.Ident = constObj.Ident + src.UnderScore + genObj.Suffix + src.UnderScore + src.ConstID
					genObj.Suffix = ""
				}
				genObj.Typo = typo

				if bitFlag || bitFlagHex {
					if option.BitFlagNoneAll && idx == 0 {
						// add None object
						none := NewConstObject(obj, TypeEnumValue)
						none.inheritedProperty(obj)
						none.Ident = option.BitFlagNone + src.UnderScore + src.ConstID
						if len(none.Suffix) > 0 {
							none.Ident = option.BitFlagNone + src.UnderScore + none.Suffix + src.UnderScore + src.ConstID
							none.Suffix = ""
						}
						none.Typo = typo
						none.Value = "0"
						if bitFlagHex && bitFlagCount != 0 {
							none.Value = fmt.Sprintf("0x%0*X", bitFlagCount, 0)
						}
						obj.genObjs = append(obj.genObjs, none)
					}

					genObj.Value = "1 << " + strconv.FormatInt(int64(idx), 10)
					if bitFlagHex && bitFlagCount != 0 {
						genObj.Value = fmt.Sprintf("0x%0*X", bitFlagCount, 1<<idx)
					}

					obj.genObjs = append(obj.genObjs, genObj)

					allValue |= 1 << idx

					if option.BitFlagNoneAll && idx == len(obj.Children)-1 {
						// add All object
						all := NewConstObject(obj, TypeEnumValue)
						all.inheritedProperty(obj)
						all.Ident = option.BitFlagAll + src.UnderScore + src.ConstID
						if len(all.Suffix) > 0 {
							all.Ident = option.BitFlagAll + src.UnderScore + all.Suffix + src.UnderScore + src.ConstID
							all.Suffix = ""
						}
						all.Typo = typo
						all.Value = "0x" + strconv.FormatInt(int64(allValue), 16)
						if bitFlagHex && bitFlagCount != 0 {
							all.Value = fmt.Sprintf("0x%0*X", bitFlagCount, allValue)
						}
						obj.genObjs = append(obj.genObjs, all)
					}
				} else {
					genObj.Value = strconv.FormatInt(int64(idx+1), 10)

					obj.genObjs = append(obj.genObjs, genObj)
				}
			}
		}

		// gen str
		if genStr && len(obj.Children) > 0 {
			for _, constObj := range obj.Children {
				genObj := NewConstObject(obj, TypeObj)
				genObj.inheritedProperty(obj)
				genObj.Ident = constObj.Ident + src.UnderScore + src.ConstSTR
				if len(genObj.Suffix) > 0 {
					// always use constObj.Ident
					// genObj.Ident = constObj.Ident + src.UnderScore + genObj.Suffix + src.UnderScore + src.ConstSTR
					genObj.Suffix = ""
				}
				genObj.Typo = src.Typo_Str

				// always use constObj.value
				genObj.Value = strconv.Quote(constObj.value) // strconv.Quote(genIdentName(option, "", constObj))

				obj.genObjs = append(obj.genObjs, genObj)
			}
		}

		// check duplicate value
		if option.CheckValue {
			obj.CheckValue()
		}

		// clear self property
		obj.clearProperty()
	} else {
		// first detect value and type
		if len(obj.Typo) == 0 {
			obj.detectType(obj.Value)
		}

		// check value
		if len(obj.Value) == 0 {
			obj.Typo = src.Typo_Str

			// if no value, use value as Value
			obj.Value = obj.value
		}

		if len(obj.Typo) == 0 {
			obj.Typo = src.Typo_Str
		}

		if obj.Typo == src.Typo_Str {
			if len(obj.Value) == 0 {
				obj.Value = ""
			} else {
				value := src.UnquoteStr(obj.Value)

				action, ok := obj.options[src.Option_Action]
				if ok && len(action.Value) > 0 {
					str := strings.ToLower(src.UnquoteStr(action.Value))

					cmds := strings.SplitN(str, src.ActionSep, -1)

					bytes := []byte(value)
					acted := false

					for _, cmd := range cmds {
						cmd = src.UnquoteStr(cmd)

						if len(cmd) == 0 {
							continue
						}

						switch cmd {
						case src.Option_MD5:
							if len(bytes) > 0 {
								hash := md5.Sum(bytes)
								bytes = hash[:]
								acted = true
							}

						case src.Option_SHA256:
							if len(bytes) > 0 {
								hash := sha256.Sum256(bytes)
								bytes = hash[:]
								acted = true
							}

						case src.Option_SHA384:
							if len(bytes) > 0 {
								hash := sha512.Sum384(bytes)
								bytes = hash[:]
								acted = true
							}

						case src.Option_SHA512:
							if len(bytes) > 0 {
								hash := sha512.Sum512(bytes)
								bytes = hash[:]
								acted = true
							}

						case src.Option_Base64:
							if len(bytes) > 0 {
								value = base64.StdEncoding.EncodeToString(bytes)
								bytes = make([]byte, 0)
								acted = true
							}

						case src.Option_Hex:
							if len(bytes) > 0 {
								value = hex.EncodeToString(bytes)
								bytes = make([]byte, 0)
								acted = true
							}
						}
					}

					//default use base64
					if len(bytes) > 0 && acted {
						value = base64.StdEncoding.EncodeToString(bytes)
					}

					if obj.Value != value {
						if len(obj.CommentTriple) > 0 {
							obj.CommentTriple = obj.CommentTriple + ", value=" + obj.Value + ", action=" + action.Value
						} else {
							obj.CommentTriple = "// value=" + obj.Value + ", action=" + action.Value
						}
					}
				}

				obj.Value = strconv.Quote(value)
			}
		}
	}
}

// ////////////////////////////////////////////////////////////////////////////
type NameSpace struct {
	ConstObject `json:",inline" xml:",innerxml"`
	Enums       []*ConstObject `json:"enums,omitempty" xml:"enum"`
	Namespaces  []*NameSpace   `json:"namespaces,omitempty" xml:"namespace"`
}

func NewNameSpace(parent *ConstObject) *NameSpace {
	obj := &NameSpace{}
	obj.initialize(parent)
	return obj
}

func (obj *NameSpace) initialize(parent *ConstObject) {
	obj.ConstObject.initialize(parent, TypeNamespace)
	obj.Enums = make([]*ConstObject, 0)
	obj.Namespaces = make([]*NameSpace, 0)
}

func (obj *NameSpace) isEmpty() bool {
	return len(obj.Children) == 0 && len(obj.Enums) == 0 && len(obj.Namespaces) == 0
}

// check current scope identifier is redefined
func (obj *NameSpace) identRedefined(object *ConstObject, line int, column int) bool {
	if obj.ConstObject.identRedefined(object, line, column) {
		return true
	}

	for _, eumObj := range obj.Enums {
		if eumObj.Ident == object.Ident {
			fmt.Fprintf(os.Stderr, "Identifier '%s' is redefined at %d : %d.\n", object.Ident, line, column+1)
			return true
		}
	}

	for _, namespaceObj := range obj.Namespaces {
		if namespaceObj.Ident == object.Ident {
			fmt.Fprintf(os.Stderr, "Identifier '%s' is redefined at %d : %d.\n", object.Ident, line, column+1)
			return true
		}
	}

	return false
}

func (obj *NameSpace) addEnum(enum *ConstObject, line int, column int) {
	if obj.identRedefined(enum, line, column) {
		return
	}

	obj.Enums = append(obj.Enums, enum)
}

func (obj *NameSpace) addNamespace(namespace *NameSpace, line int, column int) {
	if obj.identRedefined(&namespace.ConstObject, line, column) {
		return
	}

	obj.Namespaces = append(obj.Namespaces, namespace)
}
