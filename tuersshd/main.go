package main

import (
    "flag"
    "log"
    "net"
    "bytes"
    "errors"
    "time"
    "io/ioutil"
    "golang.org/x/crypto/ssh"
)

var (
    refresh  = flag.Int("refresh", 300, "Refresh interval in seconds (pubkey synchronization)")
    endpoint = flag.String("endpoint", "https://benutzerdb.raumzeitlabor.de/BenutzerDB", "The BenutzerDB endpoint")
    privkey  = flag.String("privkey", "id_rsa", "The private keyfile for the server")
    bindAddr = flag.String("bind", "0.0.0.0:2323", "The address to listen on for SSH")
    keyGroup = flag.String("keygroup", "main", "The group from which to fetch keys")
    socket   = flag.String("socket", "/tmp/pinpad-ctrl.sock", "Path to pinpad ctrl socket")
)

var keyring *BenutzerDBKeyring

func main() {
    flag.Parse()

    keyring = &BenutzerDBKeyring{}
    go refreshKeyring()

    listener, err := net.Listen("tcp", *bindAddr)
    if err != nil {
        panic("failed to listen for connection")
    }

    conf := getServerConfig()

    for {
        conn, err := listener.Accept()
        if err != nil {
            panic("failed to accept incoming connection")
        }
        go service(conn, conf)
    }
}

func refreshKeyring() {
    for {
        log.Println("refreshing keys")
        err := keyring.Refresh(*endpoint, *keyGroup)
        if err != nil {
            log.Println("could not refresh keyring:", err)
        } else {
            log.Printf("keys refreshed (got %d)", len(*keyring.Keys))
        }
        time.Sleep(time.Duration(*refresh) * time.Second)
    }
}

func getServerConfig() *ssh.ServerConfig {
    config := &ssh.ServerConfig{}
    privateBytes, err := ioutil.ReadFile(*privkey)
    if err != nil {
        panic("Failed to load private key")
    }

    private, err := ssh.ParsePrivateKey(privateBytes)
    if err != nil {
        panic("Failed to parse private key")
    }

    config.AddHostKey(private)
    config.PublicKeyCallback = pubkeyAuthCallback

    return config
}

func pubkeyAuthCallback(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
    if keyring.Keys == nil {
        log.Println("rejecting authentication due to missing keyring")
        return nil, errors.New("no keyring available")
    }

    var keyFound *BenutzerDBKeyHandle
    for _, k := range *keyring.Keys {
        if k.ParsedPublicKey == nil {
            continue
        } else if bytes.Compare(key.Marshal(), k.ParsedPublicKey.Marshal()) == 0 {
            keyFound = &k
            break
        }
    }

    if keyFound == nil {
        log.Println("could not authenticate", conn.RemoteAddr().String(), " no key found")
        return nil, errors.New("invalid authentication")
    }

    log.Println("accepted key for user:", keyFound.Handle)
    return &ssh.Permissions{Extensions: map[string]string{"user_id": keyFound.Handle}}, nil
}

func service(conn net.Conn, config *ssh.ServerConfig) {
    _, chans, reqs, err := ssh.NewServerConn(conn, config)
    if err != nil {
        log.Println("could not complete handshake for", conn.RemoteAddr().String(), ":", err)
        return
    }

    go ssh.DiscardRequests(reqs)

    for newChannel := range chans {
        go handleChanReq(newChannel)
    }

    log.Println("finished servicing")
}

func handleChanReq(chanReq ssh.NewChannel) {
    if chanReq.ChannelType() != "session" {
        chanReq.Reject(ssh.UnknownChannelType, "unknown channel type")
        return
    }

    channel, requests, err := chanReq.Accept()
    if err != nil {
        return
    }

    for req := range requests {
        if req.Type == "exec" {
            handleExec(channel, req)
            if req.WantReply {
                req.Reply(true, nil)
            }
            break
        }
    }

    /* 
       todo: return exit status:

       byte      SSH_MSG_CHANNEL_REQUEST
       uint32    recipient channel
       string    "exit-status"
       boolean   FALSE
       uint32    exit_status
    */
    channel.Close()
}

func handleExec(ch ssh.Channel, req *ssh.Request) {
    cmd := string(req.Payload[4:])
    log.Println("received cmd", cmd)

    var msg string
    switch cmd {
    case "open", "close":
        err := sendCommand(cmd)
        if err != nil {
            log.Println("could not write to ctrl sock:", err)
            return
        }
        log.Printf("sent command '%s' to ctrl", cmd)
        msg = "ok"
    default:
        msg = "invalid command " + cmd
    }

    ch.Write([]byte(msg + "\r\n"))
}

func sendCommand(cmd string) error {
    c, err := net.Dial("unix", *socket)
    if err != nil {
        return err
    }
    _, err = c.Write([]byte(cmd + "\n"))
    if err != nil {
        return err
    }
    defer c.Close()
    return nil
}
