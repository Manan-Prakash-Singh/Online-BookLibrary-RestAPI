package models

type Book struct {
	BookID int    `json:"book_id"`
	Name   string `json:"name"`
	Genre  string `json:"genre"`
	Author string `json:"author"`
	Price  int    `json:"price"`
	Count  int    `json:"count"`
}

func AddBook(b *Book) error {
	_, err := db.Exec(`INSERT INTO books(name, genre, author, price_inr, count)
    VALUES ($1,$2,$3,$4,$5)`,
		b.Name,
		b.Genre,
		b.Author,
		b.Price,
		b.Count,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteBook(bookid int) error {
	_, err := db.Exec(`DELETE FROM books where bookid = $1`, bookid)
	if err != nil {
		return err
	}
	return nil
}

func GetBookByID(bookid int) (*Book, error) {

	row := db.QueryRow(`SELECT * FROM books WHERE bookid = $1`, bookid)

	book := &Book{}

	err := row.Scan(
		&book.BookID,
		&book.Name,
		&book.Genre,
		&book.Author,
		&book.Price,
		&book.Count,
	)

	if err != nil {
		return nil, err
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return book, err
}

func GetTotalBooksCount() (int, error) {
	row := db.QueryRow(`SELECT SUM(count) AS total_books FROM books`)
	var total_count int
	if err := row.Scan(&total_count); err != nil {
		return 0, nil
	}
	return total_count, nil
}

func GetAllBooks() ([]Book, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	var books []Book

	for rows.Next() {
		book := Book{}
		err := rows.Scan(
			&book.BookID,
			&book.Name,
			&book.Genre,
			&book.Author,
			&book.Price,
			&book.Count,
		)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
