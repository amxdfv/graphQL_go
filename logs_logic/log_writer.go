package logs_logic

import (
	"encoding/json"
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
	if err != nil {
		println(err.Error())
	}
	infoLog := log.New(FileLog, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Print(sim_log)

}

func Set_log_information(msg string, q_t string, q_i string, data any) Simple_log {
	var Log_inf Simple_log
	Log_inf.Message = msg
	Log_inf.Query_type = q_t
	Log_inf.Query_id = q_i
	var val []byte
	val, _ = json.Marshal(data)
	Log_inf.Answer = string(val[:])
	return Log_inf
}
