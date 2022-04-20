package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"

	"github.com/deFarro/fastpoke_backend.git/app/artist"
	"github.com/deFarro/fastpoke_backend.git/app/config"
)

type Database struct {
	DB       gorm.DB
	Settings config.Config
}

func (db Database) log(itemType, id, action string) {
	fmt.Printf("%s (id: %s) -> %s\n", itemType, id, action)
}

func NewDatabase(settings config.Config) (Database, error) {
	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
				settings.DatabaseAddr,
				settings.DatabaseUser,
				settings.DatabasePassword,
				settings.DatabaseName,
				settings.DatabasePort,
			),
		),
		&gorm.Config{},
	)
	if err != nil {
		panic("failed to connect database")
	}

	database := Database{
		DB:       *db,
		Settings: settings,
	}

	err = TryToCall(database.DropTables, 5)
	if err != nil {
		return Database{}, err
	}

	err = database.PrepopulateDatabase()
	if err != nil {
		return Database{}, err
	}

	return database, nil
}

func TryToCall(f func() error, attempts int) error {
	timer := time.NewTicker(time.Second)
	for {
		<-timer.C

		err := f()
		if err == nil {
			return nil
		}

		attempts--
		if attempts == 0 {
			return err
		}
	}
}

func (db *Database) PrepopulateDatabase() error {
	err := db.DB.Migrator().AutoMigrate(&artist.Artist{})

	if err == nil {
		err := db.DB.Migrator().CreateTable(&initialArtists)
		if err != nil {
			return err
		}
		fmt.Println("Database is populated with initial artists")
	} else {
		log.Println(err)
	}

	return nil
}

func (db *Database) DropTables() error {
	fmt.Println("Trying to drop the tables")

	for _, model := range []interface{}{&artist.Artist{}} {
		err := db.DB.Migrator().DropTable(model)
		if err != nil {
			return err
		}
	}
	fmt.Println("Tables were dropped")

	return nil
}
