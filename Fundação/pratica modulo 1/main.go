package main

import (
	"fmt"
	pedidopackage "pratica-1/pedido"

	"github.com/google/uuid"
)

func main() {
	pedidoService := pedidopackage.NewPedidoService()

	pedidoService.Create(1)

	pedido, err := pedidoService.GetById(uuid.NewString())
	if err != nil {
		panic(err)
	}
	println(pedido.Id.String())
	println(pedido.ClienteId)
	println(pedido.OrderCode)
	pedidos, err := pedidoService.GetAll()
	if err != nil {
		panic(err)
	}
	for key, pedido := range pedidos {
		println(pedido.Id.String())
		println(pedido.ClienteId)
		println(pedido.OrderCode)
		delete(pedidos, key)
	}
	println("Total de pedidos: ", len(pedidos))

	println("Total de pedidos: ", len(pedidos))

	idNovo := uuid.NewString()
	pedidos[idNovo] = &pedidopackage.Pedido{
		Id:        uuid.New(),
		ClienteId: 1,
		OrderCode: uuid.NewString(),
		Itens: []pedidopackage.PedidoItem{
			{
				Id:         uuid.New(),
				PedidoId:   uuid.New(),
				ProdutoId:  1,
				Quantidade: 1,
				Valor:      10.0,
				ValorTotal: 10.0,
			},
		},
	}

	fmt.Printf("Pedidos após inclusão: %+v", *pedidos[idNovo])

}
