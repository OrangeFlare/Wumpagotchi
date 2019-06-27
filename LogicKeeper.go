package main

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// LogicKeeper performs numerous checks on a Wumpus to make sure their stats
// are within range
//
// Requires 2 Arguments
// UserWumpus Wumpus should be the Wumpus you want to check and correct
//
// Returns CorrectedWumpus Wumpus which is the original Wumpus with values
// that are within range
func LogicKeeper(UserWumpus Wumpus) (CorrectedWumpus Wumpus) {
	if UserWumpus.Age > 14 {
		CorrectedWumpus.Age = 14
	} else if UserWumpus.Age < 0 {
		CorrectedWumpus.Age = 0
	} else {
		CorrectedWumpus.Age = UserWumpus.Age
	}

	if UserWumpus.Health > 10 {
		CorrectedWumpus.Health = 10
	} else if UserWumpus.Health < 0 {
		CorrectedWumpus.Health = 0
		CorrectedWumpus.Left = true
	} else {
		CorrectedWumpus.Health = UserWumpus.Health
	}

	if UserWumpus.Energy > 10 {
		CorrectedWumpus.Energy = 10
	} else if UserWumpus.Energy < 0 {
		CorrectedWumpus.Energy = 0
	} else {
		CorrectedWumpus.Energy = UserWumpus.Energy
	}

	if UserWumpus.Happiness > 10 {
		CorrectedWumpus.Happiness = 10
	} else if UserWumpus.Happiness < 0 {
		CorrectedWumpus.Happiness = 0
	} else {
		CorrectedWumpus.Happiness = UserWumpus.Happiness
	}

	if UserWumpus.Hunger > 10 {
		CorrectedWumpus.Hunger = 10
	} else if UserWumpus.Hunger < 0 {
		CorrectedWumpus.Hunger = 0
	} else {
		CorrectedWumpus.Hunger = UserWumpus.Hunger
	}

	if UserWumpus.Credits < 0 {
		CorrectedWumpus.Credits = 0
	} else {
		CorrectedWumpus.Credits = UserWumpus.Credits
	}

	if UserWumpus.Age >= 14 || UserWumpus.Health <= 0 {
		CorrectedWumpus.Left = true
	} else {
		CorrectedWumpus.Left = UserWumpus.Left
	}

	CorrectedWumpus.Color = UserWumpus.Color
	CorrectedWumpus.Sick = UserWumpus.Sick
	CorrectedWumpus.Sleeping = UserWumpus.Sleeping
	CorrectedWumpus.Name = UserWumpus.Name
	return CorrectedWumpus
}

// LeftCheck checks if the Wumpus has left
// If the user's wumpus has left, it'll tell them to run the w.view command
// This tells the user different things depending upon if the wumpus is in good conditions
// if the Wumpus hasn't left yet this will return false
func LeftCheck(UserWumpus Wumpus, session *discordgo.Session, event *discordgo.MessageCreate) (Left bool) {
	if UserWumpus.Left && UserWumpus.Health <= 0 {
		go sendMessage(session, event, event.ChannelID, UserWumpus.Name+"Appears to be in critical condition\nPlease view your Wumpus (w.view)")
		return true
	} else if UserWumpus.Age >= 14 {
		UserWumpus.Left = true
		UserWumpus.Age = 14
		UpdateWumpus(event.Author.ID, UserWumpus)
		go sendMessage(session, event, event.ChannelID, UserWumpus.Name+" has something important to tell you!\nPlease view your Wumpus (w.view)")
		return trueb
	} else if UserWumpus.Age > 9 && UserWumpus.Left {
		go sendMessage(session, event, event.ChannelID, UserWumpus.Name+" has something important to tell you!\nPlease view your Wumpus (w.view)")
		return true
	} else if UserWumpus.Age > 4 && UserWumpus.Age < 10 && UserWumpus.Left {
		go sendMessage(session, event, event.ChannelID, UserWumpus.Name+" wants to talk\nPlease view your Wumpus (w.view)")
		return true
	} else if UserWumpus.Left {
		go sendMessage(session, event, event.ChannelID, "You can't seem to find "+UserWumpus.Name+" anywhere...\nPlease view your Wumpus (w.view)")
		return true
	}
	//Wumpus hasn't left yet :D
	return false
}

// EnergyCheck checks if the Wumpus has enough energy
// If the user's wumpus doesn't have enough energy it will tell the user their wumpus is too tired and return true
// If the user's wumpus has enough energy it will return false
func EnergyCheck(UserWumpus Wumpus, requiredEnergy int, session *discordgo.Session, event *discordgo.MessageCreate) (noPass bool) {
	if UserWumpus.Energy < requiredEnergy {
		go sendMessage(session, event, event.ChannelID, UserWumpus.Name+" is too tired!")
		return true
	}
	return false
}

// SleepCheck checks if the Wumpus is sleeping and if they can be woken up
// If the user's wumpus has more than 0 energy it will wake them up from sleep and return a wumpus object that reflects that they are awake
// Additionally this will also alert the user their wumpus was woken up
// If the user's wumpus has 0 energy it will tell the user that they are sleeping and return an unmodified Wumpus
// If theu user's wumpus isn't sleeping it will just return an umodified Wumpus
func SleepCheck(UserWumpus Wumpus, session *discordgo.Session, event *discordgo.MessageCreate) (UpdatedWumpus Wumpus) {
	if UserWumpus.Sleeping {
		if UserWumpus.Energy > 0 {
			UserWumpus.Sleeping = false
			go sendMessage(session, event, event.ChannelID, UserWumpus.Name+" has woken from sleep!")
			return UserWumpus
		}
		go sendMessage(session, event, event.ChannelID, UserWumpus.Name+" is sleeping!")
		return UserWumpus
	}
	return UserWumpus
}

// CreditCheck checks if the user has enough credits
// If the user has enough credits it return false
// If the user doesn't have enough credits, it will alert them they need however many credits they need and then return true
func CreditCheck(UserWumpus Wumpus, creditsRequired int, session *discordgo.Session, event *discordgo.MessageCreate) (noPass bool) {
	if UserWumpus.Credits < creditsRequired {
		go sendMessage(session, event, event.ChannelID, "You need "+strconv.Itoa(creditsRequired)+"Ꞡ!")
		return true
	}
	return false
}
