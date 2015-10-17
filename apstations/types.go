package apinsight

type Station struct {
    ApMac            string      `json:"ap_mac"`
    AssocTime        int         `json:"assoc_time"`
    AuthTime         int         `json:"auth_time"`
    Authorized       bool        `json:"authorized"`
    Bssid            string      `json:"bssid"`
    Bytes            int         `json:"bytes"`
    Bytes_d          int         `json:"bytes-d"`
    Bytes_r          int         `json:"bytes-r"`
    CCQ              int         `json:"ccq"`
    Channel          int         `json:"channel"`
    DHCPEndTime      int         `json:"dhcpend_time"`
    DHCPStartTime    int         `json:"dhcpstart_time"`
    Essid            string      `json:"essid"`
    FirstSeen        int         `json:"first_seen"`
    Hostname         string      `json:"hostname"`
    Idletime         int         `json:"idletime"`
    IP               string      `json:"ip"`
    Is11a            interface{} `json:"is_11a"`
    Is11ac           interface{} `json:"is_11ac"`
    Is11b            interface{} `json:"is_11b"`
    Is11n            bool        `json:"is_11n"`
    IsGuest          bool        `json:"is_guest"`
    LastSeen         int         `json:"last_seen"`
    Mac              string      `json:"mac"`
    Noise            int         `json:"noise"`
    OUI              string      `json:"oui"`
    PowersaveEnabled bool        `json:"powersave_enabled"`
    QOSPolicyApplied bool        `json:"qos_policy_applied"`
    Radio            string      `json:"radio"`
    RSSI             int         `json:"rssi"`
    RxBytes          int         `json:"rx_bytes"`
    RxBytes_d        int         `json:"rx_bytes-d"`
    RxBytes_r        int         `json:"rx_bytes-r"`
    RxCrypts         int         `json:"rx_crypts"`
    RxCrypts_d       int         `json:"rx_crypts-d"`
    RxCrypts_r       int         `json:"rx_crypts-r"`
    RxDropped        int         `json:"rx_dropped"`
    RxDropped_d      int         `json:"rx_dropped-d"`
    RxDropped_r      int         `json:"rx_dropped-r"`
    RxErrors         int         `json:"rx_errors"`
    RxErrors_d       int         `json:"rx_errors-d"`
    RxErrors_r       int         `json:"rx_errors-r"`
    RxFrags          int         `json:"rx_frags"`
    RxFrags_d        int         `json:"rx_frags-d"`
    RxFrags_r        int         `json:"rx_frags-r"`
    RxMcast          int         `json:"rx_mcast"`
    RxPackets        int         `json:"rx_packets"`
    RxPackets_d      int         `json:"rx_packets-d"`
    RxPackets_r      int         `json:"rx_packets-r"`
    RxRate           int         `json:"rx_rate"`
    RxRetries        int         `json:"rx_retries"`
    Signal           int         `json:"signal"`
    SiteID           string      `json:"site_id"`
    State            int         `json:"state"`
    StateHt          bool        `json:"state_ht"`
    StatePwrmgt      bool        `json:"state_pwrmgt"`
    TxBytes          int         `json:"tx_bytes"`
    TxBytes_d        int         `json:"tx_bytes-d"`
    TxBytes_r        int         `json:"tx_bytes-r"`
    TxDropped        int         `json:"tx_dropped"`
    TxDropped_d      int         `json:"tx_dropped-d"`
    TxDropped_r      int         `json:"tx_dropped-r"`
    TxErrors         int         `json:"tx_errors"`
    TxErrors_d       int         `json:"tx_errors-d"`
    TxErrors_r       int         `json:"tx_errors-r"`
    TxPackets        int         `json:"tx_packets"`
    TxPackets_d      int         `json:"tx_packets-d"`
    TxPackets_r      int         `json:"tx_packets-r"`
    TxPower          int         `json:"tx_power"`
    TxRate           int         `json:"tx_rate"`
    TxRetries        int         `json:"tx_retries"`
    TxRetries_d      int         `json:"tx_retries-d"`
    TxRetries_r      int         `json:"tx_retries-r"`
    Uptime           int         `json:"uptime"`
    UserID           string      `json:"user_id"`
}

type StationResponse struct {
    Data    []Station   `json:"data"`
    Meta    struct {
                RC  string  `json:"rc"`
            } `json:"meta"`
}
