package wall_app

import (
	"time"
)

type WallList struct {
	wallList []*Record
	limit    int
	count    int
}

type Record struct {
	date time.Time
	text string
}

func CreateWallRecord(text string) *Record {
	return &Record{text: text, date: time.Now()}
}

func CreateWallList(countOfRecords int) *WallList {
	return &WallList{wallList: []*Record{}, limit: countOfRecords, count: 0}
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
