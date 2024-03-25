package hackgrp

import (
	"context"
	"errors"
	"github.com/draculaas/deploy/business/web/v1/response"
	"github.com/draculaas/deploy/core/web"
	"math/rand"
	"net/http"
)

func Hack(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100) % 2; n == 0 {
		return response.NewError(errors.New("TRUST ERROR"), http.StatusBadRequest)
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
