package shema

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Model struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	gorm.Model
	Name string
}

func main() {
	//var app DataAepp
	db, err := gorm.Open("sqlite3", "data.app")
	defer db.Close()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("connected to database!")
	}

	return

}

//http://motion-express.com/blog/gorm:-a-simple-guide-on-crud
