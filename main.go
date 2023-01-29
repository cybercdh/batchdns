package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/lixiangzhong/dnsutil"
	"github.com/miekg/dns"
)

var concurrency int
var domains = make(chan string, 200)

func main() {
	flag.IntVar(&concurrency, "c", 20, "set the concurrency level")

	flag.Parse()

	// perform the dig concurrently
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for domain := range domains {
				msg, err := digDomain(domain)
				if err != nil {
					log.Printf("Error processing %s", domain)
					continue
				}
				status := dns.RcodeToString[msg.MsgHdr.Rcode]
				fmt.Printf("%s,%s\n", domain, status)
			}
		}()
	}

	// get user input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		domains <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// tidy up
	close(domains)
	wg.Wait()
}

func digDomain(domain string) (*dns.Msg, error) {
	var dig dnsutil.Dig
	dig.SetDNS("1.1.1.1")
	msg, err := dig.GetMsg(dns.TypeA, domain)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return msg, nil
}
