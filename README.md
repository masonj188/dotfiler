# Dotfiler
[![Go Report Card](https://goreportcard.com/badge/github.com/masonj188/dotfiler)](https://goreportcard.com/report/github.com/masonj188/dotfiler)

Dotfiler makes it easy to manage dotfiles across multiple hosts with a single configuration file.

Example of a configuration file:
```yml
dotfiles:
- name: default_i3
  filepath: test/i3/i3
  hosts:
    - hostname: desktop
      path: /home/user/.config/test/i3/i3
      frozen: false

    - hostname: blade
      path: /home/user/.config/i3/i3
      frozen: false

- name: polybar/
  filepath: test/polybar/polybar/
  hosts:
    - hostname: desktop
      path: /home/user/.config/test/polybar/
      frozen: true
    - hostname: blade
      path: /home/user/.config/polybar
      frozen: false

- name: .vimrc
  filepath: test/vim/.vimrc
  hosts:
    - hostname: desktop
      path: /home/user/.config/test/vim/.vimrc
      frozen: false
    
    - hostname: blade
      path: /home/user/.vimrc
      frozen: false
```

A dotfile has the following fields

**name**: A dotfile name has no meaning other than as a way to differentiate it to the user

**filepath**: the path in the current directory where the dotfile/config directory resides

**hosts**: A list of hosts to apply this dotfile to

**hostname**: Hostname of the computer to apply this dotfile/directory to

**path**: Full path and filename of where to put the dotfile/directory on this host

**frozen**: boolean determining whether this host is "frozen" for this dotfile, meaning to skip it for now

## Usage
dotfiler [flags] [command]

Typical usage would be to store all of your dotfiles in a git repository along with a Dotfiler configuration file.  Then applying dotfiles is as easy as cloning the repo and running `dotfiler apply` in the git repo's directory.

### Commands
**apply**

Apply the current configuration file, will replace all configured config files that match the hostname.

**update**

Update the current directories dotfiles with ones from the specific host.

**backup** (not implemented yet)

Backup current directories dotfiles into the specified directory, must be used with -b flag.

**restore** (not implemented yet)

Restore a backed up configuration, must be used with the -b flag.


### Flags
**-c *path/to/config.yml***

Specify path and name of config file to use (default "./config.yml")

**-h *hostname***

Set hostname to use for this run (default is current machine's hostname)

**-b *path/to/backup.bak***

Specify the path for backup or restore to use (default "./backup/config.yml.bak")

## Installation
If you have a working Go installation, simply `go get -u github.com/masonj188/dotfiler`

If you don't have a Go installation, precompiled binaries will soon be available for most major operating systems.
