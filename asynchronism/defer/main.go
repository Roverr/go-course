package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// pinger pings the given URL for every period of time
type pinger struct {
	URL     url.URL
	healthy bool
	client  http.Client
}

// setHealth is for setting the health of the pinger with a pointer
func (p *pinger) setHealth(healthy *bool) {
	if p == nil || healthy == nil {
		return
	}
	p.healthy = *healthy
	p.infoLn(fmt.Sprintf("IsHealthy: %t", p.healthy))
}

// ping is for doing an actual ping to the url
func (p *pinger) ping() {
	var healthy bool
	defer p.setHealth(&healthy)

	// Doing the ping and setting healthy to true, if everything was alright
	resp, err := p.client.Get(p.URL.String())
	if err != nil {
		p.errorLn("Request error", err)
		return
	}
	if resp.StatusCode != 200 {
		p.errorLn("Status not correct", fmt.Errorf("Got %d status", resp.StatusCode))
		return
	}
	if success := isPingSuccessful(resp); !success {
		p.errorLn("Ping tagged as unsuccessful", errors.New("Random just hit ya"))
		return
	}
	healthy = true
}

// errorLn is for logging errors in the pinger
func (p *pinger) errorLn(msg string, err error) {
	fmt.Printf("PINGER ERROR || %s || %s\n", msg, err.Error())
}

func (p *pinger) infoLn(msg string) {
	fmt.Printf("PINGER INFO || %s\n", msg)
}

// schedule is for creating a new while loop which can do the request for every time period
func (p *pinger) schedule(period time.Duration) {
	for {
		p.ping()
		<-time.After(period)
	}
}

// newPinger parses the given url and creates a new pinger
func newPinger(host string) (*pinger, error) {
	URL, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &pinger{URL: *URL, client: http.Client{}}, nil
}

func main() {
	// Initialize and start to schedule the pinger
	pinger, err := newPinger("http://google.com")
	if err != nil {
		log.Fatal(err)
	}
	go pinger.schedule(time.Second * 2)

	doOtherThings()
}
