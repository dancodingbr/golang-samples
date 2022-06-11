package animals

import (
	"fmt"
	"strings"
	"testing"
)

func TestRequestOfAnimalAndEatProperty(t *testing.T) {
	cow := Animal{"grass", "walk", "moo"}
	got := cow.Eat()
	wanted := "grass"

	if strings.Compare(got, wanted) != 0 {
		t.Errorf("got: %s; wanted: %s\n", got, wanted)
	} else {
		fmt.Printf("got: %s; wanted: %s\n", got, wanted)
	}
}

func TestRequestOfAnimalAndMoveProperty(t *testing.T) {
	bird := Animal{"worms", "fly", "peep"}
	got := bird.Move()
	wanted := "fly"

	if strings.Compare(got, wanted) != 0 {
		t.Errorf("got: %s; wanted: %s\n", got, wanted)
	} else {
		fmt.Printf("got: %s; wanted: %s\n", got, wanted)
	}
}

func TestRequestOfAnimalAndSpeakProperty(t *testing.T) {
	snake := Animal{"worms", "fly", "hsss"}
	got := snake.Speak()
	wanted := "hsss"

	if strings.Compare(got, wanted) != 0 {
		t.Errorf("got: %s; wanted: %s\n", got, wanted)
	} else {
		fmt.Printf("got: %s; wanted: %s\n", got, wanted)
	}
}

func TestLoadDataSource(t *testing.T) {
	LoadDataSource()
	got := len(animals)
	wanted := 3

	if got != wanted {
		t.Errorf("got: %d; wanted: %d\n", got, wanted)
	} else {
		fmt.Printf("got: %d; wanted: %d\n", got, wanted)
	}
}

func TestTheRequestedDataAboutEat(t *testing.T) {
	animalName := "cow"
	propertyName := "eat"
	got := RequestDataAbout(animalName, propertyName)
	wanted := "grass"

	if strings.Compare(got, wanted) != 0 {
		t.Errorf("got: %s; wanted: %s\n", got, wanted)
	} else {
		fmt.Printf("got: %s; wanted: %s\n", got, wanted)
	}

	animalName = "bird"
	propertyName = "eat"
	got = RequestDataAbout(animalName, propertyName)
	wanted = "worms"

	if strings.Compare(got, wanted) != 0 {
		t.Errorf("got: %s; wanted: %s\n", got, wanted)
	} else {
		fmt.Printf("got: %s; wanted: %s\n", got, wanted)
	}

	animalName = "snake"
	propertyName = "eat"
	got = RequestDataAbout(animalName, propertyName)
	wanted = "mice"

	if strings.Compare(got, wanted) != 0 {
		t.Errorf("got: %s; wanted: %s\n", got, wanted)
	} else {
		fmt.Printf("got: %s; wanted: %s\n", got, wanted)
	}
}

func TestTheRequestedDataAboutMove(t *testing.T) {
	animalName := "cow"
	propertyName := "move"
	got := RequestDataAbout(animalName, propertyName)
	wanted := "walk"

	if strings.Compare(got, wanted) != 0 {
		t.Errorf("got: %s; wanted: %s\n", got, wanted)
	} else {
		fmt.Printf("got: %s; wanted: %s\n", got, wanted)
	}

	animalName = "bird"
	propertyName = "move"
	got = RequestDataAbout(animalName, propertyName)
	wanted = "fly"

	if strings.Compare(got, wanted) != 0 {
		t.Errorf("got: %s; wanted: %s\n", got, wanted)
	} else {
		fmt.Printf("got: %s; wanted: %s\n", got, wanted)
	}

	animalName = "snake"
	propertyName = "move"
	got = RequestDataAbout(animalName, propertyName)
	wanted = "slither"

	if strings.Compare(got, wanted) != 0 {
		t.Errorf("got: %s; wanted: %s\n", got, wanted)
	} else {
		fmt.Printf("got: %s; wanted: %s\n", got, wanted)
	}
}

func TestTheRequestedDataAboutSpeak(t *testing.T) {
	animalName := "cow"
	propertyName := "speak"
	got := RequestDataAbout(animalName, propertyName)
	wanted := "moo"

	if strings.Compare(got, wanted) != 0 {
		t.Errorf("got: %s; wanted: %s\n", got, wanted)
	} else {
		fmt.Printf("got: %s; wanted: %s\n", got, wanted)
	}

	animalName = "bird"
	propertyName = "speak"
	got = RequestDataAbout(animalName, propertyName)
	wanted = "peep"

	if strings.Compare(got, wanted) != 0 {
		t.Errorf("got: %s; wanted: %s\n", got, wanted)
	} else {
		fmt.Printf("got: %s; wanted: %s\n", got, wanted)
	}

	animalName = "snake"
	propertyName = "speak"
	got = RequestDataAbout(animalName, propertyName)
	wanted = "hsss"

	if strings.Compare(got, wanted) != 0 {
		t.Errorf("got: %s; wanted: %s\n", got, wanted)
	} else {
		fmt.Printf("got: %s; wanted: %s\n", got, wanted)
	}
}
