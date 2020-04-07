# Dotfiler
Dotfiler makes it easy to manage dotfiles across multiple hosts

## Usage
dotfiler [flags] [command]

Please see `example-config.yml` for an example config file.

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