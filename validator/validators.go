package validator

import (
	"regexp"
)

type (
	emailValidator struct {
		tagRegex   *regexp.Regexp
		emailRegex *regexp.Regexp
	}

	minValidator struct {
		tagRegex *regexp.Regexp
	}

	maxValidator struct {
		tagRegex *regexp.Regexp
	}

	numericValidator struct {
		tagRegex *regexp.Regexp
	}

	requiredValidator struct {
		tagRegex *regexp.Regexp
	}
)

var validators = []validator{
	emailValidator{
		tagRegex:   regexp.MustCompile(`(?i)\bemail\b`),
		emailRegex: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
	},
	minValidator{
		tagRegex: regexp.MustCompile(`min=([^\s]+)`),
	},
	maxValidator{
		tagRegex: regexp.MustCompile(`max=([^\s]+)`),
	},
	numericValidator{
		tagRegex: regexp.MustCompile(`(?i)\bnumeric\b`),
	},
	requiredValidator{
		tagRegex: regexp.MustCompile(`(?i)\brequired\b`),
	},
}
