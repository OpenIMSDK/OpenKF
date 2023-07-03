// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hooks

import (
	"fmt"

	"github.com/gin-gonic/gin"

	urltrie "github.com/OpenIMSDK/OpenKF/server/internal/middleware/hooks/url_trie"
)

func init() {
	urltrie.RegisterHook(GlobalHook{})
	fmt.Println("RegisterHook", "Register Hook[GlobalHook] success...")
}

// GlobalHook implement urltrie.Hook.
type GlobalHook struct {
	urltrie.Hook
}

// Pattern return pattern.
func (h GlobalHook) Pattern() string {
	return "/*"
}

func (h GlobalHook) Priority() int64 {
	return 0
}

// BeforeRun do something before controller run.
func (h GlobalHook) BeforeRun(c *gin.Context) {
	c.Next()
}

// AfterRun do something after controller run.
func (h GlobalHook) AfterRun(c *gin.Context) {
	c.Next()
}
