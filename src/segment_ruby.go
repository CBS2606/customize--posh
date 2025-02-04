package main

type ruby struct {
	language *language
}

func (r *ruby) string() string {
	version := r.language.string()
	// asdf default non-set version
	if version == "______" {
		return ""
	}
	return version
}

func (r *ruby) init(props properties, env environmentInfo) {
	r.language = &language{
		env:        env,
		props:      props,
		extensions: []string{"*.rb", "Rakefile", "Gemfile"},
		commands: []*cmd{
			{
				executable: "rbenv",
				args:       []string{"version-name"},
				regex:      `(?P<version>.+)`,
			},
			{
				executable: "rvm-prompt",
				args:       []string{"i", "v", "g"},
				regex:      `(?P<version>.+)`,
			},
			{
				executable: "chruby",
				args:       []string(nil),
				regex:      `\* (?P<version>.+)\n`,
			},
			{
				executable: "asdf",
				args:       []string{"current", "ruby"},
				regex:      `ruby\s+(?P<version>[^\s]+)\s+`,
			},
			{
				executable: "ruby",
				args:       []string{"--version"},
				regex:      `ruby\s+(?P<version>[^\s]+)\s+`,
			},
		},
	}
}

func (r *ruby) enabled() bool {
	return r.language.enabled()
}
