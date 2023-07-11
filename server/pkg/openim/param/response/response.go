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

package response

// BaseResponse base response.
type BaseResponse struct {
	ErrCode uint   `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	ErrDlt  string `json:"errDlt"`
}

// UserTokenResponse user token response.
type UserTokenResponse struct {
	BaseResponse
	Data TokenData `json:"data"`
}

// TokenData token data.
type TokenData struct {
	Token             string `json:"token"`
	ExpireTimeSeconds uint   `json:"expireTimeSeconds"`
}