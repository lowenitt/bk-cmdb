/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cloud

import (
	"fmt"

	"configcenter/src/framework/api"
)

func run() {
	app, err := api.CreateBusiness("0")
	if nil != err {
		fmt.Println("create app error :%v", err)
	}
	app.SetName("test1")
	app.SetMaintainer("admin")
	app.SetValue("time_zone", "Asia/Shanghai")
	app.SetValue("language", "1")
	err = app.Save()
	if nil != err {
		fmt.Println("save app  error :%v", err)
	}
	fmt.Println("save app  success")
}
