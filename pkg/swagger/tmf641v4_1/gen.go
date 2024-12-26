package tmf641v4_1

//go:generate rm -rf server
//go:generate mkdir -p server
//go:generate swagger generate server -t ./template --quiet --target server --name service-ordering-v4_1 --spec TMF641-ServiceOrdering-v4.1.0.yaml --exclude-main
