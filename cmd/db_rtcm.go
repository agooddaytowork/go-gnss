package main

import (
    "fmt"
    "github.com/umeat/go-gnss/cmd/database/models"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
    db, err := gorm.Open("sqlite3", "database/test.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    msg := models.Observation{SatelliteData: []models.SatelliteData{}}
    db.First(&msg).Preload("SignalData").Related(&msg.SatelliteData)
    fmt.Printf("%+v\n", msg)
}

