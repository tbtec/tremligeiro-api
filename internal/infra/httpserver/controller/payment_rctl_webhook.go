package controller

import (
	"context"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type PaymentWebHookRestController struct {
	controller *ctl.PaymentWebHookController
}

func NewPaymentWebHookRestController(container *container.Container) httpserver.IController {
	return &PaymentWebHookRestController{
		controller: ctl.NewPaymentWebHookController(container),
	}
}

func (ctl *PaymentWebHookRestController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {
	input := dto.PaymentWebHook{}

	err := request.ParseBody(ctx, &input)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	err2 := ctl.controller.Execute(ctx, input)
	if err2 != nil {
		return httpserver.HandleError(ctx, err2)
	}

	return httpserver.NoContent()
}
