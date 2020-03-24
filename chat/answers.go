package chat

import (
)

type Answers struct {
    answers []*Answer
}

type Answer struct {
    keywords string
    message string
}

func NewAnswers() *Answers {
    as := new(Answers)
    return as
}

func NewAnswer(keywords string, message string) *Answer {
    a := new(Answer)
    a.keywords = keywords
    a.message = message
    return a
}

func (as *Answers) SetAnswer(answer *Answer) {
    as.answers = append(as.answers, answer)
}

// делаем всякие расчеты и отдаем нужно сообщение
func (as *Answers) Answer(clientMessage string) string {
    message := ""
    for i := 0; i < len(as.answers); i++ {
        answer := as.answers[i]
        if answer.keywords == clientMessage {
           message = answer.message
        }
    }
    return message
}