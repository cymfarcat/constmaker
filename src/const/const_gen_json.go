// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"encoding/json"
	"fmt"
	"os"

	src "github.com/cymfarcat/constmaker/src"
)

func (obj *ConstParser) GenJson(option *src.Options) {
	jsonData, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "generate json err: %s.\n", err)
	} else {
		path := option.GetOutputFile("json")
		data := string(jsonData)
		obj.writeFile(path, data, false, "", "")

	}
}
