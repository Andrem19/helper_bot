package timer

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Andrem19/Timer_Job/variables"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
var m sync.Mutex

func Timer(cmd []string, chat_id string) (string, error) {
	switch strings.ToLower(cmd[1]) {
	case "start":
		go Start(cmd[2], chat_id)
		return "Process is runing", nil
		
	case "stop":
		return Stop()
	default:
		return "", nil
	}
}

func Start(minutes string, chat_id string) {
	min, err := strconv.Atoi(minutes)
		if err != nil {
			fmt.Println(err.Error())
		}
		time_finish := time.Now().Unix()
		time_finish = time_finish + int64(min * 60)

		for time.Now().Unix() < time_finish {
			if variables.Stop_timer{
				break
			}
			time.Sleep(time.Second * 4)

			left := (time_finish - time.Now().Unix())

			minutes := left /60
			sec := left % 60
			ch_id, _ := strconv.Atoi(chat_id)
			msg := tgbotapi.NewMessage(int64(ch_id), fmt.Sprintf("Left to finish %d:%v\n", int(minutes), sec))
			m.Lock()
			variables.Bot.Send(msg)
			m.Unlock()
		}
		variables.Stop_timer = false
}

func Stop() (string, error) {
	variables.Stop_timer = true
	return "Timer successfuly stoped", nil
}