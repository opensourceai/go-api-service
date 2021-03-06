/*
 *    Copyright 2020 opensourceai
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package app

import (
	"github.com/jinzhu/gorm"
	"github.com/opensourceai/go-api-service/pkg/e"
	"net/http"
)

// 异常恢复
func Recover(appG *Gin) {
	if err := recover(); err == gorm.ErrRecordNotFound.Error() {
		appG.Response(http.StatusNotFound, e.INVALID_PARAMS, nil)
	} else if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR, nil)
	}

}
