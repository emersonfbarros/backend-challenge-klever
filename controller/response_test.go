package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResSender_sendError(t *testing.T) {
	// mock context for test
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)

	rs := NewResSender()

	errMsg := "Error for test"
	rs.sendError(context, http.StatusNotFound, errMsg)

	// assertions
	assert.Equal(t, http.StatusNotFound, context.Writer.Status())
	assert.JSONEq(t, `{"message": "Error for test"}`, w.Body.String())
}

func TestResSender_sendSuccess(t *testing.T) {
	// mock context for test
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)

	rs := NewResSender()

	data := gin.H{"key": "value"}

	rs.sendSuccess(context, data)

	// assertions
	assert.Equal(t, http.StatusOK, context.Writer.Status())
	assert.JSONEq(t, `{"key": "value"}`, w.Body.String())
}
