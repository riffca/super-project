package schema

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"time"
)

var DB *gorm.DB
var Connected bool

type Model struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func init() {

	db, err := gorm.Open("sqlite3", "data")
	db.LogMode(true)
	defer db.Close()

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("connected to database!")

	DB = db
	Connected = true

	//db.CreateTable(&User{}, &Page{})

}

/*http://motion-express.com/blog/gorm:-a-simple-guide-on-crud*/
