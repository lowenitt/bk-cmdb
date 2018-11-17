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

package service

import (
	"context"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/gin-gonic/gin/json"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	types "configcenter/src/common/mapstr"
	meta "configcenter/src/common/metadata"
	"configcenter/src/common/util"
	"configcenter/src/scene_server/validator"
)

func (ps *ProcServer) CreateSwitch(req *restful.Request, resp *restful.Response) {
	language := util.GetLanguage(req.Request.Header)
	defErr := ps.CCErr.CreateDefaultCCErrorIf(language)
	pHeader := req.Request.Header
	ownerID := common.BKDefaultOwnerID
	objID := req.PathParameter(common.BKObjIDField)

	input := types.MapStr{}
	if err := json.NewDecoder(req.Request.Body).Decode(&input); err != nil {
		blog.Errorf("create object failed! decode request body err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommJSONUnmarshalFailed)})
		return
	}
	valid := validator.NewValidMap(ownerID, objID, pHeader, ps.Engine)
	if err := valid.ValidMap(input, common.ValidCreate, 0); err != nil {
		blog.Errorf("fail to valid input parameters. err:%s", err.Error())
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommFieldNotValid)})
		return
	}

	ret, err := ps.CoreAPI.ObjectController().Instance().CreateObject(context.Background(), objID, pHeader, input)
	if nil != err || !ret.Result {
		blog.Errorf("create object failed by  input :%v, return:%v, err: %v", input, ret, err)
		resp.WriteError(http.StatusInternalServerError, &meta.RespError{Msg: defErr.Error(common.CCErrProcCreateTemplateFail)})
		return
	}

	resp.WriteEntity(meta.NewSuccessResp(nil))
}
