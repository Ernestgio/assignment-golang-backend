package handler_test

import (
	"assignment-golang-backend/server"
	"assignment-golang-backend/testutils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleNotFound(t *testing.T) {
	t.Run("Should return 404 status code when request is not found", func(t *testing.T) {
		cfg := &server.RouterConfig{}

		req, _ := http.NewRequest("GET", "/not-found", nil)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}
