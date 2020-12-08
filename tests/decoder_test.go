package tests

import (
	"github.com/fsena92/meli-operacion-fuego/decoder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMessageWithAllMessagesEquals(t *testing.T){
	var messages [][]string
	messages = append(messages, []string{"este", "", "", "mensaje", ""}, []string{"", "es", "", "", "secreto"}, []string{"este", "", "un", "", ""})
	var completeMessage string = decoder.GetMessage(messages)
	assert.Equal(t, "este es un mensaje secreto", completeMessage)
}

func TestGetMessageWithAllMessagesEqualsAndOneMessageBlank(t *testing.T){
	var messages [][]string
	messages = append(messages, []string{"este", "", "un", ""}, []string{"", "es", "", "mensaje"}, []string{"", "", "", ""})
	var completeMessage string = decoder.GetMessage(messages)
	assert.Equal(t, "este es un mensaje", completeMessage)
}

func TestGetMessageWithAllMessagesStartOffsetEquals(t *testing.T){
	var messages [][]string
	messages = append(messages, []string{"", "este", "", "", "mensaje", ""}, []string{"","", "es", "", "", "secreto"}, []string{"", "este", "", "un", "", ""})
	var completeMessage string = decoder.GetMessage(messages)
	assert.Equal(t, "este es un mensaje secreto", completeMessage)
}

func TestGetMessageWithAllMessagesEndOffsetEquals(t *testing.T){
	var messages [][]string
	messages = append(messages, []string{"este", "", "", "mensaje", "", ""}, []string{"", "es", "", "", "secreto", ""}, []string{"este", "", "un", "", "", ""})
	var completeMessage string = decoder.GetMessage(messages)
	assert.Equal(t, "este es un mensaje secreto", completeMessage)
}

func TestGetMessageWithAllMessagesMiddleOffsetEquals(t *testing.T){
	var messages [][]string
	messages = append(messages, []string{"este", "", "", "mensaje", "", ""}, []string{"", "es", "", "", "", "secreto"}, []string{"este", "", "un", "", "", ""})
	var completeMessage string = decoder.GetMessage(messages)
	assert.Equal(t, "este es un mensaje secreto", completeMessage)
}
func TestGetMessageWithAllMessagesNotEquals(t *testing.T){
	var messages [][]string
	messages = append(messages, []string{"este", "", "", "mensaje"}, []string{"", "es", "", ""}, []string{"este", "", "un", ""})
	var completeMessage string = decoder.GetMessage(messages)
	assert.NotEqual(t, "este es un mensaje secreto", completeMessage)
}

func TestValidateMessagesWithMissingMessageInMessages(t *testing.T){
	var messages [][]string
	messages = append(messages, []string{"este", "", "", "mensaje"}, []string{"", "es", "", ""})
	assert.False(t, decoder.ValidateMessages(messages))
}
func TestValidateMessagesWithAllAndSameMessagesLength(t *testing.T){
	var messages [][]string
	messages = append(messages, []string{"este", "", "", "mensaje"}, []string{"", "es", "", ""}, []string{"este", "", "un", ""})
	assert.True(t, decoder.ValidateMessages(messages))
}
func TestValidateMessagesWithDiferentMessagesLength(t *testing.T){
	var messages [][]string
	messages = append(messages, []string{"este", "", "", "mensaje"}, []string{"", "es", "", "", "secreto"}, []string{"este", "", "un", "", ""})
	assert.False(t, decoder.ValidateMessages(messages))
}
func TestValidateMessageWithValidMessage(t *testing.T){
	message:= "este es un mensaje"
	assert.True(t, decoder.ValidateMessage(message))
}

func TestValidateMessageWithInvalidMessage(t *testing.T){
	message := ""
	assert.False(t, decoder.ValidateMessage(message))
}