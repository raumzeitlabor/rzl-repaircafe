package apinsight

import (
    "fmt"
    "errors"
    "io/ioutil"
    "net/url"
    "net/http"
    "strings"
    "net/http/cookiejar"
    "encoding/json"
    "crypto/tls"
)

type APClient struct {
    http    *http.Client
    base    string
}

func New(addr string) *APClient {
    ap := &APClient{}
    ap.http = &http.Client{}
    jar, err := cookiejar.New(nil)
    if err != nil {
        panic(err)
    }
    ap.http.Jar = jar
    ap.base = "https://" + addr
    return ap
}

func (ap *APClient) SetInsecureSkipVerify(insec bool) {
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: insec},
    }
    ap.http.Transport = tr
}

func (ap *APClient) Login(user, passwd string) error {
    reqdata := &url.Values{}
    reqdata.Set("login", "Login")
    reqdata.Set("username", user)
    reqdata.Set("password", passwd)

    req, err := http.NewRequest("POST", ap.base + "/login", strings.NewReader(reqdata.Encode()))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    // Looks like there is no easier way to stop following a redirect.
    wantsRedirect := false
    ap.http.CheckRedirect = func(req *http.Request, via []*http.Request) error {
        wantsRedirect = true

        // We need to throw an error to stop the redirect
        return errors.New("allfine")
    }

    resp, err := ap.http.Do(req)

    // If there was an error, check if this is our redirect cancellation
    if err != nil && !wantsRedirect {
        fmt.Println("error is " + err.Error())
        return err
    }

    // Better safe than sorry.
    if err := expectResponse(302, req, resp); err != nil {
        return err
    }

    defer resp.Body.Close()
    return nil
}

func (ap *APClient) GetStations() (interface{}, error) {
    req, err := http.NewRequest("GET", ap.base + "/api/stat/sta", nil)
    if err != nil {
        return nil, err
    }

    if err != nil {
        panic(err)
    }

    resp, err := ap.http.Do(req)
    if err != nil {
        return nil, err
    }

    if err := expectResponse(200, req, resp); err != nil {
        return nil, err
    }

    // We assume a flat key-value response.
    into := &StationResponse{}
    decoder := json.NewDecoder(resp.Body)
    if err := decoder.Decode(into); err != nil {
        return nil, err
    }

    defer resp.Body.Close()
    return into, err
}

func expectResponse(expected int, req *http.Request, resp *http.Response) error {
    if resp.StatusCode != expected {
        msg, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            panic(err)
        }
        return errors.New(fmt.Sprintf("ERROR: %s %s (status=%d, expected=%d):\n%s", req.Method, req.URL.Path, resp.StatusCode, expected, msg))
    }
    return nil
}
