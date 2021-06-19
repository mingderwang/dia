package main

import (
    "os"
    "context"
    "fmt"
    "log"
    "github.com/ethereum/go-ethereum/ethclient"
)
func Connect(host string) (*ethclient.Client, error) {
    conn, err := ethclient.Dial(host)
    if err != nil {
        return nil, err
    }
    return conn, nil
}

func main() {

    infuraProjectID := os.Getenv("INFURA_PROJECT_ID")
    if (0!=len(infuraProjectID)) {

    client, err:= Connect("https://mainnet.infura.io/v3/" + infuraProjectID)
    if err != nil {
        log.Fatal(err)
        panic(err.Error())
    }
    header, err := client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("🦄 current hight on mainnet is: ", header.Number.String()) // 12661894

    } else {
      fmt.Println("INFURA_PROJECT_ID:", os.Getenv("INFURA_PROJECT_ID"))
      fmt.Println("👻 need to define an env. variable called INFURA_PROJECT_ID, 🔥 you can get one from https://infura.io/")
    }
}

