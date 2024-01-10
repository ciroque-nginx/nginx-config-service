module ciroque/go-http-server

go 1.21

require (
	github.com/prometheus/client_golang v1.17.0
	github.com/sirupsen/logrus v1.9.3-0.20230531171720-7165f5e779a5
)
require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace (
	github.com/nginxinc/kubernetes-nginx-ingress/internal/config => ./internal/config
	github.com/nginxinc/kubernetes-nginx-ingress/internal/translation => ./internal/translation
	github.com/nginxinc/kubernetes-nginx-ingress/internal/translation/nginxplus => ./internal/translation/nginxplus
)
