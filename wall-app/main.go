package wall_app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

var wallApp WallList

type Client struct {
	Conn *websocket.Conn
}

type WallList struct {
	clients  []*Client
	wallList []*Record
	limit    int
	count    int
}

type Record struct {
	Date   time.Time `json:"date"`
	Text   string    `json:"text"`
	Client string    `json:"client"`
}

func CreateWallList(countOfRecords int) {
	wallApp = WallList{wallList: []*Record{}, limit: countOfRecords, count: 0}
}

func GetAppInstance() *WallList {
	return &wallApp
}

func (wl *WallList) CreateWallRecord(text string, clint string) *Record {
	return &Record{Text: text, Date: time.Now(), Client: clint}
}

func (wl *WallList) AddClient(client *Client) {
	wl.clients = append(wl.clients, client)
}

func (wl *WallList) GetRecords() []*Record {
	return wl.wallList
}

func (wl *WallList) PutRecord(r *Record) []*Record {
	if wl.count >= wl.limit {
		wl.wallList = append(wl.wallList[:], r)
		wl.count = 0
	} else {
		wl.wallList = append(wl.wallList, r)
	}

	wl.count++
	return wl.wallList
}

func (wl *WallList) GetClients() []*Client {
	return wl.clients
}

func (r *Record) ToJson() ([]byte, error) {
	jsonRecord, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("can't convert record to json")
	}
	return jsonRecord, nil
}
