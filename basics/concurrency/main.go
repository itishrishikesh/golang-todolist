package main

import (
	"fmt"
)

// Scenario: We're building an quiz show. Where participants can answer question at the same time.
var answers chan string = make(chan string, 10)
var questions chan string = make(chan string, 10)

func answerQuestion(user, question, answer string) {
	answers <- answer
	fmt.Println("User", user, "has answered", "\""+answer+"\"", "for question", question)
}

func sendQuestion(question string) {
	questions <- question
	fmt.Println("Sending question", "\""+question+"\"", "to all users")
}

func main() {
	fmt.Println("Welcome to goQuiz!")

	sendQuestion("What's the meaning of this?")
	sendQuestion("What do you mean?")

	// simulate user1
	go func() {
		question, ok := <-questions
		if !ok {
			fmt.Println("It's not ok here bro!")
		}
		answerQuestion("user2", question, "I'm stuborn, and I'll not answer any question.")
	}()

	//simulate user2
	go func() {
		question, ok := <-questions
		if !ok {
			fmt.Println("It's not ok here bro!")
		}
		answerQuestion("user2", question, "I've a cheat sheet accept all answers.")
	}()

	close(questions)

	// evaluate answers.
	for ans := range answers {
		fmt.Println("Your answer is incorrect", ans, "as there was no answer!")
		if len(answers) <= 0 {
			close(answers)
		}
	}
}
