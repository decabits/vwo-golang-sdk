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

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	vwoInstance, err := getInstance("./testdata/testdata.json")
	assert.Nil(t, err, "error fetching instance")
	userID := "Ash"

	tagKey := ""
	tagValue := ""
	pushed := vwoInstance.Push(tagKey, tagValue, userID)
	assert.False(t, pushed, "Invalid params")

	tagKey = "demoTagKey"
	tagValue = "demoTagVal"
	pushed = vwoInstance.Push(tagKey, tagValue, userID)
	assert.True(t, pushed, "Unable to Push")

	tagKey = "demoTagKey-Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis id tellus quis massa iaculis interdum. Morbi rutrum, lacus ac egestas lobortis, lectus lectus mollis sem, eget vehicula justo velit ut erat. Mauris ac ligula id nulla laoreet fringilla non at purus. Quisque eu risus quis mi convallis sagittis. Aliquam luctus posuere mollis. Nullam rhoncus mauris a lorem sagittis efficitur. Nulla quis risus sit amet tellus bibendum facilisis. Aliquam erat volutpat.In aliquam imperdiet nulla, sed consequat ex pharetra eget. Mauris eget vestibulum nunc. Morbi sem lectus, elementum sit amet laoreet at, euismod a purus. Aliquam ut tristique neque, tempor aliquet nisl. Aenean vestibulum lectus ut semper fringilla. Phasellus accumsan lorem at risus laoreet, non molestie est egestas. Fusce ac tellus vel nulla mollis auctor. Praesent ac laoreet lorem.Proin bibendum sodales nulla eget consectetur. Etiam auctor non lacus ac venenatis. Maecenas a magna dolor. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In id ornare nunc, vel sodales purus. Integer ultricies dui at tortor bibendum facilisis. Vestibulum mollis porttitor ligula. Fusce odio tortor, imperdiet vel lectus id, rhoncus facilisis tortor. Ut sagittis purus non sapien condimentum, vitae iaculis ligula pharetra. Donec in metus id libero pellentesque mattis sed sed metus. Maecenas a nisi ut risus volutpat posuere. Nunc id semper quam, ac vehicula lacus. Aliquam erat volutpat.Aliquam cursus lacinia odio non pretium. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam lectus ex, consectetur at augue pretium, iaculis cursus lacus. Aliquam nec porta erat. Aliquam blandit lobortis sapien, vitae maximus."
	tagValue = "demoTagVal"
	pushed = vwoInstance.Push(tagKey, tagValue, userID)
	assert.False(t, pushed, "Unable to Push")

	tagKey = "demoTagKey"
	tagValue = "demoTagVal-Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis id tellus quis massa iaculis interdum. Morbi rutrum, lacus ac egestas lobortis, lectus lectus mollis sem, eget vehicula justo velit ut erat. Mauris ac ligula id nulla laoreet fringilla non at purus. Quisque eu risus quis mi convallis sagittis. Aliquam luctus posuere mollis. Nullam rhoncus mauris a lorem sagittis efficitur. Nulla quis risus sit amet tellus bibendum facilisis. Aliquam erat volutpat.In aliquam imperdiet nulla, sed consequat ex pharetra eget. Mauris eget vestibulum nunc. Morbi sem lectus, elementum sit amet laoreet at, euismod a purus. Aliquam ut tristique neque, tempor aliquet nisl. Aenean vestibulum lectus ut semper fringilla. Phasellus accumsan lorem at risus laoreet, non molestie est egestas. Fusce ac tellus vel nulla mollis auctor. Praesent ac laoreet lorem.Proin bibendum sodales nulla eget consectetur. Etiam auctor non lacus ac venenatis. Maecenas a magna dolor. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In id ornare nunc, vel sodales purus. Integer ultricies dui at tortor bibendum facilisis. Vestibulum mollis porttitor ligula. Fusce odio tortor, imperdiet vel lectus id, rhoncus facilisis tortor. Ut sagittis purus non sapien condimentum, vitae iaculis ligula pharetra. Donec in metus id libero pellentesque mattis sed sed metus. Maecenas a nisi ut risus volutpat posuere. Nunc id semper quam, ac vehicula lacus. Aliquam erat volutpat.Aliquam cursus lacinia odio non pretium. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam lectus ex, consectetur at augue pretium, iaculis cursus lacus. Aliquam nec porta erat. Aliquam blandit lobortis sapien, vitae maximus."
	pushed = vwoInstance.Push(tagKey, tagValue, userID)
	assert.False(t, pushed, "Unable to Push")
}