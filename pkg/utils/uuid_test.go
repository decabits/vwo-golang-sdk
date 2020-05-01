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

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFor(t *testing.T) {
	vwoInstance := getInstance()

	accountID := 12345
	userID := "Chris"
	actual := generateFor(vwoInstance, userID, accountID)
	expected := "810E5C0420C2541C8E9BDACD245E6476"
	assert.Equal(t, expected, actual, "Expected and Actual UUIDs should be same")

	accountID = 12345
	userID = "__123__"
	actual = generateFor(vwoInstance, userID, accountID)
	expected = "50A5B167FB6356A796F91D8951E480EE"
	assert.Equal(t, expected, actual, "Expected and Actual UUIDs should be same")

	accountID = 12345
	userID = "We@#dcs3232.f3"
	actual = generateFor(vwoInstance, userID, accountID)
	expected = "AAB4580A6BB3525FAA31DC341752D501"
	assert.Equal(t, expected, actual, "Expected and Actual UUIDs should be same")
}
