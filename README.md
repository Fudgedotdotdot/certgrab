# certgrab
Dumps certificate information - useful to increase the attack surface of a target by discovering more domains/subdomains.
## Installation
`go get -u github.com/Fudgedotdotdot/certgrab`

## Usage
First column is the requested IP/domain - Second column is the results

```console
# certgrab -h
Usage of certgrab:
  -domain string
        root domain - if not specified, will match everything
  -threads int
        Number of threads (default 150)

# cat domains.txt | certgrab
www.google.com   www.google.com
www.google.com   www.google.com
google.com       google.com
google.com       google.com
google.com       android.com

# cat ips.txt | certgrab
35.160.0.94      www.example.com
35.160.0.11      pepeapi.com
35.160.0.11      pepeapi.com
35.160.0.11      pepeapi.com
35.160.0.11      admin.pepeapi.com
35.160.0.11      api.pepeapi.com
```

## How it works
No enough tools explain this.

This tool starts a TLS handshake with the target, and then dumps the subject name
& the alternate subject names. You can filter the output in order to only get
the domains you're interested in with the -domain flag.
