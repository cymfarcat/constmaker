// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

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
)

const (
	TypeObj       uint8 = 1
	TypeEnum      uint8 = 2
	TypeEnumValue uint8 = 3
	TypeNamespace uint8 = 4
	TypeProperty  uint8 = 5
)

// ////////////////////////////////////////////////////////////////////////////
type Object struct {
	prefix       string
	suffix       string
	Ident        string `json:"id" xml:"id,attr"`
	Typo         string `json:"type" xml:"type,attr"`
	Value        string `json:"value" xml:"value,attr"`
	value        string
	CommentDoc   string
	CommentTripe string    `json:"comment" xml:"comment,attr"`
	Children     []*Object `json:"consts,omitempty" xml:"const"`
	genObjs      []*Object

	parent  *Object
	objType uint8
	options map[string]*Object

	enumToPrefix      bool
	namespaceToPrefix bool
	freeze            bool
	macroDefine       bool
	upperIdent        bool
	lowerIdent        bool
	upperIdentCamel   bool
	lowerIdentCamel   bool

	enumSetValue bool
}

func NewObject(parent *Object, objType uint8) *Object {
	obj := &Object{}
	obj.initialize(parent, objType)
	return obj
}

func (obj *Object) cleanup() {}

func (obj *Object) initialize(parent *Object, objType uint8) {
	obj.Children = make([]*Object, 0)
	obj.genObjs = make([]*Object, 0)
	obj.options = make(map[string]*Object)
	obj.parent = parent
	obj.objType = objType
}

// check current scope identifier is redefined
func (obj *Object) identRedefined(object *Object, line int, column int) bool {
	for _, constObj := range obj.Children {
		if constObj.Ident == object.Ident {
			fmt.Fprintf(os.Stderr, "Identifier '%s' is redefined at %d : %d.\n", object.Ident, line, column+1)
			return true
		}
	}
	return false
}

func (obj *Object) addObject(object *Object, line int, column int) {
	if !obj.identRedefined(object, line, column) {
		obj.Children = append(obj.Children, object)
	}
}

func (obj *Object) addOption(option *Object) {
	obj.options[option.Ident] = option
}

func (obj *Object) sortChildren(sortAsc bool, sortDesc bool, sortAscLex bool, sortDescLex bool) {
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
			case isAlphaA:
				return a < b
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
			case isAlphaA:
				return a > b
			default:
				return a > b
			}
		})
	}
}

func (obj *Object) checkValue() {
	valueMap := make(map[string]string)

	for _, constObj := range obj.Children {

		ident, exists := valueMap[constObj.Value]
		if exists {
			fmt.Fprintf(os.Stderr, "Warning: identifier '%s' and '%s' has duplicate value: %s.\n", constObj.Ident, ident, constObj.Value)
		}

		valueMap[constObj.Value] = constObj.Ident
	}
}

func (obj *Object) inheritedProperty(parent *Object) {
	obj.prefix = parent.prefix
	obj.suffix = parent.suffix
	obj.freeze = parent.freeze
	obj.macroDefine = parent.macroDefine
	obj.upperIdent = parent.upperIdent
	obj.lowerIdent = parent.lowerIdent
	obj.upperIdentCamel = parent.upperIdentCamel
	obj.lowerIdentCamel = parent.lowerIdentCamel
}

/*
 * apply property only for children
 */
func (obj *Object) applyProperty() {
	if len(obj.Children) == 0 {
		return
	}

	// children use parent's prefix
	if len(obj.prefix) > 0 {
		for _, constObj := range obj.Children {
			constObj.prefix = obj.prefix
		}
	}

	// children use parent's suffix
	if len(obj.suffix) > 0 {
		for _, constObj := range obj.Children {
			constObj.suffix = obj.suffix
		}
	}

	// if obj freeze, children need freeze
	if obj.freeze {
		for _, constObj := range obj.Children {
			constObj.freeze = obj.freeze
		}
	}

	// if obj useDefine, children need useDefine
	if obj.macroDefine {
		for _, constObj := range obj.Children {
			constObj.macroDefine = obj.macroDefine
		}
	}

	// if obj upperIdent, children need upperIdent
	if obj.upperIdent {
		for _, constObj := range obj.Children {
			constObj.upperIdent = obj.upperIdent
		}
	}

	// if obj lowerIdent, children need lowerIdent
	if obj.lowerIdent {
		for _, constObj := range obj.Children {
			constObj.lowerIdent = obj.lowerIdent
		}
	}

	// if obj cameCaseIdent, children need cameCaseIdent
	if obj.upperIdentCamel {
		for _, constObj := range obj.Children {
			constObj.upperIdentCamel = obj.upperIdentCamel
		}
	}

	// if obj snakeCaseIdent, children need snakeCaseIdent
	if obj.lowerIdentCamel {
		for _, constObj := range obj.Children {
			constObj.lowerIdentCamel = obj.lowerIdentCamel
		}
	}
}

func (obj *Object) clearProperty() {
	obj.prefix = ""
	obj.suffix = ""
	// obj.freeze = false
	// obj.macroDefine = false
	// obj.upperIdent = false
	// obj.lowerIdent = false
	// obj.upperIdentCame = false
	// obj.lowerIdentCame = false
}

func (obj *Object) detectType(str string) {
	if len(str) == 0 {
		obj.Typo = Typo_Str
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
			obj.Typo = Typo_F32
			obj.Value = strconv.FormatFloat(floatVal, 'f', -1, 32)
		}
	} else if floatVal, err := strconv.ParseFloat(str, 64); err == nil { // then ParseFloat64
		if !strings.ContainsAny(str, "eE.") && floatVal == math.Trunc(floatVal) {
			obj.detectType(fmt.Sprintf("%.0f", floatVal))
		} else {
			obj.Typo = Typo_F64
			obj.Value = strconv.FormatFloat(floatVal, 'f', -1, 64)
		}
	} else {
		obj.Typo = Typo_Str
		obj.Value = str
	}
}

func (obj *Object) genValue(option *ConstOption) {
	// correct indent
	correctIdent(obj)

	// apply prefix option
	prefix, ok := obj.options[Option_Prefix]
	if ok && len(prefix.Value) > 0 {
		obj.prefix = unquoteStr(prefix.Value)
	}

	// apply suffix option
	suffix, ok := obj.options[Option_Suffix]
	if ok && len(suffix.Value) > 0 {
		obj.suffix = unquoteStr(suffix.Value)
	}

	// apply property option
	obj.enumToPrefix = false
	obj.namespaceToPrefix = false
	obj.freeze = false
	obj.macroDefine = false
	obj.upperIdent = false
	obj.lowerIdent = false
	obj.upperIdentCamel = false
	obj.lowerIdentCamel = false

	property, ok := obj.options[Option_Property]
	if ok && len(property.Value) > 0 {
		str := strings.ToLower(unquoteStr(property.Value))

		if strings.Contains(str, Option_EnumToPrefix) {
			obj.enumToPrefix = true
		}

		if strings.Contains(str, Option_NamespaceToPrefix) || strings.Contains(str, Option_NSToPrefix) {
			obj.namespaceToPrefix = true
		}

		if strings.Contains(str, Option_Freeze) {
			obj.freeze = true
		}

		if strings.Contains(str, Option_MacroDefine) {
			obj.macroDefine = true
		}

		if strings.Contains(str, Option_UpperIdent) {
			obj.upperIdent = true
		}

		if strings.Contains(str, Option_LowerIdent) {
			obj.lowerIdent = true
		}

		if strings.Contains(str, Option_UpperIdentCamel) {
			obj.upperIdentCamel = true
		}

		if strings.Contains(str, Option_LowerIdentCamel) {
			obj.lowerIdentCamel = true
		}
	}

	// apply action option
	bitFlag := false
	bitFlagHex := false
	genStr := false
	genId := false

	sortAsc := false
	sortDesc := false
	sortAscLex := false
	sortDescLex := false

	action, ok := obj.options[Option_Action]
	if ok && len(action.Value) > 0 {
		str := strings.ToLower(unquoteStr(action.Value))

		if strings.Contains(str, Option_BitFlag) {
			bitFlag = true
		}

		if strings.Contains(str, Option_BitFlagHex) {
			bitFlagHex = true
		}

		if strings.Contains(str, Option_GenStr) {
			genStr = true
		}

		if strings.Contains(str, Option_GenId) {
			genId = true
		}

		if strings.Contains(str, Option_SortAscLex) {
			sortAscLex = true
		} else if strings.Contains(str, Option_SortAsc) {
			sortAsc = true
		}

		if strings.Contains(str, Option_SortDescLex) {
			sortDescLex = true
		} else if strings.Contains(str, Option_SortDesc) {
			sortDesc = true
		}
	}

	// pre-check bitFlag||bitFlagHex
	if (bitFlag || bitFlagHex) && len(obj.Children) > int(BitFlagMax) {
		fmt.Fprintf(os.Stderr, "\"%s\" sub-items exceeds the maximum value of the type, can't exec bitflag.\n", obj.Ident)
		bitFlag = false
		bitFlagHex = false
	}

	// calc max bitflag
	bitFlagCount := bitFlagCount(adjustLen(len(obj.Children), option.bitFlagNoneAll))

	if obj.objType == TypeEnum {
		// do sort
		obj.sortChildren(sortAsc, sortDesc, sortAscLex, sortDescLex)

		// apply property to all children
		obj.applyProperty()

		// first check type
		typo := obj.Typo

		// check enum if set value
		obj.enumSetValue = false
		for _, constObj := range obj.Children {
			if len(constObj.Value) != 0 {
				obj.enumSetValue = true
				break
			}
		}

		// gen enum value
		if bitFlag || bitFlagHex {
			obj.enumSetValue = true

			if len(typo) == 0 {
				typo = bitFlagType(uint64(len(obj.Children)), option.bitFlagNoneAll)
			}

			allValue := 0
			for idx, constObj := range obj.Children {
				constObj.Typo = typo
				constObj.Value = "(1 << " + strconv.FormatInt(int64(idx), 10) + ")"
				if bitFlagHex && bitFlagCount != 0 {
					constObj.Value = fmt.Sprintf("0x%0*X", bitFlagCount, 1<<idx)
				}
				constObj.value = constObj.Value
				allValue |= 1 << idx
			}

			if option.bitFlagNoneAll {
				// add None object
				none := NewObject(obj, TypeEnumValue)
				none.inheritedProperty(obj)
				none.Ident = option.bitFlagNone
				none.Typo = typo
				none.Value = "0"
				if bitFlagHex && bitFlagCount != 0 {
					none.Value = fmt.Sprintf("0x%0*X", bitFlagCount, 0)
				}
				none.value = option.bitFlagNone
				obj.Children = append([]*Object{none}, obj.Children...)

				// add All object
				all := NewObject(obj, TypeEnumValue)
				all.inheritedProperty(obj)
				all.Ident = option.bitFlagAll
				all.Typo = typo
				all.Value = "0x" + strconv.FormatInt(int64(allValue), 16)
				if bitFlagHex && bitFlagCount != 0 {
					all.Value = fmt.Sprintf("0x%0*X", bitFlagCount, allValue)
				}
				all.value = option.bitFlagAll
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

		// gen str
		if genStr && len(obj.Children) > 0 {
			for _, constObj := range obj.Children {
				genObj := NewObject(obj, TypeObj)
				genObj.inheritedProperty(obj)
				genObj.Ident = constObj.Ident + UnderScore + ConstSTR
				if len(genObj.suffix) > 0 {
					genObj.Ident = constObj.Ident + UnderScore + genObj.suffix + UnderScore + ConstSTR
					genObj.suffix = ""
				}
				genObj.Typo = Typo_Str
				genObj.Value = strconv.Quote(option.genIdentNameObj("", constObj))

				obj.genObjs = append(obj.genObjs, genObj)
			}
		}

		// check duplicate value
		if option.checkValue {
			obj.checkValue()
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
			typo := bitFlagType(uint64(len(obj.Children)), option.bitFlagNoneAll)

			allValue := 0
			for idx, constObj := range obj.Children {
				genObj := NewObject(obj, TypeObj)
				genObj.inheritedProperty(obj)
				genObj.Ident = constObj.Ident + UnderScore + ConstID
				if len(genObj.suffix) > 0 {
					genObj.Ident = constObj.Ident + UnderScore + genObj.suffix + UnderScore + ConstID
					genObj.suffix = ""
				}
				genObj.Typo = typo

				if bitFlag || bitFlagHex {
					if option.bitFlagNoneAll && idx == 0 {
						// add None object
						none := NewObject(obj, TypeEnumValue)
						none.inheritedProperty(obj)
						none.Ident = option.bitFlagNone + UnderScore + ConstID
						if len(none.suffix) > 0 {
							none.Ident = option.bitFlagNone + UnderScore + none.suffix + UnderScore + ConstID
							none.suffix = ""
						}
						none.Typo = typo
						none.Value = "0"
						if bitFlagHex && bitFlagCount != 0 {
							none.Value = fmt.Sprintf("0x%0*X", bitFlagCount, 0)
						}
						obj.genObjs = append(obj.genObjs, none)
					}

					genObj.Value = "(1 << " + strconv.FormatInt(int64(idx), 10) + ")"
					if bitFlagHex && bitFlagCount != 0 {
						genObj.Value = fmt.Sprintf("0x%0*X", bitFlagCount, 1<<idx)
					}

					obj.genObjs = append(obj.genObjs, genObj)

					allValue |= 1 << idx

					if option.bitFlagNoneAll && idx == len(obj.Children)-1 {
						// add All object
						all := NewObject(obj, TypeEnumValue)
						all.inheritedProperty(obj)
						all.Ident = option.bitFlagAll + UnderScore + ConstID
						if len(all.suffix) > 0 {
							all.Ident = option.bitFlagAll + UnderScore + all.suffix + UnderScore + ConstID
							all.suffix = ""
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

		// check duplicate value
		if option.checkValue {
			obj.checkValue()
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
			obj.Typo = Typo_Str

			// if no value, use value as Value
			obj.Value = obj.value
		}

		if len(obj.Typo) == 0 {
			obj.Typo = Typo_Str
		}

		if obj.Typo == Typo_Str {
			if len(obj.Value) == 0 {
				obj.Value = ""
			} else {
				value := unquoteStr(obj.Value)

				action, ok := obj.options[Option_Action]
				if ok && len(action.Value) > 0 {
					str := strings.ToLower(unquoteStr(action.Value))

					cmds := strings.SplitN(str, ActionSep, -1)

					bytes := []byte(value)
					acted := false

					for _, cmd := range cmds {
						cmd = unquoteStr(cmd)

						if len(cmd) == 0 {
							continue
						}

						switch cmd {
						case Option_MD5:
							if len(bytes) > 0 {
								hash := md5.Sum(bytes)
								bytes = hash[:]
								acted = true
							}

						case Option_SHA256:
							if len(bytes) > 0 {
								hash := sha256.Sum256(bytes)
								bytes = hash[:]
								acted = true
							}

						case Option_SHA384:
							if len(bytes) > 0 {
								hash := sha512.Sum384(bytes)
								bytes = hash[:]
								acted = true
							}

						case Option_SHA512:
							if len(bytes) > 0 {
								hash := sha512.Sum512(bytes)
								bytes = hash[:]
								acted = true
							}

						case Option_Base64:
							if len(bytes) > 0 {
								value = base64.StdEncoding.EncodeToString(bytes)
								bytes = make([]byte, 0)
								acted = true
							}

						case Option_Hex:
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
						if len(obj.CommentTripe) > 0 {
							obj.CommentTripe = obj.CommentTripe + ", value=" + obj.Value + ", action=" + action.Value
						} else {
							obj.CommentTripe = "// value=" + obj.Value + ", action=" + action.Value
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
	Object     `json:",inline" xml:",innerxml"`
	Enums      []*Object    `json:"enums,omitempty" xml:"enum"`
	Namespaces []*NameSpace `json:"namespaces,omitempty" xml:"namespace"`
}

func NewNameSpace(parent *Object) *NameSpace {
	obj := &NameSpace{}
	obj.initialize(parent)
	return obj
}

func (obj *NameSpace) initialize(parent *Object) {
	obj.Object.initialize(parent, TypeNamespace)
	obj.Enums = make([]*Object, 0)
	obj.Namespaces = make([]*NameSpace, 0)
}

func (obj *NameSpace) isEmpty() bool {
	return len(obj.Children) == 0 && len(obj.Enums) == 0 && len(obj.Namespaces) == 0
}

// check current scope identifier is redefined
func (obj *NameSpace) identRedefined(object *Object, line int, column int) bool {
	if obj.Object.identRedefined(object, line, column) {
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

func (obj *NameSpace) addEnum(enum *Object, line int, column int) {
	if obj.identRedefined(enum, line, column) {
		return
	}

	obj.Enums = append(obj.Enums, enum)
}

func (obj *NameSpace) addNamespace(namespace *NameSpace, line int, column int) {
	if obj.identRedefined(&namespace.Object, line, column) {
		return
	}

	obj.Namespaces = append(obj.Namespaces, namespace)
}
