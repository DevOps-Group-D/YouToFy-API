package authentication

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
	"github.com/DevOps-Group-D/YouToFy-API/utils"
)

const AUTHORIZE_ROUTE = "%s://%s:%s/authorize"

func Authorize(username string) bool {
	protocol := configs.Cfg.AuthenticationConfig.Protocol
	host := configs.Cfg.AuthenticationConfig.Host
	port := configs.Cfg.ApiConfig.Port
	authorizeRoute := fmt.Sprintf(AUTHORIZE_ROUTE, protocol, host, port)

	jsonData := fmt.Sprintf(`{
		"username:" %s
	}`, username)

	req, err := http.NewRequest(http.MethodPost, authorizeRoute, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Println("Error creating post authorize request:", err.Error())
		return false
	}

	res, err := utils.Client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err.Error())
		return false
	}
	defer res.Body.Close()

	return res.StatusCode == http.StatusOK
}
