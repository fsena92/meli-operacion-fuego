package decoder

import (
	"strings"
)

/*GetMessage returns the message generated from the ship*/
func GetMessage(messages [][]string) (msg string){
	
	var firstMessage []string = messages[0]
	var otherMessages [][]string

	otherMessages = append(otherMessages, messages[1], messages[2])

	var completeMessage []string
	word := ""	
		for i := 0; i < len(firstMessage); i++ {
			word = firstMessage[i];
			if word == "" {
				for _, otherMessage := range otherMessages {
					word = otherMessage[i];
					if word != ""{
						break
					}
						
				}
			}
			completeMessage = append(completeMessage, word);
			
			if completeMessage[len(completeMessage) - 1] != "" && (i + 1) != len(firstMessage) {
				completeMessage = append(completeMessage, " ");
			}
		}
		msg = strings.Join(completeMessage, "")
		return strings.Join(strings.Fields(msg), " ")
		
}
/*ValidateMessages validates the length of the messages in the request*/
func ValidateMessages(messages [][]string) bool{
	
	if len(messages) != 3 {
		return false
	}
	
	firstMessageLength := len(messages[0])
	var actualMessageLength int

	for i:=1; i < len(messages); i++ {
		actualMessageLength = len(messages[i])
		if actualMessageLength != firstMessageLength {
			return false
		} 
	}
	return true
}

/*ValidateMessage validates the message generated */
func ValidateMessage(message string) bool{
	return message != ""
}

