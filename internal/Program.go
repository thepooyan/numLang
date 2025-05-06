package internal

import (
	"fmt"
	"strconv"
)

type Program struct {
	intVariables    map[string]int
	decimalVariables map[string]float64
}

func NewProgram() *Program {
	return &Program{
		intVariables:    make(map[string]int),
		decimalVariables: make(map[string]float64),
	}
}

func (p *Program) addVariable( varName string, varType string, value string,) error {
	switch varType {
	case "int":
		val, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("value is not a valid integer string: %w", err)
		}
		p.intVariables[varName] = val
		return nil
	case "decimal":
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("value is not a valid decimal string: %w", err)
		}
		p.decimalVariables[varName] = val
		return nil
	default:
		return fmt.Errorf("invalid variable type")
	}
}

func (p *Program) retrieveValue(
	varName string,
) (any, string, error) {
	if val, ok := p.intVariables[varName]; ok {
		return val, "int", nil
	}
	if val, ok := p.decimalVariables[varName]; ok {
		return val, "decimal", nil
	}
	return nil, "", fmt.Errorf("variable %s not found", varName)
}


