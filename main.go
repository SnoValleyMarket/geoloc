package main

import (
  "context"
  "log"
  "github.com/aws/aws-sdk-go-v2/config"
)

type LocationService struct {
    *client.Client
}

func main(){
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
      log.Fatalf("failed to load configuration, %v", err)
    }

    mySession := session.Must(session.NewSession())
    // Create a LocationService client from just a session.
    svc := locationservice.New(mySession)    
}
