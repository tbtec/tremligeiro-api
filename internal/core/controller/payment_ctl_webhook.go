package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type PaymentWebHookController struct {
	usc *usecase.UscPaymentWebHook
}

func NewPaymentWebHookController(container *container.Container) *PaymentWebHookController {
	return &PaymentWebHookController{
		usc: usecase.NewUseCasePaymentWebHook(
			gateway.NewOrderGateway(container.OrderRepository),
			gateway.NewPaymentGateway(container.PaymentService, container.PaymentRepository)),
	}
}

func (ctl *PaymentWebHookController) Execute(ctx context.Context, input dto.PaymentWebHook) error {
	return ctl.usc.ProcessWebHook(ctx, input)
}
