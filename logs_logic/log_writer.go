package logs_logic

import (
	"log"
	"os"
)

type Simple_log struct {
	Message    string
	Query_type string
	Query_id   string
	Answer     string
}

func Write_usual_log(sim_log Simple_log) {
	var FileLog, err = os.OpenFile("logs\\log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755) // TODO делай разбивку файлов по датам
	// os.Create("C:\\Users\\Лёшка\\GolandProjects\\graphQL_go\\logs\\log.txt")
	if err != nil {
		println(err.Error())
	}
	infoLog := log.New(FileLog, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Print(sim_log)

}
