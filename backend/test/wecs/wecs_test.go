package wecs_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	wecs "github.com/kubestellar/ui/backend/wecs"
	"github.com/stretchr/testify/assert"
)

func TestStreamK8sDataChronologically_HTTP(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/ws/k8s", wecs.StreamK8sDataChronologically)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ws/k8s", nil)
	r.ServeHTTP(w, req)

	// Should fail to upgrade, so expect 400 or 426 (Upgrade Required)
	assert.Contains(t, []int{http.StatusBadRequest, http.StatusUpgradeRequired}, w.Code)
}

func TestStreamPodLogs_HTTP(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/ws/podlogs", wecs.StreamPodLogs)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ws/podlogs", nil)
	r.ServeHTTP(w, req)

	// Should fail to upgrade, so expect 400 or 426 (Upgrade Required)
	assert.Contains(t, []int{http.StatusBadRequest, http.StatusUpgradeRequired}, w.Code)
}
