package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	var wg sync.WaitGroup
	password := getenv("PASSWORD", "SERVER_UDP")
	port := getenv("PORT", "39000")
	req := getenv("REQUEST", "")

	pc, err := net.ListenPacket("udp4", ":"+port)
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	wg.Add(1)
	go func() {
		buf := make([]byte, 1024)
		fmt.Printf("Server run pass: %s \n", password)
		fmt.Printf("Server run PORT: %s \n", port)
		for {
			n, addr, err := pc.ReadFrom(buf)
			if err != nil {
				panic(err)
			}
			fmt.Printf("[%s] %s\n", addr, buf[:n])
			if bytes.Compare(buf[:n], []byte(password)) == 0 {
				pc.WriteTo([]byte("OK"), addr)
			}
		}
	}()

	if len(req) > 0 {
		wg.Add(1)
		go func() {
			for {
				time.Sleep(time.Second)
				addr, err := net.ResolveUDPAddr("udp4", req)
				if err != nil {
					panic(err)
				}

				if _, err := pc.WriteTo([]byte(password), addr); err != nil {
					panic(err)
				}
			}
		}()
	}
	wg.Wait()

}
