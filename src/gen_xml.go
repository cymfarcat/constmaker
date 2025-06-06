// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"encoding/xml"
	"fmt"
	"os"
)

func (obj *ConstParser) genXml(option *ConstOption) {
	// change root name
	obj.XMLName = xml.Name{Local: upperCamelCase(option.fileName)}

	jsonData, err := xml.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "generate xml err: %s.\n", err)
	} else {
		path := option.getOutputFile("xml")
		data := string(jsonData)
		obj.writeFile(path, data, false, "", "")

	}
}
