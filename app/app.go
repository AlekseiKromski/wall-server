package app

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"path/filepath"
	"wall-server/app/actions"
	wall_app "wall-server/wall-app"
)

type App struct {
	config                 Config
	server                 string
	WallApp                *wall_app.WallList
	httpConnectionUpgraded websocket.Upgrader
	clients                []*Client
}

type WebSocket struct {
}

func Start() (App, error) {
	//Try to load configuration
	path := filepath.Join(".", "config.json")
	config, err := LoadConfig(path)
	if err != nil {
		return App{}, err
	}

	app := App{config: config}
	//Start application
	app.runApp()
	//Up server and handle controller
	err = app.serverUp()
	if err != nil {
		return App{}, err
	}
	return app, nil
}

func (app *App) runApp() {
	//Up inmemory storage for records
	app.WallApp = wall_app.CreateWallList(app.config.RecordLimit)
	actions.RegisterActions()
}

func (app *App) serverUp() error {
	fmt.Println("Start server")
	app.httpConnectionUpgraded = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := app.httpConnectionUpgraded.Upgrade(w, r, nil)
		if err != nil {
			fmt.Printf("problem while upgrade http connection to webscket: %v", err)
			return
		}
		fmt.Println("Client was connected")
		client := app.addClient(conn)
		client.Handler(app)
	})

	http.ListenAndServe(app.config.GetServerString(), nil)
	return nil
}

func (app *App) addClient(conn *websocket.Conn) *Client {
	client := CreateNewClient(conn, &app.config)
	app.clients = append(app.clients, client)
	return client
}

func (app *App) sendNewClientConnectedMessage(ip string) {
	for _, client := range app.clients {
		client.conn.WriteMessage(1, []byte(fmt.Sprintf("New User was connected <%s>", ip)))
	}
}
