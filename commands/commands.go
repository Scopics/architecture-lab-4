package commands

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/Scopics/architecture-lab-4/engine"
)

var commandsArr = []engine.Command{
	&add{},
	&cat{},
	&print{},
	&delete{},
	&multiply{},
	&palindrom{},
	&printc{},
	&random{},
	&reverse{},
	&sha{},
	&split{},
	&subtract{},
}

func setField(field reflect.Value, str string) error {
	var val interface{}

	switch field.Type().Name() {
	case "int":
		val, _ = strconv.Atoi(str)
	case "float":
		val, _ = strconv.ParseFloat(str, 32)
	default:
		val = str
	}

	if field.Type() != reflect.ValueOf(val).Type() {
		return fmt.Errorf("error: wrong argument")
	}

	field.Set(reflect.ValueOf(val))

	return nil
}

func checkArgs(cmdReflection reflect.Value, args []string) bool {
	for i, v := range args {
		field := cmdReflection.Elem().Field(i)
		err := setField(field, v)
		if err != nil {
			fmt.Println("err")
		}
	}

	return true

}

func Compose(cmdName string, args []string) engine.Command {
	var command engine.Command

	for _, v := range commandsArr {
		commandValue := reflect.ValueOf(v)
		name := commandValue.Elem().Type().Name()
		if cmdName == name {
			if checkArgs(commandValue, args) {
				command = v
				break
			}
		}
	}

	if command == nil {
		command = &print{Arg: "Error: command not found"}
	}

	return command
}
