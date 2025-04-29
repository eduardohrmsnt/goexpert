package pedido

import "github.com/google/uuid"

type Pedido struct {
	Id        uuid.UUID
	ClienteId int
	OrderCode string
	Itens     []PedidoItem
}

type PedidoItem struct {
	Id         uuid.UUID
	PedidoId   uuid.UUID
	ProdutoId  int
	Quantidade int
	Valor      float64
	ValorTotal float64
}
