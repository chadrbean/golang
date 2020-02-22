package main

import (
	"fmt"
	"strings"
	"database/sql"
	"strconv"
	"os/exec"
	"os"
	"log"
	"io/ioutil"
	_ "github.com/mattn/go-sqlite3"
)

func hash() []byte {
	cmd := "find . -ls | sort | sha1sum"
    cmdOutput, err := exec.Command("bash", "-c", cmd).Output()
    if err != nil {
        log.Fatal(err)
    	}
	return cmdOutput
	}
func walkDir() []os.FileInfo {
	files, err := ioutil.ReadDir(".")
	if err != nil {
    	log.Fatal(err)
		}
	return files
	}
func bytesToString(data []byte) string {
	return string(data[:])
	}
func main() {
	cmdOutput := hash()
	files := walkDir()
	for _, f := range files {
		fmt.Println(f.Name())
		}
	hashString := bytesToString(cmdOutput)
	hashStrip := strings.Fields(hashString)[0]

	fmt.Println(hashStrip)

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
