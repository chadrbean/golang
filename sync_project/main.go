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
	"path/filepath"
	_ "github.com/mattn/go-sqlite3"
)

func bytesToString(data []byte) string {
	return string(data[:])
	}


func hash(dir string) string{
	cmd := "find " + string(dir[:]) + " -type f -execdir ls -alh --time-style='+' '{}' ';' | sort "
    cmdOutput, err := exec.Command("bash", "-c", cmd).Output()
    if err != nil {
        log.Fatal(err)
		}
	hashString := bytesToString(cmdOutput)
	hashStrip := strings.Fields(hashString)[0]
	return hashStrip
	}

func walkDir(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
    	log.Fatal(err)
		}
	return files
	}
func walkAllFilesInDir(dir string) error {
	var files []string
	return filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		// check if it is a regular file (not dir)
		if info.Mode().IsRegular() {
			files = append(files, path)
			for _, file := range files {
				cmd := "sha1sum " + file
				fmt.Println(cmd)
				fileHash, _ := exec.Command("bash", "-c", cmd).Output()
				fmt.Println(bytesToString(fileHash))
				}	
		}
		return nil
	})
}

func database() {
	database, _ := 
		sql.Open("sqlite3", "./movies.db")
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


func main() {
	srcdir := "src/"
	dstdir := "dst/"
	srchash := hash(srcdir)
	dsthash := hash(dstdir)
	
	// Check and see if source hash matches dest hash and sync if not. Hash can be modified to whatever
	if srchash == dsthash {
		walkAllFilesInDir(srcdir)
		database()
		fmt.Printf("%s %s", srchash, dsthash)
	} else {
		fmt.Println("All Files Are Synced")
	}


}
