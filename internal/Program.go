package internal

import "fmt"

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

func (p *Program) addVariable(
	varName string,
	varType string,
	value any,
) error {
	switch varType {
	case "int":
		if val, ok := value.(int); ok {
			p.intVariables[varName] = val
			return nil
		} else {
			return fmt.Errorf("value is not an integer")
		}
	case "decimal":
		if val, ok := value.(float64); ok {
			p.decimalVariables[varName] = val
			return nil
		} else {
			return fmt.Errorf("value is not a decimal")
		}
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


