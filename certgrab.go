package main


import (
	"crypto/tls"
	"net"
	"time"
	"fmt"
	"bufio"
	"strings"
	"os"
	"sync"
	"flag"
)


func serverCert(ip string, port string, domain string){
	// setting the timeout
	d := &net.Dialer{
		Timeout: time.Duration(5) * time.Second,
	}

	conn, err := tls.DialWithDialer(d, "tcp", ip+":"+port, &tls.Config{
			InsecureSkipVerify: true,
	})
	if err != nil {
		//fmt.Println(err)
		return
	}
	cert := conn.ConnectionState().PeerCertificates[0].DNSNames

	conn.Close()

	// print the CommonName of the cert
	CommonName := conn.ConnectionState().PeerCertificates[0].Subject.CommonName
	if strings.HasSuffix(CommonName, domain){
			fmt.Println(ip, "\t", strings.TrimLeft(CommonName, "*."))
	}

	// print the Subject Alternate Names
	for _, name := range(cert){
		if strings.HasSuffix(name, domain){
			fmt.Println(ip, "\t", strings.TrimLeft(name, "*."))
		}
	}
	return
}

func main(){
		// parsing flags
		var domain string
		flag.StringVar(&domain, "domain", "", "root domain - if not specified, will match everything")

		var threads int
		flag.IntVar(&threads, "threads", 150, "Number of threads ( default is 150 )")
		flag.Parse()

		ips := make(chan string)
		var wg sync.WaitGroup

		// spawning goroutines
		for i := 0; i < threads; i++{
			wg.Add(1)
			go func(){
				for ip_port := range ips{
					ip := strings.Split(ip_port, ":")[0]
					port := strings.Split(ip_port, ":")[1]
					serverCert(ip, port, domain)
				}
				wg.Done()
			}()
		}

		// scanning stdin and launching goroutines
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
				// one day, we shall be able to run this script fast enough 
				// to add more ports
			ips <- fmt.Sprintf("%s:%s", sc.Text(), "443")

			/*
			ports := []string{"443"}
			for _, port := range ports {
				ips <- fmt.Sprintf("%s:%s", sc.Text(), port)
			}
			*/
		}
		close(ips)
		wg.Wait()
}

