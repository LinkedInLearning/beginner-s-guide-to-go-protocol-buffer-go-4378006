module example.com/customer

go 1.18

require (
	github.com/LinkedInLearning/beginner-s-guide-to-go-Proto-protocol-buffer-go-4378006/go/character v0.0.0-00010101000000-000000000000
	github.com/LinkedInLearning/beginner-s-guide-to-go-Proto-protocol-buffer-go-4378006/go/header v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.3.0
	github.com/mattn/go-sqlite3 v1.14.16
	google.golang.org/grpc v1.54.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace github.com/LinkedInLearning/beginner-s-guide-to-go-Proto-protocol-buffer-go-4378006/go/character => ../../../beginner-s-guide-to-go-Proto-protocol-buffer-go-4378006/go/character

replace github.com/LinkedInLearning/beginner-s-guide-to-go-Proto-protocol-buffer-go-4378006/go/header => ../../../beginner-s-guide-to-go-Proto-protocol-buffer-go-4378006/go/header
