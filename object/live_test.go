// Copyright 2022 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package object

import (
	"fmt"
	"testing"
	"time"
)

func TestGetLive(t *testing.T) {
	InitConfig()

	for {
		room := getRoom("admin", "seminar_1")
		fmt.Printf("%v, %v\n", getLiveStreamOnlineMap(room), getLiveDomainOnlineCount(room))

		time.Sleep(time.Second)
	}
}
