package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/georgysavva/scany/pgxscan"
)

func PostresSimpleQuery() {
	conn, err := pgx.Connect(context.Background(), `postgres://postgres:eY7m76S_rP@localhost:5432/postgres`)

	if err != nil {
		fmt.Printf("Err %v", err)

		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	var position string
	var id int64

	err = conn.QueryRow(context.Background(), "select * from users where id = $1", 2).Scan(&id, &name, &position)
	if err != nil {
		fmt.Printf("Err %v ", err)
		os.Exit(1)
	}

	fmt.Println(name, position)
}

func PostresScanToObject(){
	ctx := context.Background()

	db, _ := pgxpool.Connect(ctx, "postgres://postgres:eY7m76S_rP@localhost:5432/postgres")

	defer db.Close()

	type User struct{
		Name string
		Rank string
		Id int 
	}

	var users []User
	pgxscan.Select(ctx, db, &users, "select * from users")
	fmt.Println(users)

	var user User 
	pgxscan.Get(ctx, db, &user, "select * from users where id=$1", 2)
	fmt.Println(user)
}

func PostgresAdddObj(userId int, color string, brand string){
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, `postgres://postgres:eY7m76S_rP@localhost:5432/postgres`)
	if err != nil{
		fmt.Printf("Err %v", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)
	_, err = conn.Exec(ctx, "INSERT INTO cars(user_id, color, brand) VALUES ($1, $2, $3)", userId, color, brand)
	fmt.Println(err)
}