package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Error(c, http.StatusBadRequest, gin.H{"error": "Bad request"})

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"Bad request"}`, w.Body.String())
}

func TestSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Success(c, http.StatusOK, gin.H{"message": "Success"})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"Success"}`, w.Body.String())
}

func TestBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	BadRequest(c, nil)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"Bad request"}`, w.Body.String())
}

func TestNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	NotFound(c, nil)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, `{"error":"Not found"}`, w.Body.String())
}

func TestInternalServerError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	InternalServerError(c, nil)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, `{"error":"Internal server error"}`, w.Body.String())
}
