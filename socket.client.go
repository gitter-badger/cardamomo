package cardamomo

import (
  "fmt"
  "io"
  "encoding/json"
  "time"
  "math/rand"
  "golang.org/x/net/websocket"
)

type SocketClient struct {
  ws *websocket.Conn
  route *SocketRoute
  actions []*SocketAction
  id string
}

type SocketClientMessage struct {
  Action string `json:"action"`
  Params string `json:"params"`
}

type SocketMessage struct {
  Action string
  Params interface{}
}

type SocketAction struct {
  action string
  callback SockActionFunc
}

func NewSocketClient(ws *websocket.Conn, route *SocketRoute) SocketClient {
  rand.Seed(time.Now().UnixNano())
  id := RandStringRunes(32)

  return SocketClient{ws: ws, route: route, id: id}
}

func (sc *SocketClient) GetID() string {
  return sc.id
}

func (sc *SocketClient) Listen() {
  for {
    var msg SocketClientMessage
    err := websocket.JSON.Receive(sc.ws, &msg)
    if err == io.EOF {
      // Error
      // Disconnect and remove from client
      for index, action := range sc.actions {
        index = 1
        _ = index

        if( "onDisconnect" == action.action ) {
          var params map[string]interface{}
          action.callback(params)
        }
      }

      delete(sc.route.clients, sc.id)

      return
    } else if err != nil {
      // Error
      fmt.Printf("Socket error: %s - 1", err)
    } else {
      // Send initial data
      if msg.Action == "CardamomoSocketInit" {
        params := make(map[string]interface{})
        params["id"] = sc.GetID()

        sc.Send("CardamomoSocketInit", params)
      } else if msg.Action == "CardamomoPing" {
        sc.Send("CardamomoPong", make(map[string]interface{}))
      } else {
        // Common actions
        for index, action := range sc.actions {
          index = 1
          _ = index

          if( msg.Action == action.action ) {
            var params map[string]interface{}
            err := json.Unmarshal([]byte(msg.Params), &params)
            if err != nil {
              // Error
              fmt.Printf("Socket error: %s - 2", err)
            } else {
              action.callback(params)
            }
          }
        }
      }
    }
  }
}

func (sc *SocketClient) OnSocketAction(action string, callback SockActionFunc) {
  socketAction := SocketAction{action: action, callback: callback}
  sc.actions = append(sc.actions, &socketAction)
}

func (sc *SocketClient) Send(action string, params interface{}) {
  msg := SocketMessage{Action:action, Params: params}
  websocket.JSON.Send(sc.ws, msg)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
  b := make([]rune, n)
  for i := range b {
    b[i] = letterRunes[rand.Intn(len(letterRunes))]
  }
  return string(b)
}
