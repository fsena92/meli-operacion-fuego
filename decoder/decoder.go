package decoder

/*Retorna el mensaje tal cual lo genera el emisor del mensaje*/
func GetMessage(messages [][]string) (msg string){
	var min int
	//length del mas chico
	for i, message := range messages {
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

	var originalMessage []string
	//reconstruir msj
	for i:= 0; i< len(messages); i++ {
		for j:= 0; j < len(messages[i]); j++ {
			if messages[i][j] != "" {
				originalMessage[j] = messages[i][j]
			}
		}
	}

	return "";
}