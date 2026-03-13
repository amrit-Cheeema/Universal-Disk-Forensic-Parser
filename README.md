# Universal Disk Forensics Parser (Go)
Status: Currently In Progress (Active Development)

A systems-level utility written in Go designed to interface with raw disk images. This project aims to provide a unified interface for parsing and traversing diverse filesystem structures, beginning with legacy FAT and expanding into NTFS (Windows) and APFS (Apple).

---

### Project Goals
The objective is to build a cross-platform forensic tool that can:

* Identify File Systems: Detect magic bytes and signatures (e.g., NTFS at offset 3 or NXSB for APFS).
* Map Geometry: Calculate offsets for Data Regions, MFT (Master File Table), and Checkpoint Areas.
* Low-Level Extraction: Read raw binary data without mounting the drive, ensuring forensic integrity.

---

### Current Capabilities (FAT12/16)
The initial module focuses on the BIOS Parameter Block (BPB), the foundation of volume recognition:

* Signature Verification: Validates the JmpBoot instruction (0xEB or 0xE9) used by BIOS.
* Metadata Extraction: Decodes OEM labels and allocation unit sizes (Sectors per Cluster).
* Endian-Aware Parsing: Utilizes encoding/binary for precise Little-Endian conversion of disk metadata.