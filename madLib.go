package main

import (
	"fmt"
	"math/rand"
)

func GetMadLib() string {

	type Responses struct {
		nameField            string
		ageField             string
		occupationField      string
		computingDeviceField string
		bodyPartField        string
		moodWordField        string
		actionWordField      string
	}

	names := []string{
		"Poppy", "Kelly", "Sus", "Cameron", "Lucy", "Miles", "Penny", "Molly", "Princess Peach",
	}

	occupations := []string{
		"accountant", "bus driver", "plumber", "crane operator", "firefighter", "pilot", "teacher", "lawyer",
		"veterinarian", "service advisor", "electrician", "farmer", "superhero", "musician", "artist", "lifeguard",
		"supermodel",
	}

	computingDevices := []string{
		"laptop", "tablet", "chromebook", "PC", "server", "raspberry pi", "Zune music player", "iPod", "Sony Walkman",
		"Samsung refrigerator", "Nintendo 64", "Panasonic VHS player",
	}

	bodyParts := []string{
		"arm", "leg", "torso", "butt", "back", "head", "hand", "foot", "mouth",
	}

	moodWords := []string{
		"happy", "sad", "melancholy", "angry", "ecstatic", "depressed", "bored", "wired", "focused",
	}

	actionWords := []string{
		"exploding", "farting", "eating", "reading", "sleeping", "combusting", "running", "throwing up", "jumping",
		"erupting", "imploding", "hugging", "thinking",
	}

	name := names[rand.Intn(len(names))]
	age := rand.Intn(150)
	ageStr := fmt.Sprintf("%d", age)
	occupation := occupations[rand.Intn(len(occupations))]
	computingDevice := computingDevices[rand.Intn(len(computingDevices))]
	bodyPart := bodyParts[rand.Intn(len(bodyParts))]
	moodWord := moodWords[rand.Intn(len(moodWords))]
	actionWord := actionWords[rand.Intn(len(actionWords))]

	responses := Responses{
		nameField:            name,
		ageField:             ageStr,
		occupationField:      occupation,
		computingDeviceField: computingDevice,
		bodyPartField:        bodyPart,
		moodWordField:        moodWord,
		actionWordField:      actionWord,
	}

	madLib := fmt.Sprintf("%s is a %s-year old %s who has been struggling with a lot of job-related stress."+
		" They decided to try a new application to relieve stress, which runs on a %s to help improve their mood."+
		"\n\nThe application senses the subject's mood through a device they wear on their %s.\n\nWhen the device "+
		"senses that they are %s it responds by %s.", responses.nameField, responses.ageField, responses.occupationField,
		responses.computingDeviceField, responses.bodyPartField, responses.moodWordField, responses.actionWordField)

	return madLib

}
