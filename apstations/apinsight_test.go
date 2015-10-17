package apinsight

import (
    "flag"
    "testing"
)

var (
    addr    = flag.String("ctrl", "unifi.vm.rzl:8443", "address of the controller")
    user    = flag.String("user", "noc", "username for login")
    passwd  = flag.String("passwd", "", "password for login")
    insecure = flag.Bool("insecure", false, "accept any certificate")
)

func TestLogin(t *testing.T) {
    c := New(*addr)
    c.SetInsecureSkipVerify(*insecure)
    if err := c.Login(*user, *passwd); err != nil {
        t.Fatalf(err.Error())
    }
}

func TestGetStations(t *testing.T) {
    c := New(*addr)
    c.SetInsecureSkipVerify(*insecure)
    if err := c.Login(*user, *passwd); err != nil {
        t.Fatalf(err.Error())
    }
    _, err := c.GetStations()
    if err != nil {
        t.Fatalf(err.Error())
    }
}
