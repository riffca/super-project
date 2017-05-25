package shema

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	//"os"
	"time"
)

var DB *gorm.DB

type Model struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	Name string
}

type Page struct {
	Content string
}

func init() {

	db, err := gorm.Open("sqlite3", "data.app")
	db.LogMode(true)
	defer db.Close()

	if err != nil {

		log.Println(err)

	} else {

		log.Println("connected to database!")

		DB = db

	}

	//DB.DropTable(&User{}, &Page{})
	//DB.CreateTable(&User{}, &Page{})

	// if len(os.Args) >= 0 {
	//  switch os.Args[1] {
	//  case "create":
	//    switch os.Args[2] {
	//    case "user":
	//      DB.CreateTable(&User{})
	//    case "page":
	//      DB.CreateTable(&Page{})
	//    case "all":
	//      DB.CreateTable(&User{}, &Page{})
	//    }
	//  case "drop":
	//    switch os.Args[2] {
	//    case "user":
	//      DB.DropTable(&User{})
	//    case "page":
	//      DB.DropTable(&Page{})
	//    case "all":
	//      DB.DropTable(&User{}, &Page{})
	//    }
	//  }

	// }

}

func New() gorm.DB {
	return &DB
}

//http://motion-express.com/blog/gorm:-a-simple-guide-on-crud
