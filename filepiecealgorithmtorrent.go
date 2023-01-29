package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

const pieceSize = 256 * 1024 // 256 KB

func divideIntoPieces(file *os.File) ([][20]byte, error) {
	reader := bufio.NewReader(file)
	pieces := [][20]byte{}
	var piece []byte
	for {
		chunk := make([]byte, pieceSize)
		n, err := reader.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		chunk = chunk[:n]
		if len(piece)+n > pieceSize {
			pieces = append(pieces, sha1.Sum(piece))
			piece = chunk
		} else {
			piece = append(piece, chunk...)
		}
	}
	if len(piece) > 0 {
		pieces = append(pieces, sha1.Sum(piece))
	}
	return pieces, nil
}

func main() {
	file, err := os.Open("sample.torrent")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	pieces, err := divideIntoPieces(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, piece := range pieces {
		fmt.Printf("Piece %d: %x\n", i, piece)
	}
}
