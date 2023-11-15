package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	//"strconv"
	"strings"

	// this has to be the same as the go.mod module,
	// followed by the path to the folder the proto file is in.
	gRPC "github.com/hannaStokes/handin/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)