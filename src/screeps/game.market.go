package screeps

import "syscall/js"

type OrderType string

const (
	OrderSell OrderType = "sell"
	OrderBuy  OrderType = "buy"
)

type Transaction struct {
	TransactionId string
	Time          int
	Sender        *struct{ username string }
	Recipent      *struct{ username string }
	ResourceType  ResourceConstant
	Amount        int
	From          string
	To            string
	Description   string
	Order         *struct {
		Id    string
		Type  OrderType
		Price float64
	}
}

type Order struct {
	Id               string
	Created          int
	CreatedTimestamp *int
	Active           *bool
	Type             string
	ResourceType     ResourceConstant //TODO - MarketResourceConstant
	RoomName         *string
	Amount           int
	RemainingAmount  int
	TotalAmount      *int
	Price            float64
}

type Market struct {
	ref js.Value
}

func (market Market) Credits() int {
	return market.ref.Get("credits").Int()
}

func (market Market) IncomingTransactions() []Transaction {
	jsIncomingTransactions := market.ref.Get("incomingTransactions")
	transactionCount := jsIncomingTransactions.Length()
	result := make([]Transaction, transactionCount)
	for i := 0; i < transactionCount; i++ {
		jsTransaction := jsIncomingTransactions.Index(i)
		result[i] = Transaction{
			TransactionId: jsTransaction.Get("transactionid").String(),
			Time:          jsTransaction.Get("time").Int(),
			Sender:        nil,
			Recipent:      nil,
			ResourceType:  ResourceConstant(jsTransaction.Get("resourceType").String()),
			Amount:        jsTransaction.Get("amount").Int(),
			From:          jsTransaction.Get("from").String(),
			To:            jsTransaction.Get("to").String(),
			Description:   jsTransaction.Get("description").String(),
			Order:         nil,
		}

		// sender
		sender := jsTransaction.Get("sender")
		if !sender.IsUndefined() {
			result[i].Sender = &struct{ username string }{sender.Get("username").String()}
		}

		// recipent
		recipent := jsTransaction.Get("recipent")
		if !sender.IsUndefined() {
			result[i].Recipent = &struct{ username string }{recipent.Get("username").String()}
		}

		// order
		order := jsTransaction.Get("order")
		if !order.IsUndefined() {
			result[i].Order = &struct {
				Id    string
				Type  OrderType
				Price float64
			}{
				Id:    order.Get("id").String(),
				Type:  OrderType(order.Get("type").String()),
				Price: order.Get("price").Float(),
			}
		}
	}
	return result
}

func (market Market) OutgoingTransactions() []Transaction {
	jsOutgoingTransactions := market.ref.Get("outgoingTransactions")
	transactionCount := jsOutgoingTransactions.Length()
	result := make([]Transaction, transactionCount)
	for i := 0; i < transactionCount; i++ {
		jsTransaction := jsOutgoingTransactions.Index(i)
		result[i] = Transaction{
			TransactionId: jsTransaction.Get("transactionid").String(),
			Time:          jsTransaction.Get("time").Int(),
			Sender:        nil,
			Recipent:      nil,
			ResourceType:  ResourceConstant(jsTransaction.Get("resourceType").String()),
			Amount:        jsTransaction.Get("amount").Int(),
			From:          jsTransaction.Get("from").String(),
			To:            jsTransaction.Get("to").String(),
			Description:   jsTransaction.Get("description").String(),
			Order:         nil,
		}

		// sender
		sender := jsTransaction.Get("sender")
		if !sender.IsUndefined() {
			result[i].Sender = &struct{ username string }{sender.Get("username").String()}
		}

		// recipent
		recipent := jsTransaction.Get("recipent")
		if !sender.IsUndefined() {
			result[i].Recipent = &struct{ username string }{recipent.Get("username").String()}
		}

		// order
		order := jsTransaction.Get("order")
		if !order.IsUndefined() {
			result[i].Order = &struct {
				Id    string
				Type  OrderType
				Price float64
			}{
				Id:    order.Get("id").String(),
				Type:  OrderType(order.Get("type").String()),
				Price: order.Get("price").Float(),
			}
		}
	}
	return result
}

func (market Market) Orders() map[string]Order {
	jsOrders := market.ref.Get("orders")
	result := map[string]Order{}

	entries := object.Call("entries", jsOrders)
	length := entries.Length()

	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		jsOrder := entry.Index(1)
		order := Order{
			Id:               jsOrder.Get("id").String(),
			Created:          jsOrder.Get("created").Int(),
			CreatedTimestamp: nil,
			Active:           nil,
			Type:             jsOrder.Get("type").String(),
			ResourceType:     ResourceConstant(jsOrder.Get("resourceType").String()),
			RoomName:         nil,
			Amount:           jsOrder.Get("amount").Int(),
			RemainingAmount:  jsOrder.Get("remainingAmount").Int(),
			TotalAmount:      nil,
			Price:            jsOrder.Get("price").Float(),
		}

		// active
		active := jsOrder.Get("active")
		if !active.IsUndefined() {
			activeBool := active.Bool()
			order.Active = &activeBool
		}

		// roomName
		roomName := jsOrder.Get("roomName")
		if !roomName.IsUndefined() {
			roomNameStr := roomName.String()
			order.RoomName = &roomNameStr
		}

		// totalAmount
		totalAmount := jsOrder.Get("totalAmount")
		if !totalAmount.IsUndefined() {
			totalAmountInt := totalAmount.Int()
			order.TotalAmount = &totalAmountInt
		}

		result[key] = order
	}

	return result
}

func (market Market) CalcTransactionCost(amount int, roomName1 string, roomName2 string) int {
	return market.ref.Call("calcTransactionCost", amount, roomName1, roomName2).Int()
}

func (market Market) CancelOrder(orderId string) ErrorCode {
	result := market.ref.Call("cancelOrder", orderId).Int()
	return ErrorCode(result)
}

func (market Market) ChangeOrderPrice(orderId string, newPrice int) ErrorCode {
	result := market.ref.Call("changeOrderPrice", orderId, newPrice).Int()
	return ErrorCode(result)
}

type CreateOrderParams struct {
	Type         string
	ResourceType string
	Price        float64
	TotalAmount  int
	RoomName     *string
}

func (market Market) CreateOrder(params CreateOrderParams) ErrorCode {
	jsParams := map[string]interface{}{}
	jsParams["type"] = params.Type
	jsParams["resourceType"] = params.ResourceType
	jsParams["price"] = params.Price
	jsParams["totalAmount"] = params.TotalAmount
	if params.RoomName != nil {
		jsParams["roomName"] = *params.RoomName
	}
	return ErrorCode(market.ref.Call("changeOrderPrice", jsParams).Int())
}

func (market Market) Deal(orderId string, amount int, yourRoomName *string) ErrorCode {
	var jsYourRoomName js.Value
	if yourRoomName == nil {
		jsYourRoomName = js.Undefined()
	} else {
		jsYourRoomName = js.ValueOf(*yourRoomName)
	}
	result := market.ref.Call("deal", orderId, amount, jsYourRoomName).Int()
	return ErrorCode(result)
}

func (market Market) ExtendOrder(orderId string, addAmount int) ErrorCode {
	result := market.ref.Call("extendOrder", orderId, addAmount).Int()
	return ErrorCode(result)
}

func (market Market) GetAllOrders() []Order { // TODO - filter
	jsOrders := market.ref.Get("orders")
	length := jsOrders.Length()
	result := make([]Order, length)

	for i := 0; i < length; i++ {
		jsOrder := jsOrders.Index(i)
		order := Order{
			Id:               jsOrder.Get("id").String(),
			Created:          jsOrder.Get("created").Int(),
			CreatedTimestamp: nil,
			Active:           nil,
			Type:             jsOrder.Get("type").String(),
			ResourceType:     ResourceConstant(jsOrder.Get("resourceType").String()),
			RoomName:         nil,
			Amount:           jsOrder.Get("amount").Int(),
			RemainingAmount:  jsOrder.Get("remainingAmount").Int(),
			TotalAmount:      nil,
			Price:            jsOrder.Get("price").Float(),
		}

		// roomName
		roomName := jsOrder.Get("roomName")
		if !roomName.IsUndefined() {
			roomNameStr := roomName.String()
			order.RoomName = &roomNameStr
		}

		result[i] = order
	}

	return result
}

type PriceHistory struct {
	ResourceType ResourceConstant //TODO - MarketResourceConstant
	Date         string
	Transactions int
	Volume       int
	AvgPrice     float64
	StddevPrice  float64
}

func (market Market) GetHistory(resourceType *ResourceConstant) []PriceHistory { // TODO - MarketResourceConstant}) {
	var jsResourceType js.Value
	if resourceType == nil {
		jsResourceType = js.Undefined()
	} else {
		jsResourceType = js.ValueOf(string(*resourceType))
	}

	jsHistory := market.ref.Call("getHistory", jsResourceType)
	length := jsHistory.Length()
	result := make([]PriceHistory, length)
	for i := 0; i < length; i++ {
		entry := jsHistory.Index(i)
		result[i] = PriceHistory{
			ResourceType: ResourceConstant(entry.Get("resourceType").String()),
			Date:         entry.Get("date").String(),
			Transactions: entry.Get("transactions").Int(),
			Volume:       entry.Get("volume").Int(),
			AvgPrice:     entry.Get("avgPrice").Float(),
			StddevPrice:  entry.Get("stddevPrice").Float(),
		}

	}
	return result
}

func (market Market) GetOrderById(id string) Order {
	jsOrder := market.ref.Call("getOrderById", id)
	order := Order{
		Id:               jsOrder.Get("id").String(),
		Created:          jsOrder.Get("created").Int(),
		CreatedTimestamp: nil,
		Active:           nil,
		Type:             jsOrder.Get("type").String(),
		ResourceType:     ResourceConstant(jsOrder.Get("resourceType").String()),
		RoomName:         nil,
		Amount:           jsOrder.Get("amount").Int(),
		RemainingAmount:  jsOrder.Get("remainingAmount").Int(),
		TotalAmount:      nil,
		Price:            jsOrder.Get("price").Float(),
	}

	// roomName
	roomName := jsOrder.Get("roomName")
	if !roomName.IsUndefined() {
		roomNameStr := roomName.String()
		order.RoomName = &roomNameStr
	}

	return order
}
