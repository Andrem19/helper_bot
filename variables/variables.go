package variables

import (
	"database/sql"
	"log"

	db "github.com/Andrem19/Timer_Job/db/sqlc"
	_ "github.com/lib/pq"
	
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Conf struct {
	TELEGTRAM_BOT_TOKEN string `mapstructure:"TELEGTRAM_BOT_TOKEN"`
	USER_ID string `mapstructure:"USER_ID"`
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

var (
	Bot* tgbotapi.BotAPI
	Config Conf
	Stop_timer bool = false
	Queries *db.Queries
	database *sql.DB
	
)

func StartWithDb() {
	database, err := sql.Open(Config.DBDriver, Config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	Queries = db.New(database)
}