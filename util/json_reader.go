// Copyright (c) 2018, Google, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package util

import (
	"encoding/json"
	"errors"
	"os"
)

func ParseJsonFromFileOrStdin(filePath string) (map[string]interface{}, error) {
	var fromFile *os.File
	var err error
	var jsonContent map[string]interface{}

	if filePath != "" {
		fromFile, err = os.Open(filePath)
		if err != nil {
			return nil, err
		}
	} else {
		fromFile = os.Stdin
	}

	fi, err := fromFile.Stat()
	if err != nil {
		return nil, err
	}

	pipedStdin := (fi.Mode() & os.ModeCharDevice) == 0
	if fi.Size() <= 0 && !pipedStdin {
		return nil, errors.New("No json input to parse.")
	}

	err = json.NewDecoder(fromFile).Decode(&jsonContent)
	if err != nil {
		return nil, err
	}
	return jsonContent, nil
}
