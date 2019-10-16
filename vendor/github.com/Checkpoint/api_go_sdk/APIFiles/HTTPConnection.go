package api_go_sdk

import (
	"crypto/sha1"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// HTTP Client wrapper
type Client struct {
	client      *http.Client
	server      string
	sid         string
	fingerprint string
	debugLevel  string
}

// Init and returns new instance of HTTP client wrapper
func CreateClient(server string, sid string, timeout time.Duration) (*Client, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	var netClient = &http.Client{
		Timeout: timeout,
	}
	return &Client{netClient, server, sid, "", ""}, nil
}

// Init and returns new instance of HTTP proxy client
func CreateProxyClient(server string, serverProxy string, sid string, portProxy int, timeout time.Duration) (*Client, error) {
	proxyURL, _ := url.Parse("http://" + serverProxy + ":" + strconv.Itoa(portProxy))
	http.DefaultTransport = &http.Transport{
		Proxy:           http.ProxyURL(proxyURL),
		TLSNextProto:    make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	var netClient = &http.Client{
		Timeout: timeout,
	}

	return &Client{netClient, server, sid, "", ""}, nil
}

// Set debug level for client
func (c *Client) SetDebugLevel(level string) {
	c.debugLevel = level
}

// Returns pointer to HTTP client
func (c *Client) GetClient() *http.Client {
	return c.client
}

// Returns client server
func (c *Client) GetServer() string {
	return c.server
}

// Get fingerprint from server
func getFingerprint(server string, port int) (string, error) {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", server+":"+strconv.Itoa(port), conf)
	if err != nil {
		log.Println(err)
		return "", errors.New("Connection to " + server + " from port " + strconv.Itoa(port) + " failed")
	}
	defer conn.Close()

	var fp string = ""
	state := conn.ConnectionState()
	for _, v := range state.PeerCertificates {
		fp = fp + fmt.Sprintf("%x", sha1.Sum(v.Raw))
	}

	//no error occurred
	return fp, nil

}

// Set fingerprint to client
func (c *Client) setFingerprint(fp string) {
	c.fingerprint = fp
}
