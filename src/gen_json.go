// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"encoding/json"
	"fmt"
	"os"
)

func (obj *ConstParser) genJson(option *ConstOption) {
	jsonData, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "generate json err: %s.\n", err)
	} else {
		path := option.getOutputFile("json")
		data := string(jsonData)
		obj.writeFile(path, data, false, "", "")

	}
}
