package utils

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

// DEBUG - Global debug flag
var DEBUG = false

// Reference for PEM file: https://gist.github.com/exAspArk/f738f0771e2675e7f4c3b5d11403efd8

// ClientManager - Hold information to make the connection between this client and wrangler worker possible
type ClientManager struct {
	// Private
	prof Profile
	conn net.Conn

	// Public
	Times   int
	IsHTTPS bool
	URL     string
	Port    string
	Path    string
}

// TCPConnector - Init the TCP connector version to dial to HTTP endpoint
func (c *ClientManager) TCPConnector() ([]byte, error, error) {
	if DEBUG {
		fmt.Println("Initializing TCP Connector Client Manager...")
		fmt.Println("Curr URL + Port: " + c.URL + c.Port)
	}

	// Setup TCP address with the URL given. If error, exit program
	tcpAddr, tcpErr := net.ResolveTCPAddr("tcp", c.URL+c.Port)
	if tcpErr != nil {
		errHelper("Error at resolving tcp address for the endpoint", tcpErr)
	}

	// Setup Connection address with the URL given. If error, exit program
	connection, connectionErr := net.DialTCP("tcp", nil, tcpAddr)
	if connectionErr != nil {
		errHelper("Error in Client Manager while dialing to endpoint", connectionErr)
	}
	c.conn = connection

	// Get access to given URL content
	_, err := c.conn.Write([]byte(fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n", c.Path, c.URL)))
	if err != nil {
		errHelper("Error writing Connection Data", err)
	}
	// Write requested Data
	requestData, requestErr := ioutil.ReadAll(c.conn)
	if requestErr != nil {
		errHelper("Error reading Connection Data", requestErr)
	}
	return requestData, err, requestErr
}

// TLSConnector - Init the TLS connector version to dial to HTTPS endpoint
func (c *ClientManager) TLSConnector() ([]byte, error, error) {
	if DEBUG {
		fmt.Println("Initializing Secured TLS Connector Client Manager...")
		fmt.Println("Curr URL + Port: " + c.URL + c.Port)
	}

	// Config TLS connector
	conf, confErr := tlsConfig()
	if confErr != nil {
		errHelper("Error configurating TLS", confErr)
		os.Exit(2)
	}

	// Setup Connection address with the URL given. If error, exit program
	connection, dialErr := tls.Dial("tcp", c.URL+c.Port, conf)
	if dialErr != nil {
		errHelper("Error while TLS dialing connection", dialErr)
		os.Exit(2)
	}
	c.conn = connection

	// Get access to given URL content
	_, err := c.conn.Write([]byte(fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n", c.Path, c.URL)))
	if err != nil {
		errHelper("Error writing Connection Data", err)
	}
	// Write requested Data
	requestData, requestErr := ioutil.ReadAll(c.conn)
	if requestErr != nil {
		errHelper("Error reading Connection Data", requestErr)
	}
	return requestData, err, requestErr
}

// MakeRequest - Making Request from the client to URL given.
func (c *ClientManager) MakeRequest() {
	var err error
	var requestErr error
	var requestData []byte

	if c.Times <= 0 { // Case invalid Profile Time, only print the response to the console.
		if c.IsHTTPS {
			requestData, err, requestErr = c.TLSConnector()
		} else {
			requestData, err, requestErr = c.TCPConnector()
		}

		if err != nil || requestErr != nil {
			os.Exit(2)
		}

		fmt.Println(string(requestData))
	} else { // Case Profile Time valid, request the URL with the amount given and calculate other info for Profile

		c.prof.InitProfile()
		c.prof.numberRequest = c.Times

		// Do the request ammount
		for i := 0; i < c.Times; i++ {
			// Start timing this request
			start := time.Now()
			if c.IsHTTPS {
				requestData, err, requestErr = c.TLSConnector()
			} else {
				requestData, err, requestErr = c.TCPConnector()
			}

			// If there's an error, increase the failure count and move on to the next request
			if err != nil {
				c.prof.failTime++
				c.prof.HandleError(err)
				continue
			}
			if requestErr != nil {
				c.prof.failTime++
				c.prof.HandleError(err)
				continue
			}
			totalTime := time.Since(start)

			// Print only once!
			if i == 0 {
				fmt.Println(string(requestData))
			}
			c.prof.CalculateRequest(int(totalTime.Milliseconds()), len(requestData))
		}
		// Write requested Data along with Profile info if valid
		c.prof.PrintInfo()
	}
}

// ----------- HELPER -----------

// errHelper - Print out message along with error
func errHelper(message string, err error) {
	fmt.Println(message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
}

// tlsConfig - Configurate the TLS handshake certification file
func tlsConfig() (*tls.Config, error) {
	// NOTE: Attempted to make Handshake certification authorization file to communicate between HTTPS, but failed

	// crt, err := ioutil.ReadFile("./certs/public.crt")
	// if err != nil {
	// 	errHelper("Error while reading certification file", err)
	// 	return nil, err
	// }

	// rootCAs := x509.NewCertPool()
	// rootCAs.AppendCertsFromPEM(crt)

	return &tls.Config{
		// RootCAs: rootCAs,
		InsecureSkipVerify: true,
	}, nil
}
