/*
   Copyright 2019-2020 Wingify Software Pvt. Ltd.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package core

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/google/logger"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	DSL                         map[string]interface{} `json:"dsl"`
	Expected                    bool                   `json:"expectation"`
	CustomVariable              map[string]interface{} `json:"custom_variables"`
	VariationTargetingVariables map[string]interface{} `json:"variation_targeting_variables"`
}

func TestSegmentEvaluator(t *testing.T) {
	var testdata map[string]map[string]TestCase
	data, err := ioutil.ReadFile("./testdata/testSegment.json")
	if err != nil {
		logger.Info("Error: " + err.Error())
	}

	if err = json.Unmarshal(data, &testdata); err != nil {
		logger.Info("Error: " + err.Error())
	}

	for parent, v := range testdata {
		for child, value := range v {
			var actual bool
			if value.CustomVariable != nil {
				actual = SegmentEvaluator(value.DSL, value.CustomVariable)
			} else {
				actual = SegmentEvaluator(value.DSL, value.VariationTargetingVariables)
			}
			expected := value.Expected
			assert.Equal(t, expected, actual, parent+" "+child)
		}
	}
}
