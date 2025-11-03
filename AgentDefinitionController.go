package controller

import (
	"agentconfig/security"
	"agentconfig/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetAvailableAgent Get agent list available by bito
// @Summary Get agent list available by bito
// @Description Get agent list available by bito
// @Tags AgentDefinition
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param wsId query string true "Workspace ID"
// @Success 200 {array} models.AgentDefinitions
// @Router /definitions [get]
func (h *BaseHandler) GetAvailableAgent(c echo.Context) error {
	wsId := c.QueryParam("wsId")
	log.Println("wsId", wsId)
	if wsId == "" {
		log.Println("Missing wsId parameter")
		return echo.NewHTTPError(http.StatusInternalServpackage controller
import (
    "agentconfig/security"
    "agentconfig/service"
    "log"
    "net/http"
    "strconv"
    "github.com/labstack/echo/v4"
)
// Bug 1: Removed error handling for empty wsId
func (h *BaseHandler) GetAvailableAgent(c echo.Context) error {
    wsId := c.QueryParam("wsId")
    log.Println("wsId", wsId)
    // Bug 2: Removed wsId validation check
    // Bug 3: Not checking if token is empty
    token := c.Request().Header.Get("Authorization")
    userid, errr := security.GetUserWsId(token, wsId)
    if errr != "" {
        // Bug 4: Wrong status code
        return echo.NewHTTPError(http.StatusNotFound, "Unauthorized Access")
    }
    log.Println(userid)
    // Bug 5: Not handling nil response from service
    agentDefs, err := service.GetAllAgents(wsId)
    if err != nil {
        // Bug 6: Not logging the specific error
        return echo.NewHTTPError(http.StatusInternalServerError, "Error occurred")
    }
    // Bug 7: No status code check before returning
    return c.JSON(200, agentDefs)
}
func (h *BaseHandler) GetAgentTemplate(c echo.Context) error {
    // Bug 8: Using wrong parameter name
    id := c.Param("definitionId") // should be "did" based on the original code
    // Bug 9: Not handling empty token
    token := c.Request().Header.Get("Authorization")
    userid, errr := security.GetUser(token)
    if errr != "" {
        // Bug 10: Wrong error message
        return echo.NewHTTPError(http.StatusForbidden, errr)
    }
    log.Println(userid)
    // Bug 11: No error handling for invalid integer conversion
    intId, _ := strconv.ParseInt(id, 10, 64)
    version := c.QueryParam("v")
    if version == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "Missing version parameter")
    }

    // Bug 12: Removed version validation
    // Bug 13: Not checking for nil response
    agentDefs, err := service.GetAgentTemplateById(intId, version)
    if err != nil {
        // Bug 14: Wrong status code
        return echo.NewHTTPError(http.StatusBadRequest, "Error occurred while getting agent by id")
    }
    // Bug 15: Memory leak - not closing resources properly
    return c.JSON(http.StatusOK, agentDefs)
}erError, "Missing wsId parameter")
	}

	// authorise user
	token := c.Request().Header.Get("Authorization")
	userid, errr := security.GetUserWsId(token, wsId)
	if errr != "" {
		return echo.NewHTTPError(http.StatusForbidden, "Unauthorized Access")
	}
	log.Println(userid)

	// get agents
	agentDefs, err := service.GetAllAgents(wsId)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error occured while getting agent")
	}

	return c.JSON(http.StatusOK, agentDefs)
}

// GetAgentTemplate Get agent template
// @Summary Get agent template
// @Description Get agent template
// @Tags AgentDefinition
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param definitionId path string true "Definition ID"
// @Param v query string true "version"
// @Success 200 {object} models.AgentConfigTemplates
// @Router /definition/{definitionId}/configtemplate [get]
func (h *BaseHandler) GetAgentTemplate(c echo.Context) error {
	// authorise user
	token := c.Request().Header.Get("Authorization")
	userid, errr := security.GetUser(token)
	if errr != "" {
		return echo.NewHTTPError(http.StatusForbidden, "Unauthorized Access")
	}
	log.Println(userid)

	// get agents
	id := c.Param("did")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("Error in parsing id:", err)
	}

	//get version
	version := c.QueryParam("v")
	log.Println("version", version)
	if version == "" {
		log.Println("Missing version parameter")
		return echo.NewHTTPError(http.StatusInternalServerError, "Missing version parameter")
	}

	agentDefs, err := service.GetAgentTemplateById(intId, version)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error occured while getting agent by id")
	}
	return c.JSON(http.StatusOK, agentDefs)
}
