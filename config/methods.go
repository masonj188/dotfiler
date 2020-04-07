package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
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
	for _, d := range c.Dotfiles {
		for _, h := range d.Hosts {
			if h.Hostname == hostname && !h.Frozen {
				file, err := os.Open(h.Path)
				if err != nil {
					return err
				}
				defer file.Close()
				stat, err := file.Stat()
				if err != nil {
					return err
				}
				switch stat.IsDir() {
				case true:
					err = copyDir(d.Filepath, h.Path)
					if err != nil {
						return err
					}
				case false:
					err = copyFile(d.Filepath, h.Path)
					if err != nil {
						return err
					}
				}
				break
			}
		}
	}
	return nil
}

//Apply replaces all configured files on a host with the file in the current directory
func (c *Config) Apply(hostname string) error {
	for _, d := range c.Dotfiles {
		for _, h := range d.Hosts {
			// Check for which dotfiles to apply; you can cheese this by specifying a hostname in the command line
			if h.Hostname == hostname && !h.Frozen {
				file, err := os.Open(d.Filepath)
				if err != nil {
					return err
				}
				defer file.Close()
				stat, err := file.Stat()
				if err != nil {
					return err
				}
				switch stat.IsDir() {
				case true:
					err = copyDir(h.Path, d.Filepath)
					if err != nil {
						return err
					}
				case false:
					err = copyFile(h.Path, d.Filepath)
					if err != nil {
						return err
					}
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
			if h.Hostname == hostname {
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

func copyDir(dest, src string) error {
	err := copy.Copy(src, dest)
	if err != nil {
		return fmt.Errorf("error copying directory %s: %v", src, err)
	}
	fmt.Printf("Moved %s to %s\n", src, dest)
	return nil
}

func copyFile(dest, src string) error {
	source, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening config dotfile: %v", err)
	}
	defer source.Close()

	// Create the dotfile directory if needed
	if _, err := os.Stat(filepath.Dir(dest)); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(dest), os.ModePerm)
		if err != nil {
			return err
		}
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("Error creating destination dotfile: %v", err)
	}
	defer destFile.Close()

	copied, err := io.Copy(destFile, source)
	if err != nil {
		return err
	}
	if copied < 1 {
		return fmt.Errorf("error copying file")
	}
	fmt.Printf("Moved %s to %s\n", src, dest)

	return nil
}
