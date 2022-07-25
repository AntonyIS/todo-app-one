package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/AntonyIS/todo-app-one/app"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}

var (
	host      = os.Getenv("HOST")
	port      = os.Getenv("PORT")
	user      = os.Getenv("USER")
	password  = os.Getenv("PASSWORD")
	dbname    = os.Getenv("DBNAME")
	tableName = os.Getenv("TABLENAME")
)

type userRepository struct {
	client *sql.DB
}

func newPostgresClient() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=verify-full", host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("ERROR ACESSING DB :", err)

	}
	err = db.Ping()

	if err != nil {
		log.Fatalf("ERROR CONNECTING TO DB :", err)
	}
	return db, nil
}

func NewPostgresRepository() (app.UserRepository, error) {
	repo := userRepository{}
	client, err := newPostgresClient()
	if err != nil {
		return nil, errors.Wrap(err, "user.NewPostgresRepository")
	}
	repo.client = client

	return repo, nil
}

func (u userRepository) Create(user *app.User) (*app.User, error) {
	insert := fmt.Sprintf("INSERT INTO %s values ('%s','%s','%s','%s','%s','%s','%s');", os.Getenv("TABLENAME"), user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.Avater, user.Todo)
	_, err := u.client.Exec(insert)
	if err != nil {
		return nil, errors.Wrap(err, "user.Create")
	}
	return user, nil
}

func (u userRepository) Read(id string) (*app.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", os.Getenv("TABLENAME"), id)
	user := &app.User{}
	row := u.client.QueryRow(query, id)

	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Avater, &user.Todo)

	if err != nil {
		return nil, errors.Wrap(err, "user.Read")
	}
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return user, nil

}

func (u userRepository) ReadAll() (*[]app.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", os.Getenv("TABLENAME"))
	users := []app.User{}
	rows, err := u.client.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "user.ReadAll")
	}

	defer rows.Close()

	for rows.Next() {
		var id, firstname, lastname, email, avater, todo string
		err = rows.Scan(&id, firstname, lastname, email, avater, todo)
		user := app.User{}
		if err != nil {
			return nil, errors.Wrap(err, "user.ReadAll")
		}

		users = append(users, user)
	}

	return &users, nil

}

func (u userRepository) Update(user *app.User) (*app.User, error) {
	query := fmt.Sprintf("UPDATE %s SET firstname=$2, lastname=$3 email=$4 avater=$4")

	_, err := u.client.Exec(query, user.FirstName, user.LastName, user.Email, user.Avater)

	if err != nil {
		return nil, errors.Wrap(err, "user.Update")
	}
	return user, nil

}

func (u userRepository) Delete(id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, os.Getenv("TABLENAME"))
	_, err := u.client.Exec(query, id)

	if err != nil {
		return errors.Wrap(err, "user.Update")
	}
	return nil
}
