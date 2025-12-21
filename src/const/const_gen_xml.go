// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"encoding/xml"
	"fmt"
	"os"

	src "github.com/cymfarcat/constmaker/src"
)

func (obj *ConstParser) GenXml(option *src.Options) {
	// change root name
	obj.XMLName = xml.Name{Local: src.UpperCamelCase(option.FileName, true)}

	jsonData, err := xml.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "generate xml err: %s.\n", err)
	} else {
		path := option.GetOutputFile("xml")
		data := string(jsonData)
		obj.writeFile(path, data, false, "", "")

	}
}
