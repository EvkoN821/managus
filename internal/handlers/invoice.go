package handlers

import (
	"github.com/IlyaZayats/managus/internal/requests"
	"github.com/IlyaZayats/managus/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type InvoiceHandlers struct {
	svc    *services.InvoiceService
	engine *gin.Engine
}

func NewInvoiceHandlers(engine *gin.Engine, svc *services.InvoiceService) (*InvoiceHandlers, error) {
	h := &InvoiceHandlers{
		svc:    svc,
		engine: engine,
	}
	h.initRoute()
	return h, nil
}

func (h *InvoiceHandlers) initRoute() {
	h.engine.GET("/invoice", h.GetInvoices)           //
	h.engine.POST("/invoice/delete", h.DeleteInvoice) //
	h.engine.PUT("/invoice", h.InsertInvoice)         //
	h.engine.POST("/invoice", h.UpdateInvoice)        //
}

func (h *InvoiceHandlers) GetInvoices(c *gin.Context) {
	invoices, err := h.svc.GetInvoices()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get invoices error", "text": err.Error()})
		return
	}

	logrus.Debug(invoices)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "item": invoices})
}

func (h *InvoiceHandlers) DeleteInvoice(c *gin.Context) {

	req, ok := GetRequest[requests.DeleteInvoiceRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete invoice request error", "text": ok})
		return
	}

	if err := h.svc.DeleteInvoice(req.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete invoice error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *InvoiceHandlers) InsertInvoice(c *gin.Context) {

	req, ok := GetRequest[requests.InsertInvoiceRequest](c)
	logrus.Debug(req)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert invoice request error", "text": ok})
		return
	}

	if err := h.svc.InsertInvoice(req.InvoiceId, req.ContDate, req.ExecDate, req.Handed, req.Accepted, req.AddInfo, req.BasisDoc, req.ClientId, req.SumTotal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert invoice error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *InvoiceHandlers) UpdateInvoice(c *gin.Context) {

	req, ok := GetRequest[requests.UpdateInvoiceRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update invoice request error", "text": ok})
		return
	}

	if err := h.svc.UpdateInvoice(req.InvoiceId, req.ContDate, req.ExecDate, req.Handed, req.Accepted, req.AddInfo, req.BasisDoc, req.Id, req.SumTotal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update invoice error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
