package main

func main() {
	// Conexão com o banco de dados SQLite3
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
}
defer db.Close()

BookService :=
}