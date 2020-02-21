package main

import (
	"fmt"
	"database/sql"
	"strconv"
	"os/exec"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func hash() {
	cmd := "find testsrc -ls | sort | sha1sum"
    cmdOutput, err := exec.Command("bash", "-c", cmd).Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", cmdOutput)

}

func main() {
	hash()
	database, _ := 
		sql.Open("sqlite3", "./bogo.db")
	statement, _ := 
		database.Prepare("CREATE TABLE IF NOT EXISTS movies (id INTEGER PRIMARY KEY, movie TEXT, folderHash TEXT, datetime TEXT )")
	statement.Exec()
	statement, _ = 
		database.Prepare("INSERT INTO movies (movie, folderHash, datetime) VALUES (?, ?, ?)")
    statement.Exec("TIME MACHINE", "12323sdfgsdfg345", "2020202001")
	rows, _ := 
		database.Query("SELECT id, movie, folderHash, datetime FROM movies")
	
    var id int
    var movie string
	var folderHash string
	var datetime string
    for rows.Next() {
        rows.Scan(&id, &movie, &folderHash, &datetime)
        fmt.Println(strconv.Itoa(id) + ": " + movie + " " + folderHash + " " + datetime)
	}
	database.Exec("DELETE FROM movies;")
}
