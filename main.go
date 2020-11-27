package kata

// Main expects file name as the first argument (location 1).
// By default Args[0] contains the name of the executable.
//func main() {
//	fmt.Println("bundle: main{}")
//	if len(os.Args) < 2 {
//		fmt.Println("*** error: no input file ***")
//		os.Exit(-1)
//	}
//	data := readAll(os.Args[1])
//	bundle := unmarshal(data)
//	jsonData := marshalBundle(bundle)
//	fmt.Printf("Bundle:%v\n", jsonData)
//}

/**
 * Need:
 * go get github.com/neocortical/flexjson
 *
 * To run this:
 *
 * cd /home/student/go/src/kata
 * 
 * export GOPATH=/home/student/go
 * export GOBIN=/home/student/go/bin
 * export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
 * 
 * go install
 * kata {stix_file_name}
 */