package logconfig

import (
	"log"
	"os"
)

func init() {
	_, err := os.Stat("./SendEmailLog")
	if err != nil {
		os.MkdirAll("./SendEmailLog", os.ModePerm)
	}
	logfile, err := os.OpenFile("./SendEmailLog\\log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		//fmt.Println(err)
		return
	}
	log.SetOutput(logfile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
