# Project Kata
Coding exercises for Golang.
Process the STIX 2.1 examples located here: 
`https://oasis-open.github.io/cti-documentation/stix/examples`

## Activites

1. read data from files into byte stream
2. unmarshal into a bundle object & child objects
3. marshal it back into a json string
4. print the content on console
5. visually inspect the result.

Do these activities for each of the examples STIX provided on the above website.
Uses the STIX 2.1 spec to ensure compliant.


## To run this program

```
 * cd /home/student/go/src/kata
 * 
 * export GOPATH=/home/student/go
 * export GOBIN=/home/student/go/bin
 * export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
 * 
 * go install
 * kata {stix_file_name}
```
