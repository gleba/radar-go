package models

import (
	"time"
)

type ActiveAlert struct {
	Id             int
	CoinName       string
	AlertId        int
	DetectTime     time.Time
	LastActiveTime time.Time
}

type Market struct {
	Slug string
	Name string
}

func Restore() {
	//aa :=[]ActiveAlert{}
	//service.DB.Model(&aa).Where()
}
