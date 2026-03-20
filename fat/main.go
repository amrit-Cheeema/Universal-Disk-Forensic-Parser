package fat
import (

	"encoding/binary"
	"fmt"
	"log"
	"os"
)
type Fat struct {
    // BIOS PARAMETER BLOCK
    BPB [512]byte

    // Jump instruction to boot code
    // This field has two allowed forms
    //  - [0xEB, offset, 0x90(NOP)] (jmp to (signed)offset)
    //  - [0xE9, offset1, offset2] (jmp to (signed)(offset2(H) offset1(L)) bytes)
    JmpBoot [3]byte

    OEMName string

    // This value may take on only the following values: 512, 1024, 2048 or 4096
    SecPerClus uint16
    
    // Number of file allocation tables
    NumFATs uint8

    // Count of 32-byte directory entries in the root directory
    //  - [FAT32] Value should be 00 00 for 32 bit volumes
    //  - [FAT12, FAT16] Default = 512. This should follow the form: fat.RootEntCnt * 32 = k * fat.BytsPerSec, where k is Even
    RootEntCnt uint16
}
func (f *Fat) Init(file *os.File) {
    buffer := make([]byte, 512)

	bytesRead, err := file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("Read %d bytes successfully.\n", bytesRead)
	f.BPB = [512]byte(buffer)
    

    // Configuring data structure
    copy(f.JmpBoot[:], f.BPB[0:3])
    f.OEMName = string(f.BPB[3:11])

    f.SecPerClus = binary.LittleEndian.Uint16(f.BPB[11:13])
    f.NumFATs = f.BPB[16]
    f.RootEntCnt = binary.LittleEndian.Uint16(f.BPB[17:19])

}