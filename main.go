package main

import (
    "bytes"
    "io"
	"log"
    "fmt"
	"os"
	// "github.com/AmritKhalsa/fossInteligence/fat"
)
// findSignature scans an os.File for a byte pattern without loading the whole file.
//  - It returns the offset array of matchs found, or empty array if not found
//  - Update bufferSize for large signitures. sig > 10kB
func findSignature(file *os.File , signature []byte) (sigAddrs []int64, err error) {
	
	const bufferSize = 64 * 1024 // 64KB chunks
	buffer := make([]byte, bufferSize)
    var offset int64 = 0 // used to store where we are in the file
	sigLen := len(signature)

    sigAddrs = []int64{} // return array 

	for {
        // Reading n bytes at offset
		_, err := file.Seek(offset, io.SeekStart)
        if err != nil {return sigAddrs, err}
        n, err := file.Read(buffer)
        if (n == 0 || err == io.EOF) { break } // Break condition: EOF -> break
        if err != nil {return sigAddrs, err}
        // buffer[:n] now has the correct data to search through
        
        
        
        // Look for all signature in the current buffer, if found append to sigAddrs
		searchBuf := buffer[:n]
		localPos := 0
		for {
			idx := bytes.Index(searchBuf[localPos:], signature) // see if there is signature in local search buffer
			if idx == -1 {break} // nothing found
			absolutePos := offset + int64(localPos) + int64(idx)
			sigAddrs = append(sigAddrs, absolutePos)
			
			// Move localPos past this match to find the next one in the same buffer, until we have found all signitures
			localPos += idx + 1 
		}


        // Calculating next offset
		// We move forward by (n - sigLen + 1) to ensure the next read starts late enough to catch a signature split across the boundary.
		nextStep := int64(n) - int64(sigLen) + 1
		if nextStep <= 0 {
			offset += int64(n) // If the signature is larger than what we read, just move by n
		} else {
			offset += nextStep
		}
    }

	return sigAddrs, nil
}
func main(){
    file, err := os.Open("firmware.bin")
    if err != nil {
		log.Fatal(err)
	}
    defer file.Close()

    
    sigs := findSignature(file, []byte{0x5F, 0x46, 0x56, 0x48})
	fmt.Println(sigs)
}