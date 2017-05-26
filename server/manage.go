package main

import (
	schema "./shema"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

// package main

// import (
//   "fmt"
//   "os"

//   "github.com/urfave/cli"
// )

// func main() {
//   app := cli.NewApp()

//   app.Action = func(c *cli.Context) error {
//     fmt.Printf("Hello %q", c.Args().Get(0))
//     return nil
//   }

//   app.Run(os.Args)
// }

func main() {

	db, _ := gorm.Open("sqlite3", "data")
	db.LogMode(true)

	defer db.Close()

	if len(os.Args) >= 0 {
		switch os.Args[1] {
		case "create":
			switch os.Args[2] {
			case "user":
				db.CreateTable(&shema.User{})
			case "page":
				db.CreateTable(&shema.Page{})
			case "all":
				db.CreateTable(&shema.User{}, &shema.Page{})
			}
		case "drop":
			switch os.Args[2] {
			case "user":
				db.DropTable(&shema.User{})
			case "page":
				db.DropTable(&shema.Page{})
			case "all":
				db.DropTable(&shema.User{}, &shema.Page{})
			}
		}
	}

}
