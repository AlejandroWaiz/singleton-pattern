package singleton

import (
	"fmt"
	"sync"
)

//Its safe thanks to Mutual exception, so no other GO Routine can access the variable at the same time
type prisonerPrincess struct {
}

var lock = &sync.Mutex{}

var peachToadstool *prisonerPrincess

func defeatBowser() {
	fmt.Println("You've defeated Bowser... Nice job!")
}

func RescueThePrincess() *prisonerPrincess {

	if peachToadstool == nil {

		lock.Lock()
		defer lock.Unlock()

		if peachToadstool == nil {

			defeatBowser()

			peachToadstool = &prisonerPrincess{}

			fmt.Println("Thanks for rescuing me! Lets have a good night baby ;)")

		} else {

			fmt.Println("Sorry, the princess is in another castle.")

		}

	} else {

		fmt.Println("Sorry, the princess is in another castle.")

	}

	return peachToadstool
}
