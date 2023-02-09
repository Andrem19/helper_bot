package variables

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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
	
)