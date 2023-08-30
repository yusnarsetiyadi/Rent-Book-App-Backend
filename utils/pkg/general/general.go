package general

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GenerateUUID() string {
	newUUID := uuid.New()
	uuidString := newUUID.String()
	return uuidString
}

func GeneratePassword(passwordLength, minSpecialChar, minNum, minUpperCase, minLowerCase int) string {
	var password strings.Builder
	var lowerCharSet string = "abcdedfghijklmnopqrstuvwxyz"
	var upperCharSet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var specialCharSet string = "!@#$%&*"
	var numberSet string = "0123456789"
	var allCharSet string = lowerCharSet + upperCharSet + specialCharSet + numberSet

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	//Set lowercase
	for i := 0; i < minLowerCase; i++ {
		random := rand.Intn(len(lowerCharSet))
		password.WriteString(string(lowerCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase - minLowerCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	PASSWORD := string(inRune)
	return PASSWORD
}

func DateTodayLocal() *time.Time {
	now := time.Now().UTC().Add(time.Hour * 7)
	return &now
}

func ParseStringToTime(date string) time.Time {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		logrus.Error("Error Parse String To Time")
	}
	return parsedTime.UTC()
}
