package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/yourusername/url-shortener/internal/service"
)

type Handler struct {
	generator service.Generator
	redirect  service.Redirect
	analytics service.Analytics
}

func NewHandler(g service.Generator, r service.Redirect, a service.Analytics) *Handler {
	return &Handler{generator: g, redirect: r, analytics: a}
}

type shortenReq struct {
	URL string `json:"url" binding:"required,url"`
}

type shortenResp struct {
	Code     string `json:"code"`
	ShortURL string `json:"short_url"`
}

func (h *Handler) CreateShortURL(c *gin.Context) {
	var req shortenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code, shortURL, err := h.generator.GenerateCode(req.URL, c.Request.Host)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate code"})
		return
	}

	c.JSON(http.StatusCreated, shortenResp{Code: code, ShortURL: shortURL})
}

func (h *Handler) Redirect(c *gin.Context) {
	code := c.Param("code")
	referer := c.GetHeader("Referer")
	ua := c.GetHeader("User-Agent")
	ip := c.ClientIP()
	target, err := h.redirect.UpdateClickEvent(code, referer, ua, ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Redirect(http.StatusFound, target)
}

func (h *Handler) GetAnalytics(c *gin.Context) {
	code := c.Param("code")
	a, err := h.analytics.GetAnalytics(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, a)
}
