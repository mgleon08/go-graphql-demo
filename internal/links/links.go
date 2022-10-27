package links

import (
	"fmt"
	"log"

	database "example/internal/pkg/db/mysql"
	"example/internal/users"
)

// #1
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
	Userid    string
}

//#2
func (link Link) Save() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address,UserID) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(link.Title, link.Address, link.Userid)
	if err != nil {
		log.Fatal(err)
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

func GetAll() []Link {
	stmt, err := database.Db.Prepare("select id, title, address, userid from Links")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &link.Userid)
		if err != nil{
			log.Fatal(err)
		}

		fmt.Println("查詢使用者成功", link.Userid)

		if len(link.Userid) != 0 {
			var user users.User
			row := database.Db.QueryRow("select * from Users where id=?", link.Userid)
			if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
				log.Fatal(err)
			}
			fmt.Println("查詢使用者成功", user)
			link.User = &user
		}

		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	log.Print("Row search!")
	return links
}
