package hydraconfigurator

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

//ConfigFields to reflect fields
type ConfigFields map[string]reflect.Value

func (f ConfigFields) add(name, value, t string) error {
	switch t {
	case "STRING":
		f[name] = reflect.ValueOf(value)
	case "INTEGER":
		i, err := strconv.Atoi(value)

		if err != nil {
			return err
		}

		f[name] = reflect.ValueOf(i)
	case "FLOAT":
		fl, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return err
		}

		f[name] = reflect.ValueOf(fl)
	case "BOOL":
		b, err := strconv.ParseBool(value)

		if err != nil {
			return err
		}

		f[name] = reflect.ValueOf(b)
	}

	return nil
}

func marshalCustomConfig(v reflect.Value, filename string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic ocurred on marshalling", r)
		}
	}()

	if !v.CanSet() {
		return errors.New("Value passed is not settable")
	}

	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	fields := make(ConfigFields)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println("Processing line", line)

		if strings.Count(line, ";") != 2 {
			continue
		}

		args := strings.Split(line, ";")

		name, value, vType := strings.TrimSpace(args[0]), strings.TrimSpace(args[1]), strings.ToUpper(strings.TrimSpace(args[2]))

		fields.add(name, value, vType)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	vt := v.Type()

	for i := 0; i < vt.NumField(); i++ {
		fieldType := vt.Field(i)
		fieldValue := v.Field(i)

		name := fieldType.Tag.Get("name")

		if name == "" {
			name = fieldType.Name
		}

		if v, ok := fields[name]; ok {
			fieldValue.Set(v)
		}
	}

	return nil
}
