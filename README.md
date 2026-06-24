Go-Pot

A minimal SSH & HTTP honeypot written in Go, intended for observing and logging activity in potentially hostile network environments.
Features

The honeypot logs the following for each connection:

    Unique session ID for each concurrent connection
    Auth username
    Auth password
    Remote address (for ID)
    Command entered

FileSystem mockup

To add extra mock filepaths to your liking, simple edit the "FakeFS" map in fs.go
Notes:

    This is a lightweight simulaton, not a full shell or OS environment
    No commands are executed on the host system
