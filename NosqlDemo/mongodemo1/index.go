package main

import (
	"context"
	"mondodemo1/mnops1"
	"time"
)

func main() {
	ctx := context.Background()
	ctx1, _ := context.WithTimeout(ctx, time.Second)
	mnops1.ConnectToDb(ctx1) //for db ops to close within 1 sec
	defer mnops1.CloseConnection(ctx1)
	mnops1.InsertNewUser(ctx1)
}
