package app

import (
	"fmt"
	"net"
	"path/filepath"
	"time"
)

type App struct {
	config Config
	server string
}

func Start() (App, error) {
	//Try to load configuration
	path := filepath.Join(".", "config.json")
	config, err := LoadConfig(path)
	if err != nil {
		return App{}, err
	}

	app := App{config: config}

	//Up server and handle controller
	err = app.serverUp()
	if err != nil {
		return App{}, err
	}
	return app, nil
}

func (app *App) serverUp() error {
	fmt.Println("Start server")
	server, err := net.Listen("tcp", app.config.GetServerString())
	if err != nil {
		return fmt.Errorf("server can't start: %v", err)
	}
	defer server.Close()

	fmt.Println("Listening on " + app.config.GetServerString())
	fmt.Println("Waiting for clients...")
	for {
		connection, err := server.Accept()
		if err != nil {
			return fmt.Errorf("can't accept: %v", err)
		}
		client := Client{Info: connection.RemoteAddr().String()}
		fmt.Printf("Client [%s] <%s> was connected", time.Now().String(), client.Info)
		go client.Handler(connection)
	}
}