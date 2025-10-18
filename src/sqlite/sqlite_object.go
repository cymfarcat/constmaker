// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package sqlite

import (
	"fmt"
	"os"

	src "github.com/cymfarcat/constmaker/src"
)

const (
	SQLITE_ARG string = "--" + src.SQLITE
)

const (
	TypeTable        uint8 = 1
	TypeIndex        uint8 = 2
	TypeView         uint8 = 3
	TypeTrigger      uint8 = 4
	TypeVirtualTable uint8 = 5
	TypeField        uint8 = 6
)

type SQLiteObject struct {
	name    map[string]string
	objType uint8
	fields  map[string]string
}

func NewSQLiteObject(name string, objType uint8) *SQLiteObject {
	obj := &SQLiteObject{}
	obj.objType = objType

	obj.name = make(map[string]string)

	name = src.UnwrapIdent(name)
	obj.name[name] = name

	obj.fields = make(map[string]string)
	return obj
}

func (obj *SQLiteObject) addField(option *src.Options, field string /*, inputFile string, create_action string, line int, column int*/) {
	field = src.UnwrapIdent(field)

	if option.IsFieldWhiteList(field) {
		fmt.Fprintf(os.Stdout, "field: %s in while list.", field)
		return
	}

	_, ok := obj.fields[field]
	if ok {
		//fmt.Fprintf(os.Stderr, "file: %s, %s duplicate name: %s, at %d : %d.\n", inputFile, field, create_action, line, column)
		return
	}

	obj.fields[field] = field
}

func (obj *SQLiteObject) applyOption(option *src.Options, constsMap map[string]string) {
	for name := range obj.name {
		constKey := getSQLiteName(option, obj.objType, "", name)
		obj.name[name] = constKey
		constsMap[constKey] = option.GenerateRandomString()

		for field := range obj.fields {
			constKey := getSQLiteName(option, TypeField, name, field)
			obj.fields[field] = constKey
			constsMap[constKey] = option.GenerateRandomString()
		}
	}
}

func (obj *SQLiteObject) altName(name string, constsMap map[string]string) (string, bool) {
	var value string
	var ok bool

	// first check name
	var name2 = src.UnwrapIdent(name)
	value, ok = obj.name[name2]
	if !ok {
		// then check fields
		value, ok = obj.fields[name2]
	}

	// then get alt value
	if ok {
		value, ok = constsMap[value]
		if ok {
			return value, true
		}
	}

	return name, false
}
