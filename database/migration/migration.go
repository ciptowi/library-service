package migration

import (
	"fmt"
	"library-sevice/database"
	"library-sevice/internal/models"
	"reflect"
	"strings"
)

var tables = []interface{}{
	&models.User{},
	&models.Book{},
}

func Migrate() {

	conn := database.GetConnection()

	if err := conn.AutoMigrate(tables...); err != nil {
		panic(err)
	}
}

func Rollback() {

	conn := database.GetConnection()

	for i := len(tables) - 1; i >= 0; i-- {
		if err := conn.Migrator().DropTable(tables[i]); err != nil {
			panic(err)
		}
	}
}

func Status() {
	var (
		conn        = database.GetConnection()
		colorReset  = "\033[0m"
		colorGreen  = "\033[32m"
		colorYellow = "\033[33m"
	)

	fmt.Printf("In database %s:\n", conn.Migrator().CurrentDatabase())
	for _, table := range tables {
		var name string

		t := reflect.TypeOf(table)
		if t.Kind() == reflect.Ptr {
			name = strings.ToLower(t.Elem().Name()) + "s"
		} else {
			name = strings.ToLower(t.Name()) + "s"
		}

		if conn.Migrator().HasTable(table) {
			fmt.Println("\t", name, "===>", string(colorGreen), "migrated", string(colorReset))
		} else {
			fmt.Println("\t", name, "===>", string(colorYellow), "not migrated", string(colorReset))
		}
	}
}
