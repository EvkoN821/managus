package handlers

import (
	"github.com/IlyaZayats/managus/internal/requests"
	"github.com/IlyaZayats/managus/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ClientHandlers struct {
	svc    *services.ClientService
	engine *gin.Engine
}

func NewClientHandlers(engine *gin.Engine, svc *services.ClientService) (*ClientHandlers, error) {
	h := &ClientHandlers{
		svc:    svc,
		engine: engine,
	}
	h.initRoute()
	return h, nil
}

func (h *ClientHandlers) initRoute() {
	h.engine.GET("/client", h.GetClients)           //
	h.engine.POST("/client/delete", h.DeleteClient) //
	h.engine.PUT("/client", h.InsertClient)         //
	h.engine.POST("/client", h.UpdateClient)        //
}

func (h *ClientHandlers) GetClients(c *gin.Context) {

	clients, err := h.svc.GetClients()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get clients error", "text": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "item": clients})
}

func (h *ClientHandlers) DeleteClient(c *gin.Context) {

	req, ok := GetRequest[requests.DeleteClientRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete client request error", "text": ok})
		return
	}

	if err := h.svc.DeleteClient(req.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete client error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *ClientHandlers) InsertClient(c *gin.Context) {

	req, ok := GetRequest[requests.InsertClientRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert client request error", "text": ok})
		return
	}

	if err := h.svc.InsertClient(req.ManagerId, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert client error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *ClientHandlers) UpdateClient(c *gin.Context) {

	req, ok := GetRequest[requests.UpdateClientRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update client request error", "text": ok})
		return
	}

	if err := h.svc.UpdateClient(req.Id, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update client error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
