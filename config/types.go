package config

//Host represents a host within a specific Dotfile
type Host struct {
	Name   string
	Path   string
	Frozen bool
}

//Dotfile is a specific Dotfile configuration
type Dotfile struct {
	Name     string
	Filepath string
	Hosts    []Host
}

//Config represents a config file, which is a collection of Dotfiles
type Config struct {
	Dotfiles []Dotfile
}
