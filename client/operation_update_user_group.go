package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func (csClient *CSClient) UpdateUserGroup(ctx context.Context, req *CreateUserGroupRequest) (*Group, error) {
	method := "PUT"

	url := fmt.Sprintf("%s/user/groups", csClient.config.URL)
	log.Debug("Url-->", url)
	log.Debug("req-->", req)
	bodyAsBytes, err := marshalCreateUserGroupRequest(req)
	if err != nil {
		return nil, err
	}
	log.Debug("method-->", method)
	log.Debug("bodyAsBytes-->", bodyAsBytes)

	httpReq, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(bodyAsBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}
	log.Debug(" adding headers...")

	log.Warn("httpReq-->", httpReq)
	httpResp, err := csClient.signV2AndDo(req.AuthToken, httpReq, bodyAsBytes)
	log.Warn("httpResp-->", httpResp)
	if err != nil {
		return nil, fmt.Errorf("failed to %s to %s: %s", method, url, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			_ = fmt.Errorf("failed to Close response body  %s", err)
		}
	}(httpResp.Body)

	var readUserGroupResp []Group
	if err := csClient.unmarshalJSONBody(httpResp.Body, &readUserGroupResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response body sdjhskdhskdskdskdksdkskjd: %s", err)
	}

	log.Info("readUserGroupResp", readUserGroupResp[0].Id)

	return &readUserGroupResp[0], err

}

func marshalUpdateUserGroupRequest(req *CreateUserGroupRequest) ([]byte, error) {
	body := []interface{}{
		map[string]interface{}{
			"id":          req.Id,
			"name":        req.Name,
			"permissions": req.Permission,
		},
	}
	log.Info("marshalCreateUserGroupRequest -->", marshalCreateUserGroupRequest)

	bodyAsBytes, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	return bodyAsBytes, nil
}
