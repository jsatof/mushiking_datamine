package main

import (
	"bufio"
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

		fmt.Printf("0x%.4x: %.2x %.2x %.2x %.2x %.2x %.2x %.2x %.2x %.2x %.2x %.2x %.2x %.2x %.2x %.2x %.2x\n",
			counter,
			byteDump[i],
			byteDump[i+1],
			byteDump[i+2],
			byteDump[i+3],
			byteDump[i+4],
			byteDump[i+5],
			byteDump[i+6],
			byteDump[i+7],
			byteDump[i+8],
			byteDump[i+9],
			byteDump[i+10],
			byteDump[i+11],
			byteDump[i+12],
			byteDump[i+13],
			byteDump[i+14],
			byteDump[i+15],
		)

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

func fetchProbableSpriteTile() error {
	romContent, err := os.ReadFile("KouchuOujaMushiking_J.gba")
	if err != nil {
		return err
	}

	fmt.Println("Length of rom:", len(romContent))

	target := []byte{0x3e, 0xdb, 0x40, 0xc0, 0x0e, 0xe3, 0x0e, 0xa8, 0x40, 0x44, 0x55, 0x00, 0x44, 0x34, 0x00, 0x24}

	for i := 0; i < len(romContent)-16; i++ {
		compareString := romContent[i : i+16]
		if SlicesEqual(compareString, target) {
			fmt.Printf("Found bytes %s at index %d\n", target, i)
			break
		}
	}

	return nil
}

func main() {
	err := fetchProbableSpriteTile()
	if err != nil {
		panic(err)
	}
}
