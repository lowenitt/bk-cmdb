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
	"time"

	"configcenter/src/framework/api"
	"configcenter/src/framework/core/input"
)

func init() {

	api.RegisterFrequencyInputer(cloud, time.Minute*5)
}

var cloud = &hostInputer{}

type hostInputer struct {
}

// Init initialization method
func (cli *hostInputer) Init(ctx input.InputerContext) error {

	return nil
}

// Name the Inputer name.
// This information will be printed when the Inputer is abnormal, which is convenient for debugging.
func (cli *hostInputer) Name() string {
	return "host_inputer"
}

// Run the input should not be blocked
func (cli *hostInputer) Run(ctx input.InputerContext) *input.InputerResult {
	run()
	return nil

}

func (cli *hostInputer) Stop() error {
	return nil
}
