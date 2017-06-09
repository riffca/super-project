package schema

import (
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

const (
	driver  = "mysql"                                             //"sqlite3"
	options = "root:@/work?charset=utf8&parseTime=True&loc=Local" //"data.db"
)

var DB *gorm.DB

type Model struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func ConnectDB() *gorm.DB {

	db, err := gorm.Open("mysql", "root:@/work?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Println(err)
		//panic(err)
		return db
	}

	db.LogMode(true)
	defer db.Close()

	log.Println("connected to database!")

	//db.CreateTable(&Page{})
	//db.DropTableIfExists(&User{}, &Page{})
	db.Create(&Page{Name: "sqwtass", Content: "stas"})

	return db

}

/*http://motion-express.com/blog/gorm:-a-simple-guide-on-crud*/
