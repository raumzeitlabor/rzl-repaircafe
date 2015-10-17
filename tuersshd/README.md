A standalone daemon that starts an SSH server to which users can authenticate
using their BenutzerDB ssh keys. Any user can be used.

### Usage

Starting the daemon is self-explanatory, make sure it can write to the pinpad
controller socket file.

    % ./tuersshd --help
    Usage of ./tuersshd:
      -bind="0.0.0.0:2323": The address to listen on for SSH
      -endpoint="https://benutzerdb.raumzeitlabor.de/BenutzerDB": The BenutzerDB endpoint
      -keygroup="main": The group from which to fetch keys
      -privkey="id_rsa": The private keyfile for the server
      -refresh=300: Refresh interval in seconds (pubkey synchronization)
      -socket="/tmp/pinpad-ctrl.sock": Path to pinpad ctrl socket

Only two commands are supported.

    ssh pinpad -p 2323 open
    ssh pinpad -p 2323 close

Any user can be used.

### Bugs

Currently the exit status is not returned properly.

© Simon Elsbrock 2015
