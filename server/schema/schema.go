package schema

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"time"
)

var DB *gorm.DB

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

	} else {

		log.Println("connected to database!")

		DB = db

	}

}

/*http://motion-express.com/blog/gorm:-a-simple-guide-on-crud*/
