package app

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"path/filepath"
	wall_app "wall-server/wall-app"
)

type App struct {
	config                 Config
	server                 string
	httpConnectionUpgraded websocket.Upgrader
	clients                []*Client
	ActionsWorker          *ActionsWorker
	TriggersWorker         *TriggersWorker
}

type WebSocket struct {
}

func Start(actions []*ActionHandler, triggers []*TriggerHandler) (App, error) {
	//Try to load configuration
	path := filepath.Join(".", "config.json")
	config, err := LoadConfig(path)
	if err != nil {
		return App{}, err
	}

	app := App{config: config}
	//Start application
	app.runApp(actions, triggers)
	//Up server and handle controller
	err = app.serverUp()
	if err != nil {
		return App{}, err
	}
	return app, nil
}

func (app *App) registerTriggers(triggers []*TriggerHandler) {
	for _, trigger := range triggers {
		app.TriggersWorker.registerHandler(trigger)
	}
}

func (app *App) registerActions(actions []*ActionHandler) {
	for _, action := range actions {
		app.ActionsWorker.registerHandler(action)
	}
}

func (app *App) registerWorkers() {
	app.ActionsWorker = &ActionsWorker{}
	app.TriggersWorker = &TriggersWorker{}
}

func (app *App) runApp(actions []*ActionHandler, triggers []*TriggerHandler) {
	//Up inmemory storage for records
	wall_app.CreateWallList(app.config.RecordLimit)

	app.registerWorkers()
	app.registerActions(actions)
	app.registerTriggers(triggers)
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
		wall_app.GetAppInstance().AddClient(&wall_app.Client{
			Conn: client.Conn,
		})
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
