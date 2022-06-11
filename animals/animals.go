/**

The program allow the user to get information about a predefined set of animals.

Animal	|	Food eaten	|	Locomotion method	|	Spoken sound	|
---------------------------------------------------------------------
cow		|	grass		|	walk				|	moo				|
bird	|	worms		|	fly					|	peep			|
snake	|	mice		|	slither				|	hsss			|

Examples:

> cow eat
Outputs: grass
> bird move
Outputs: fly
> snake speak
Outputs: hsss

*/

package animals

type Animal struct {
	food       string
	locomotion string
	noise      string
}

var animals map[string]Animal

func (animal Animal) Eat() string {
	return animal.food
}

func (animal Animal) Move() string {
	return animal.locomotion
}

func (animal Animal) Speak() string {
	return animal.noise
}

func RequestDataAbout(animalName string, propertyName string) string {
	switch propertyName {
	case "eat":
		return animals[animalName].Eat()
	case "move":
		return animals[animalName].Move()
	case "speak":
		return animals[animalName].Speak()
	}
	return "not found"
}

func LoadDataSource() {
	animals = make(map[string]Animal)
	animals["cow"] = Animal{"grass", "walk", "moo"}
	animals["bird"] = Animal{"worms", "fly", "peep"}
	animals["snake"] = Animal{"mice", "slither", "hsss"}
}
