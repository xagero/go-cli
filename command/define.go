package command

const (
	ArgumentOptional = "argument_optional"
	ArgumentRequired = "argument_required"
)

const (
	OptionValueNone     = "option_value_none"
	OptionValueOptional = "option_value_optional"
	OptionValueRequire  = "option_value_required"
)

type Callback func() error

// Command simple command structure
type Command struct {
	// config
	name        string
	description string

	// callback fn
	callback       Callback
	callbackBefore Callback
	callbackAfter  Callback

	// cli parameters
	arguments map[string]*CmdArgument
	options   map[string]*CmdOption
}

// CmdArgument command argument, always ordered in console
type CmdArgument struct {
	// config
	name        string
	position    int
	input       string
	description string

	// cli input
	value string
}

// CmdOption command option, not ordered in console
type CmdOption struct {
	// config
	name        string
	input       string
	description string

	// cli input
	exists bool
	value  string
}
