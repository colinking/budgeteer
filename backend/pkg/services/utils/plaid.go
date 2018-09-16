package utils

import (
	"strings"

	proto "github.com/colinking/budgeteer/backend/pkg/proto/plaid"
	"github.com/plaid/plaid-go/plaid"
)

// ToEnvironment converts a string to a Plaid environment.
func ToEnvironment(env string) plaid.Environment {
	switch env {
	case "development":
		return plaid.Development
	case "production":
		return plaid.Production
	default:
		return plaid.Sandbox
	}
}

// ToTransactionType converts a Plaid API transaction type to a proto transaction type.
func ToTransactionType(str string) proto.Transaction_Type {
	return proto.Transaction_Type(proto.Transaction_Type_value[strings.ToUpper(str)])
}

// ToCurrencyType converts a Plaid API currency type to a proto currency type.
func ToCurrencyType(str string) proto.Transaction_Currency {
	return proto.Transaction_Currency(proto.Transaction_Currency_value[strings.ToUpper(str)])
}

// ToTransactionLocation converts a Plaid API location to a proto location.
func ToTransactionLocation(loc plaid.Location) *proto.Transaction_Location {
	return &proto.Transaction_Location{
		Address:     loc.Address,
		City:        loc.City,
		Lat:         loc.Lat,
		Lon:         loc.Lon,
		State:       loc.State,
		StoreNumber: loc.StoreNumber,
		Zip:         loc.Zip,
	}
}

// ToTransactionPaymentMeta converts a Plaid API payment meta to a proto payment meta.
func ToTransactionPaymentMeta(meta plaid.PaymentMeta) *proto.Transaction_PaymentMeta {
	return &proto.Transaction_PaymentMeta{
		ByOrderOf:        meta.ByOrderOf,
		Payee:            meta.Payee,
		Payer:            meta.Payer,
		PaymentMethod:    meta.PaymentMethod,
		PaymentProcessor: meta.PaymentProcessor,
		Ppdid:            meta.PPDID,
		Reason:           meta.Reason,
		ReferenceNumber:  meta.ReferenceNumber,
	}
}

// ToTransaction converts a Plaid API transaction to a proto transaction.
func ToTransaction(tx plaid.Transaction) *proto.Transaction {
	return &proto.Transaction{
		Id:                   tx.ID,
		AccountId:            tx.AccountID,
		Category:             tx.Category,
		CategoryId:           tx.CategoryID,
		Type:                 ToTransactionType(tx.Type),
		MerchantName:         tx.Name,
		Amount:               tx.Amount,
		CurrencyType:         ToCurrencyType(tx.ISOCurrencyCode),
		Date:                 tx.Date,
		Location:             ToTransactionLocation(tx.Location),
		PaymentMeta:          ToTransactionPaymentMeta(tx.PaymentMeta),
		Pending:              tx.Pending,
		PendingTransactionId: tx.PendingTransactionID,
	}
}
