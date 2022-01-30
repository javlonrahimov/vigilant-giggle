package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var sqliteDatabase *sql.DB

func CreateDatabase() {
	os.Remove("sqlite-database.db") // I delete the file to avoid duplicated records.
	// SQLite is a file based database.

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ = sql.Open("sqlite3", "./sqlite-database.db")
}

func CreateTable(tableName string) {
	createStudentTableSQL := fmt.Sprintf(`CREATE TABLE %s (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,	
		"idStudent" TEXT,	
		"word" TEXT
	  );`, tableName)

	log.Printf("Create %s table...", tableName)
	statement, err := sqliteDatabase.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Printf("%s table created", tableName)
}

// We are passing db reference connection from main to our method with other parameters
func InsertWord(idStudent, word, tableName string) {
	log.Println("Inserting student record ...")
	insertStudentSQL :=  `INSERT INTO %s(idStudent, word) VALUES (?, ?)`
	statement, err := sqliteDatabase.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(idStudent, word)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func DisplayStudentWords(id, tablename string) []string {
	log.Println("Getting student record ...")
	query := fmt.Sprintf("SELECT * FROM %s", tablename)
	row, err := sqliteDatabase.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	words := make([]string, 10)
	for row.Next() { // Iterate and fetch the records from result cursor
		var idStudent string
		var word string
		row.Scan(&idStudent, &word)
		log.Printf("Table: %s, id: %s, word: %s", tablename, idStudent, word)
		words = append(words, word)
	}
	return words
}
