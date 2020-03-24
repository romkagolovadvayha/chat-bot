package main

import (
    "bufio"
    "os"
    "log"
    "chat-bot/chat"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    c := chat.NewChat()
    c.InitAnswers()

    for scanner.Scan() {
        str := scanner.Text()
        c.Message(str)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalln(err)
    }

}
