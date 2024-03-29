package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func analyzeVRAMDump() {
	binaryDumpFile, err := os.Open("dumps/vba_vram_dump")
	if err != nil {
		panic(err)
	}
	defer binaryDumpFile.Close()

	byteDump := make([]byte, 0)

	scanner := bufio.NewScanner(binaryDumpFile)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		byteDump = append(byteDump, bytes...)
		bytesRead := len(bytes)
		fmt.Println("Bytes Read:", bytesRead)
	}

	counter := 0
	for i := 0; i < len(byteDump)-16; i += 16 {
		chunk := byteDump[i : i+16]
		fmt.Printf("0x%.4x: %s\n", counter, formatByteLine16(chunk))
		counter += 16
	}
	fmt.Println("sizeof slice:", len(byteDump))

}

func SlicesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// returns the specific chunk of where a specific sprite may be
func fetchProbableSpriteTile() ([]byte, error) {
	romContent, err := os.ReadFile("KouchuOujaMushiking_J.gba")
	if err != nil {
		return nil, err
	}

	target := []byte{0x3e, 0xdb, 0x40, 0xc0, 0x0e, 0xe3, 0x0e, 0xa8, 0x40, 0x44, 0x55, 0x00, 0x44, 0x34, 0x00, 0x24}

	// parse bytes 16 at a time
	for i := 0; i < len(romContent)-16; i++ {
		compareString := romContent[i : i+16]
		if SlicesEqual(compareString, target) {
			fmt.Printf("Found bytes %s at index %d\n", target, i)

			return romContent[i : i+16*15], nil
		}
	}

	return nil, errors.New("Not Found")
}

func formatByteLine16(b []byte) string {
	if len(b) != 16 {
		return ""
	}
	return fmt.Sprintf("%.2x%.2x %.2x%.2x %.2x%.2x %.2x%.2x %.2x%.2x %.2x%.2x %.2x%.2x %.2x%.2x",
		b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15])
}

func main() {
	spriteChunk, err := fetchProbableSpriteTile()
	if err != nil {
		panic(err)
	}

	fmt.Println("The Sprite Chunk")
	for i := 0; i < len(spriteChunk)-16; i += 16 {
		chunk := spriteChunk[i : i+16]
		fmt.Printf("%s\n", formatByteLine16(chunk))
	}
}
