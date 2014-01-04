virtualdocker
=============

**Warning:** Work in progress.

Runs [boot2docker](https://github.com/steeve/boot2docker) inside VirtualBox and configures Docker to use it.

    $ virtualdocker up
    $ docker version

It is intended for use on platforms which do not support the Docker daemon such as Mac OS X.

Commands
--------

### destroy

Removes the VM.

### halt

Stops the VM.

### ssh

Logs into the VM via SSH.

### status

Prints whether the VM is running or not.

### up

Starts a VM with boot2docker running on it.

