package commands

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

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
	&sha{},
}

func setField(field reflect.Value, str string) error {
	var val interface{}
	var err error

	switch field.Type().Name() {
	case "int":
		val, err = strconv.Atoi(str)
	case "float":
		val, err = strconv.ParseFloat(str, 32)
	default:
		val = str
	}

	if err != nil {
		return err
	}

	if field.Type() != reflect.ValueOf(val).Type() {
		return fmt.Errorf("error: wrong argument")
	}

	field.Set(reflect.ValueOf(val))

	return nil
}

func setArgs(cmdReflection reflect.Value, args []string) error {
	cmdReflectionElem := cmdReflection.Elem()
	if cmdReflectionElem.NumField() != len(args) {
		return fmt.Errorf("error: wrong number of arguments")
	}

	for i, v := range args {
		field := cmdReflectionElem.Field(i)
		err := setField(field, v)
		if err != nil {
			return err
		}
	}

	return nil

}

func clone(oldObj interface{}) interface{} {
	newObj := reflect.New(reflect.TypeOf(oldObj).Elem())
	oldVal := reflect.ValueOf(oldObj).Elem()
	newVal := newObj.Elem()
	for i := 0; i < oldVal.NumField(); i++ {
		newValField := newVal.Field(i)
		if newValField.CanSet() {
			newValField.Set(oldVal.Field(i))
		}
	}
	return newObj.Interface()
}

func compose(cmdName string, args []string) engine.Command {
	var command engine.Command

	for _, v := range commandsArr {
		commandValue := reflect.ValueOf(v)
		name := commandValue.Elem().Type().Name()
		if cmdName == name {
			err := setArgs(commandValue, args)
			if err == nil {
				command = clone(v).(engine.Command)
			} else {
				command = &print{Arg: fmt.Sprintf("SYNTAX ERROR: %s", err)}
			}
			break
		}
	}

	if command == nil {
		command = &print{Arg: "Error: command not found"}
	}

	return command
}

func Parse(line string) engine.Command {
	splittedLine := strings.Fields(line)

	command, params := splittedLine[0], splittedLine[1:]
	return compose(command, params)
}
