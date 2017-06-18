package schema

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

type Model struct {
	ID        uint64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func mysql() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/work?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

/*http://motion-express.com/blog/gorm:-a-simple-guide-on-crud*/
