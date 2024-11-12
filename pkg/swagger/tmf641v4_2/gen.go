package tmf641v4_1

//go:generate rm -rf server
//go:generate mkdir -p server
//go:generate swagger generate server --quiet --target server --name service-ordering-v4_2 --spec TMF641-ServiceOrdering-v4.2.0.json --exclude-main
