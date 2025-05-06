package internal

import (
	"fmt"
	"strconv"
)

type Program struct {
	variables    map[string]string
}

func NewProgram() *Program {
	return &Program{
		variables: make(map[string]string),
	}
}

func (p *Program) addVariable( varName string, varType string, value string,) error {
	switch varType {
	case "int":
		_, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("value is not a valid integer string: %w", err)
		}
		p.variables[varName] = value
		return nil
	case "decimal":
		_, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("value is not a valid decimal string: %w", err)
		}
		p.variables[varName] = value
		return nil
	default:
		return fmt.Errorf("invalid variable type")
	}
}

func (p *Program) retrieveValue( varName string,) (string, string, error) {
	if val, ok := p.variables[varName]; ok {
		return val, "int", nil
	}
	return "", "", fmt.Errorf("variable %s not found", varName)
}
