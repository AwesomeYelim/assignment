package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080") // Dial : 전화걸다
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	go func(c net.Conn) {
		send := []string{"피카츄", "라이츄", "파이리", "꼬부기", "버터풀", "야도란"}
		for i := range send {
			_, err = c.Write([]byte(send[i]))

			if err != nil {
				fmt.Println("Failed to write data : ", err)
				break
			}
			defer conn.Close()

			i++
			time.Sleep(3 * time.Second)

		}
	}(conn)

	// 내부 동기식 처리로 전송(쓰기)되고 나서 연결이 끊어지는 현상이 생겨 따로 함수로 비동기식 처리
	go func(c net.Conn) {
		recv := make([]byte, 4096)

		for {
			n, err := c.Read(recv)
			if err != nil {
				os.Exit(0) // 인터럽트 신호로 터미널 닫음
				//fmt.Println(err)
				return
			}
			fmt.Println("client : ", string(recv[:n]))
		}
	}(conn)

	fmt.Scanln() // 고루틴을 이어나가기 위해 fmt.Scanln() 넣어줌
}
