package chat

import (
    "fmt"
)

type Chat struct {
    answers *Answers
}

func NewChat() *Chat {
    c := new(Chat)
    return c
}

func (c *Chat) InitAnswers() {
    as := NewAnswers()
    as.SetAnswer(NewAnswer("Привет", "и тебе привет"))
    as.SetAnswer(NewAnswer("Как дела?", "хорошо, а у тебя?"))
    as.SetAnswer(NewAnswer("хорошо", "что делаешь?"))
    c.answers = as
}

/**
* Обработка сообщения
*/
func (c *Chat) Message(message string) {
    fmt.Println("Ответ:", c.answers.Answer(message))
}