package mapper

import (
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/rajaabluu/commerce/backend/internal/entity"
)

func ToMidtransItem(item entity.OrderItem) midtrans.ItemDetails {
	id := strconv.Itoa(int(item.ID))
	return midtrans.ItemDetails{
		ID:    id,
		Name:  item.ProductName,
		Price: item.Price,
		Qty:   int32(item.Quantity),
	}
}

func ToSnapRequest(order *entity.Order, user *entity.User, address *entity.Address) *snap.Request {
	orderId := strconv.Itoa(int(order.ID))
	return &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: order.TotalPrice,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Phone: address.Phone,
			ShipAddr: &midtrans.CustomerAddress{
				FName:    address.RecipientName,
				Phone:    address.Phone,
				City:     address.City,
				Postcode: address.PostalCode,
				Address:  address.StreetAddress,
			},
		},
	}
}
