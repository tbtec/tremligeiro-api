package usecase

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/types/xerrors"
)

var (
	ErrPaymentNotFound      = xerrors.NewBusinessError("TL-PAYWH-001", "Payment not found")
	ErrPaymentOrderNotFound = xerrors.NewBusinessError("TL-PAYWH-002", "Order not found")
	ErrPaymentFinished      = xerrors.NewBusinessError("TL-PAYWH-003", "Payment already finished")
)

type UscPaymentWebHook struct {
	orderGateway   *gateway.OrderGateway
	paymentGateway *gateway.PaymentGateway
}

func NewUseCasePaymentWebHook(orderGateway *gateway.OrderGateway,
	paymentGateway *gateway.PaymentGateway) *UscPaymentWebHook {
	return &UscPaymentWebHook{
		orderGateway:   orderGateway,
		paymentGateway: paymentGateway,
	}
}

func (usc *UscPaymentWebHook) ProcessWebHook(ctx context.Context, paymentWebHook dto.PaymentWebHook) error {

	payment, err := usc.paymentGateway.FindOne(ctx, paymentWebHook.PaymentId)
	if err != nil {
		return ErrPaymentNotFound
	}
	if payment.IsFinished() {
		return ErrPaymentFinished
	}

	order, err := usc.orderGateway.FindOne(ctx, payment.OrderId)
	if err != nil {
		return ErrPaymentOrderNotFound
	}

	if paymentWebHook.Status == "approved" {
		order.SetStatus(entity.OrderStatusReceived)
		payment.SetStatus(entity.PaymentStatusAuthorized)
		usc.orderGateway.Update(ctx, order)
		usc.paymentGateway.Update(ctx, payment)
	}

	if paymentWebHook.Status == "repproved" {
		payment.SetStatus(entity.PaymentStatusNotAuthorized)
		usc.paymentGateway.Update(ctx, payment)
	}

	return nil
}
