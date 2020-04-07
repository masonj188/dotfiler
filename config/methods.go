package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//Parse parses a config file
func (c *Config) Parse(path string) error {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configFile, c)
	if err != nil {
		return err
	}

	for _, v := range c.Dotfiles {
		if _, err := os.Stat(v.Filepath); os.IsNotExist(err) {
			return fmt.Errorf("unable to stat %s", v.Filepath)
		}
	}

	return nil
}

//Update searches for each file in a config, and if it matches hostname
//replaces the file in the current directory with the file on that host
func (c *Config) Update(hostname string) error {
	fmt.Println("Update")
	return nil
}

//Apply replaces all configured files on a host with the file in the current directory
func (c *Config) Apply(hostname string) error {
	for _, d := range c.Dotfiles {
		for _, h := range d.Hosts {
			if h.Name == hostname {
				source, err := os.Open(d.Filepath)
				if err != nil {
					return fmt.Errorf("error opening config dotfile: %v", err)
				}
				defer source.Close()

				if _, err := os.Stat(filepath.Dir(h.Path)); os.IsNotExist(err) {
					err = os.MkdirAll(filepath.Dir(h.Path), os.ModePerm)
					if err != nil {
						return err
					}
				}

				dest, err := os.Create(h.Path)
				if err != nil {
					return fmt.Errorf("Error creating destination dotfile: %v", err)
				}
				defer dest.Close()

				copied, err := io.Copy(dest, source)
				if err != nil {
					return err
				}
				if copied < 1 {
					return fmt.Errorf("error copying file")
				}
				break
			}
		}
	}
	return nil
}

//Backup creates a copy of each file for a specified host and places it in path
func (c *Config) Backup(hostname, path string) error {
	fmt.Println("Backup")
	for _, d := range c.Dotfiles {
		for _, h := range d.Hosts {
			if h.Name == hostname {
				// grab h.Path and move it to path
				break
			}
		}
	}
	return nil
}

//Restore replaces all matching files on the host with the file specified in path
func (c *Config) Restore(hostname, path string) error {
	fmt.Println("Restore")
	return nil
}
