package internal

import "database/sql"

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

type BookService struct {
	db *sql.DB
}

func (s *BookService) CreateBook(boo *Book) error {
	query := "Insert into books (title, author, genre) values(?,?,?)"
	result, err := s.db.Exec(query, book.Title, book.Author, book.Genre)
	if err != nil {
		return err
	}
	result.LastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	book.ID = int(LastInsertId())
	return nil
}

func (s*BookService) getBooks() ([]Book, error){
	query:= "Select idm title, author, genre from books"
	rows, err:=s.db.Query(query)
	if err != nill{
		return nill, err
	}
	var books []Book
	fo rows.Next(){
		var book Book 
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
func (s *BookService) GetBookByID(id int) (*Book, error) {
	query := "select id, title, author, genre from books where id = ?"
	row:= s.db.QueryRow(query, id)
	
	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
	if err != nil {
		return nil, err 
	}
	return &book, nil
}

func (s *BookService) UpdateBook(book *Book) error{
	query := "update books set title=?, author?, genre=?, where id=?"
	result, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.ID)
	return err

}

func (s *BookService) DeleteBook(id int) error {
	query := "delete from books whre id?"
	_, err := s.db.Exec(query, id)
	return err
}