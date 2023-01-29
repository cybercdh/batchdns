# batchdns

A simple tool which takes an input of domains from stdin and outputs the DNS Status

## Installation

```bash
go install github.com/cybercdh/batchdns@latest
```

## Usage
```bash
cat domains.txt | batchdns
```

Sample output

```bash
www.yahoo.com,NOERROR
uk.yahoo.com,NOERROR
yahoo.com,NOERROR
be.yahoo.com,NOERROR
de.yahoo.com,NOERROR
br.yahoo.com,NOERROR
es.yahoo.com,NOERROR
qc.yahoo.com,NOERROR
mslnp1int.qa.paypal.com,NXDOMAIN
mslnp.qa.paypal.com,NXDOMAIN
```

You can use the concurrency flag to process the input faster if needed

```bash
cat domains.txt | batchdns -c 100
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)