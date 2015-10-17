package main

import (
    "flag"
    "testing"
)

var (
    endpoint = flag.String("endpoint", "https://benutzerdb.raumzeitlabor.de/BenutzerDB/sshkeys/main", "The BenutzerDB Endpoint")
)

func TestFetchKeys(t *testing.T) {
    _, err := GetKeyring(*endpoint)
    if err != nil {
        t.Fatalf(err.Error())
    }
}
