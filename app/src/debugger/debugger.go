package debugger

import (
	"log"
	"os"
	"runtime"
	"strconv"

	"time"
)

//context, title, error
func Log(errorType string, message string, err error) {

	where, fileName, line, ok := runtime.Caller(1)
	details := runtime.FuncForPC(where)

	functionName := ""

	if ok && details != nil {
		functionName = details.Name()
	}

	errorText := ""

	if err != nil {
		errorText = string(err.Error())
	}

	serviceName, _ := os.LookupEnv("SERVICE_NAME")

	functionCallLine := strconv.Itoa(line)

	errorText, _ = strconv.Unquote(errorText)

	messageSend := `{	"service":"` + serviceName + `",
						"type":"` + errorType + `",
						"file":"` + fileName + `",
						"line":"` + functionCallLine + `",
						"function":"` + functionName + `",
						"message":"` + string(message) + `",
						"error":"` + errorText + `",
						"timestamp":"` + string(time.Now().UTC().Format(time.RFC3339Nano)) + `"}`

	log.Print("\n\n-=>>>DEBUG:>>>>>>>>>")
	log.Print(messageSend)
	log.Print("<<<<<<<<<<<<<<<<<<<<<<<<</DEBUG<<<=-\n")

	//NOTE:
	// could add some function here to send info to some
	// kind of remote debug recipient e.g. ELK, Graylog
}
