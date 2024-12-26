package tmf641v4_1

//go:generate rm -rf server
//go:generate mkdir -p server -t ./template
//go:generate swagger generate server --quiet --target server --name service-ordering-v4_1 --spec TMF641-ServiceOrdering-v4.1.0.yaml --exclude-main
