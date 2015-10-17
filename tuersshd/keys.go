package main

import (
    "log"
    "sync"
    "net/http"
    "encoding/json"
    "golang.org/x/crypto/ssh"
)

type BenutzerDBKeyHandle struct {
    KeyId     int
    Handle    string
    PublicKey       string    `json:"pubkey"`
    ParsedPublicKey ssh.PublicKey
}

type BenutzerDBKeyring struct {
    sync.RWMutex
    Keys *[]BenutzerDBKeyHandle
}

func (keyring *BenutzerDBKeyring) Refresh(endpoint, keygroup string) error {
    resp, err := http.Get(endpoint + "/sshkeys/" + keygroup)
    if err != nil {
        return err
    }
    keyring.Lock()
    decoder := json.NewDecoder(resp.Body)
    decoder.Decode(&keyring.Keys)
    for i, k := range *keyring.Keys {
        keystr := k.PublicKey + " " + k.Handle
        if k.ParsedPublicKey == nil {
            parsed, _, _, _, err := ssh.ParseAuthorizedKey([]byte(keystr))
            if err != nil {
                log.Println("could not parse key:", keystr)
                continue
            }
            (*keyring.Keys)[i].ParsedPublicKey = parsed
        }
    }
    keyring.Unlock()
    return nil
}
