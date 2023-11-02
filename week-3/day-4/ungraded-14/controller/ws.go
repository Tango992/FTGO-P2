package controller

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
	inventoryClients []*websocket.Conn
	crimeClients []*websocket.Conn
	randomClients []*websocket.Conn
)

func InventoryWebsocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	inventoryClients = append(inventoryClients, ws)

	for {
		var msgTmp any
		err := ws.ReadJSON(&msgTmp)
		if err != nil {
			c.Logger().Error(err)
		}
		
		msg :=  msgTmp.(map[string]any)
		sender := msg["name"].(string)
		message := msg["chat_message"].(string)

		for _, client := range inventoryClients {
			notification := fmt.Sprintf(`
				<div hx-swap-oob="beforeend:#notifications" >
					<p class="py-0 my-0"><small>%s</small></p>
					<p class="lead">%s</p>
				</div>`, sender, message)

			// Write
			if err := client.WriteMessage(websocket.TextMessage, []byte(notification)); err != nil {
				c.Logger().Error(err)
			}
		}
	}
}

func CrimeWebsocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	crimeClients = append(crimeClients, ws)

	for {
		var msgTmp any
		err := ws.ReadJSON(&msgTmp)
		if err != nil {
			c.Logger().Error(err)
		}
		
		msg :=  msgTmp.(map[string]any)
		sender := msg["name"].(string)
		message := msg["chat_message"].(string)

		for _, client := range crimeClients {
			notification := fmt.Sprintf(`
			<div hx-swap-oob="beforeend:#notifications" >
				<p class="py-0 my-0"><small>%s</small></p>
				<p class="lead">%s</p>
			</div>`, sender, message)

			// Write
			if err := client.WriteMessage(websocket.TextMessage, []byte(notification)); err != nil {
				c.Logger().Error(err)
			}
		}
	}
}

func RandomWebsocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	randomClients = append(randomClients, ws)

	for {
		var msgTmp any
		err := ws.ReadJSON(&msgTmp)
		if err != nil {
			c.Logger().Error(err)
		}
		
		msg :=  msgTmp.(map[string]any)
		sender := msg["name"].(string)
		message := msg["chat_message"].(string)

		for _, client := range randomClients {
			notification := fmt.Sprintf(`
			<div hx-swap-oob="beforeend:#notifications" >
				<p class="py-0 my-0"><small>%s</small></p>
				<p class="lead">%s</p>
			</div>`, sender, message)

			// Write
			if err := client.WriteMessage(websocket.TextMessage, []byte(notification)); err != nil {
				c.Logger().Error(err)
			}
		}
	}
}

