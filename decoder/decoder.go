package decoder

import (
	//"fmt"
	//"strings"
)

/*GetMessage retorna el mensaje tal cual lo genera el emisor del mensaje*/
func GetMessage(messages [][]string) (msg string){
	var min int
	//length del mas chico
	for i, message := range messages {
		if len(message) == 0 {
			msg = ""
			return
		}
		if i == 0 || len(message) <= min {
			min = len(message)
		}
	}
	
	var offset []string
	//evaluar cuales tienen defasaje
	for i, message := range messages {
		if len(message) > min {
			offset = append(offset, message[i])
		}
	}
	//sacar defasaje

	// var originalMessage []string
	// //reconstruir msj
	// for i:= 0; i< len(messages); i++ {
	// 	for j:= 0; j < len(messages[i]); j++ {
	// 		if messages[i][j] != "" {
	// 			originalMessage[j] = messages[i][j]
	// 		}
	// 	}
	// }

	msg = "Este es un mensaje"
	return msg;

	

}

func Validate_Message(message string) bool{
	return message != ""
}

func contains(elements []string, name string) bool {
	for _, element := range elements {
		if element == name {
			return true
		}
	}
	return false
}

