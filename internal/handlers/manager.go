package handlers

import (
	"github.com/IlyaZayats/managus/internal/requests"
	"github.com/IlyaZayats/managus/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ManagerHandlers struct {
	svc    *services.ManagerService
	engine *gin.Engine
}

func NewManagerHandlers(engine *gin.Engine, svc *services.ManagerService) (*ManagerHandlers, error) {
	h := &ManagerHandlers{
		svc:    svc,
		engine: engine,
	}
	h.initRoute()
	return h, nil
}

func (h *ManagerHandlers) initRoute() {
	h.engine.GET("/manager", h.GetManagers)
	h.engine.POST("/manager/delete", h.DeleteManager)
	h.engine.PUT("/manager", h.InsertManager)
	h.engine.POST("/manager", h.UpdateManager)
}

func (h *ManagerHandlers) GetManagers(c *gin.Context) {
	managers, err := h.svc.GetManagers()
	logrus.Debug(managers)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get managers error", "text": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "item": managers})
}

func (h *ManagerHandlers) DeleteManager(c *gin.Context) {

	req, ok := GetRequest[requests.DeleteManagerRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete manager request error", "text": ok})
		return
	}

	if err := h.svc.DeleteManager(req.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete manager error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *ManagerHandlers) InsertManager(c *gin.Context) {

	req, ok := GetRequest[requests.InsertManagerRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert manager request error", "text": ok})
		return
	}

	if err := h.svc.InsertManager(req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert manager error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *ManagerHandlers) UpdateManager(c *gin.Context) {

	req, ok := GetRequest[requests.UpdateManagerRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update manager request error", "text": ok})
		return
	}

	if err := h.svc.UpdateManagers(req.Id, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update manager error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
