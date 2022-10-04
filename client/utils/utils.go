package utils

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SubmitRequestError(requestType string, url string, err error) error {
	return fmt.Errorf("Failed to submit %s request to %s => Error: %s", requestType, url, err)
}

func CreateRequestError(err error) error {
	return fmt.Errorf("Failed to create request => Error %s", err)
}

func CloseResponseError(err error) error {
	return fmt.Errorf("Failed to close response body => Error: %s", err)
}

func ResponseCodeError(res *http.Response, req *http.Request, bodyAsBytes []byte, respAsBytes []byte) error {
	return fmt.Errorf(`Client returned with a non 2xx status code => 
			Code: %d
			Method: %s
			URL: %s
			Request body: %s
			Response body: %s
		`,
		res.StatusCode, req.Method, req.URL, bodyAsBytes, respAsBytes,
	)
}

func UnmarshalXmlError(err error) error {
	return fmt.Errorf("Failed to unmarshal XML => Error: %s", err)
}

func UnmarshalJsonError(err error) error {
	return fmt.Errorf("Failed to unmarshal JSON => Error: %s", err)
}

func ReadResponseError(err error) error {
	return fmt.Errorf("Failed to read response body => Error: %s", err)
}

func MarshalJsonError(err error) error {
	return fmt.Errorf("Failed to marshal JSON => Error: %s", err)
}

func ConfigurationError(value string) diag.Diagnostics {
	return diag.Errorf("Failed to configure provider => Expected '%s' to be defined", value)
}

func NormalizingJsonError(err error) error {
	return fmt.Errorf("Failed to normalize JSON structure => Error: %s", err)
}

func CreateObjectGroupError(msg string) diag.Diagnostics {
	return diag.Errorf("Failure Creating Object Group => %s", msg)
}

func ValidateAuthType(keyAuthEnabled bool) error {
	if keyAuthEnabled {
		return fmt.Errorf("Failed to authenticate => API Keys used with incompatible resource")
	}

	return nil
}

func ContainsString(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false
}

func ConvertSetToMap(item interface{}) map[string]interface{} {
	return item.(*schema.Set).List()[0].(map[string]interface{})
}
