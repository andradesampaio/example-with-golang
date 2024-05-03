package main

import (
	"context"
	"fmt"
	"os"
	"io"
	"net/http"
	"time"
)



func main(){
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//Prepare the request
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		fmt.Println("Error request")
	}

	//Send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error response")
	}
	defer res.Body.Close()
    
	io.Copy(os.Stdout, res.Body)
}