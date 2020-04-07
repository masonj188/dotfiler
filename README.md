# Dotfiler
Dotfiler makes it easy to manage dotfiles across multiple hosts with a single configuration file.

Example of a configuration file:
```yml
dotfiles:
- name: default_i3
  filepath: test/i3/i3
  hosts:
    - name: desktop
      path: /home/user/.config/test/i3/i3
      frozen: false

    - name: blade
      path: /home/user/.config/i3/i3
      frozen: false

- name: polybar/
  filepath: test/polybar/polybar
  hosts:
    - name: desktop
      path: /home/user/.config/test/polybar/polybar
      frozen: true
    - name: blade
      path: /home/user/.config/polybar/
      frozen: false

- name: .vimrc
  filepath: test/vim/.vimrc
  hosts:
    - name: desktop
      path: /home/user/.config/test/vim/.vimrc
      frozen: false
    
    - name: blade
      path: /home/user/.vimrc
      frozen: false
```

A dotfile has the following fields

**name**: A dotfile name has no meaning other than as a way to differentiate it to the user

**filepath**: the path in the current directory where the dotfile resides

**hosts**: A list of hosts to apply this dotfile to

**name**: Hostname of the computer to apply this dotfile to

**path**: Full path and filename of where to put the dotfile on this host

**frozen**: boolean determining whether this host is "frozen" for this dotfile, meaning to skip it for now

## Usage
dotfiler [flags] [command]

Typical usage would be to store all of your dotfiles in a git repository along with a Dotfiler configuration file.  Then applying dotfiles is as easy as cloning the repo and running `dotfiler apply` in the git repo's directory.

## Commands
**apply**

Apply the current configuration file, will replace all configured config files that match the hostname.

**update**

Update the current directories dotfiles with ones from the specific host.

**backup**

Backup current directories dotfiles into the specified directory, must be used with -b flag.

**restore**

Restore a backed up configuration, must be used with the -b flag.


### Flags
**-c *path/to/config.yml***

Specify the path for backup or restore to use (default "./backup/config.yml.bak")

**-h *hostname***

Specify path and name of config file to use (default "./config.yml")


**-b *path/to/backup.bak***

Set hostname to use for this run (default is current machine's hostname)