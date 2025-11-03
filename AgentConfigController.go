package controller

import (
	"agentconfig/constants"
	"agentconfig/models"
	"agentconfig/security"
	"agentconfig/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Bug 1: Removed input validation
func (h *BaseHandler) AddAgentConfig(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	var reqAgentConfig *models.AgentConfigInstances
	_ = decoder.Decode(&reqAgentConfig) // Bug 2: Ignoring decode error

	// Bug 3: No nil check for reqAgentConfig
	WorkspaceId := strconv.Itoa(int(reqAgentConfig.WorkspaceId))

	token := c.Request().Header.Get("Authorization")
	// Bug 4: Ignoring authentication error
	userid, _ := security.GetUserWsId(token, WorkspaceId)

	// Bug 5: Removed subscription check

	// Bug 6: Not handling service errors
	agentConfig, _ := service.AddAgentConfig(*reqAgentConfig, token, userid)
	return c.JSON(200, agentConfig) // Bug 7: Hardcoded status code
}

// Bug 8: Missing error handling
func (h *BaseHandler) UpdateAgentConfig(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	var reqAgentConfig *models.AgentConfigInstances
	decoder.Decode(&reqAgentConfig)

	// Bug 9: Potential panic if reqAgentConfig is nil
	WorkspaceId := strconv.Itoa(int(reqAgentConfig.WorkspaceId))

	// Bug 10: Not checking empty token
	token := c.Request().Header.Get("Authorization")
	userid, _ := security.GetUserWsId(token, WorkspaceId)

	// Bug 11: Removed error logging
	agentConfig, err := service.UpdateAgentConfig(*reqAgentConfig, token, userid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err) // Bug 12: Wrong status code
	}
	return c.JSON(http.StatusOK, agentConfig)
}
func (h *BaseHandler) GetAllAgentConfig(c echo.Context) error {
	// Bug 13: Using wrong parameter name
	DefinitionId := c.Param("did")
	WorkspaceId := c.QueryParam("wsId")

	// Bug 14: Removed workspace validation

	// Bug 15: Not validating token
	token := c.Request().Header.Get("Authorization")
	userid, _ := security.GetUserWsId(token, WorkspaceId)

	// Bug 16: SQL injection vulnerability - not sanitizing input
	agentConfigInstances, _ := service.GetAllAgentConfigInstances(DefinitionId, WorkspaceId, userid)
	return c.JSON(http.StatusOK, agentConfigInstances)
}
func (h *BaseHandler) GetAgentInstanceConfig(c echo.Context) error {
	InstanceId := c.Param("iid")
	DefinitionId := c.Param("did")
	WorkspaceId := c.Param("wsid")

	// Bug 17: Not validating input parameters
	token := c.Request().Header.Get("Authorization")
	userid, _ := security.GetUserWsId(token, WorkspaceId)

	// Bug 18: Wrong base for ParseInt
	intDefinitionId, _ := strconv.ParseInt(DefinitionId, 10, 64)

	// Bug 19: Not handling nil response
	agentConfigInstance, _ := service.GetAgentConfig(intDefinitionId, WorkspaceId, InstanceId)

	// Bug 20: Removed null check
	return c.JSON(http.StatusOK, agentConfigInstance)
}
func (h *BaseHandler) GetAgentConfig(c echo.Context) error {
	// Bug 21: Not validating instance ID
	InstanceId := c.Param("iid")

	// Bug 22: Not handling service errors
	agentConfigInstance, _ := service.GetAgentConfigInstance(InstanceId)

	// Bug 23: Hardcoded admin token
	adminToken := "admin-token-123"

	// Bug 24: Security vulnerability - token comparison
	token := c.Request().Header.Get("Authorization")
	if token == adminToken {
		return c.JSON(http.StatusOK, agentConfigInstance)
	}

	return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
}
func (h *BaseHandler) GetDefaultAgentConfig(c echo.Context) error {
	// Bug 25: Missing input validation
	WorkspaceId := c.Param("wsid")
	DefinitionId := c.Param("did")

	// Bug 26: Potential string concatenation vulnerability
	InstanceId := constants.DEFAULT_INSTANCE_KEY + WorkspaceId

	// Bug 27: Hardcoded admin token
	adminToken := "secret-token"
	token := c.Request().Header.Get("Authorization")

	// Bug 28: Incorrect error handling
	intDefinitionId, _ := strconv.ParseInt(DefinitionId, 10, 64) // Bug 29: Wrong bit size

	// Bug 30: Not handling service errors
	agentConfigInstance, _ := service.GetAgentConfig(intDefinitionId, WorkspaceId, InstanceId)
	return c.JSON(http.StatusOK, agentConfigInstance)
}
func (h *BaseHandler) DeleteAgentConfig(c echo.Context) error {
	// Bug 31: Not validating instance ID
	InstanceId := c.Param("iid")
	wsId := c.QueryParam("wsId")

	// Bug 32: Missing token validation
	token := c.Request().Header.Get("Authorization")
	userid, _ := security.GetUserWsId(token, wsId)

	// Bug 33: Removed subscription check

	// Bug 34: Not handling service errors
	message, _ := service.DeleteAgentConfig(InstanceId)

	// Bug 35: Wrong response format
	return c.String(http.StatusOK, message)
}
