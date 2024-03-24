package command

import (
	"github.com/xagero/go-helper/helper"
	"regexp"
)

// CmdOption command option, not ordered in console
type CmdOption struct {
	// config
	name        string
	short       string
	input       string
	description string

	// cli input
	exists bool
	value  string
}

// GetOption return CmdOption
func (cmd *Command) GetOption(key string) *CmdOption {
	if opt, ok := cmd.options[key]; ok {
		return opt
	}
	return nil
}

func (cmd *Command) ListOptions() map[string]*CmdOption {
	return cmd.options
}

// SetOptionExist set if option exists in console
func (cmd *Command) SetOptionExist(key string, b bool) {
	if opt, ok := cmd.options[key]; ok {
		opt.exists = b
	} else {
		// @todo fallback InvalidOptionFallback
		panic("Option [" + key + "] not exist")
	}
}

func (cmd *Command) SetOptionExistByShort(short string, b bool) {
	for _, option := range cmd.options {
		if option.short == short {
			option.exists = b
			return
		}
	}

	panic("Short option [" + short + "] not exist")
}

// SetOptionValue set option value if option exists in console
func (cmd *Command) SetOptionValue(key string, value string) {
	if opt, ok := cmd.options[key]; ok {
		if opt.exists {
			opt.value = value
		}
	} else {
		// @todo fallback InvalidOptionFallback
		panic("Option " + key + " not exist")
	}
}

func (cmd *Command) ValidateOptionRequirement() error {
	for _, option := range cmd.options {

		if false == option.Exists() {
			continue // skip non-exists option
		}

		if option.input == OptionValueRequire {
			if helper.IsBlank(option.value) {
				panic("Invalid option [ " + option.name + " ], require value")
			}
		}
	}

	return nil
}

// Name return CmdOption name
func (opt CmdOption) Name() string {
	return opt.name
}

func (opt CmdOption) Input() string {
	return opt.input
}

func (opt CmdOption) Description() string {
	return opt.description
}

// Exists return CmdOption exists
func (opt CmdOption) Exists() bool {
	return opt.exists
}

// Value return CmdOption value
func (opt CmdOption) Value() string {
	return opt.value
}

func (opt CmdOption) Short() string {
	return opt.short
}

func (opt *CmdOption) SetShortSyntax(short string) {
	if opt.input != OptionValueNone {
		panic("Short syntax is for option_value_none")
	}

	if helper.IsBlank(short) || len(short) > 1 {
		panic("Invalid command option short syntax")
	}

	bytes := []byte(short)
	if match, _ := regexp.Match(`[a-z]`, bytes); !match {
		panic("Invalid command option short syntax")
	}

	opt.short = short
}
