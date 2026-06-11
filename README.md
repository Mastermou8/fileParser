# fileParser

A Go CLI tool that extracts the first page from every PDF file in a directory and saves the result.

## How It Works

The program reads all PDF files from the `./pdf` directory and uses [`pdfcpu`](https://github.com/pdfcpu/pdfcpu) to extract the first page of each file, writing the output back into the same directory.

You can also pass a specific file path as a command-line argument to override the directory scan:

```bash
go run main.go path/to/file.pdf
```

## Project Structure

```
.
├── main.go                  # Entry point — reads PDFs from ./pdf and processes them
├── internal/
│   └── utils/
│       └── utils.go         # ExtractFirstPage helper using pdfcpu
├── pdf/                     # Input directory — place PDF files here
├── out/                     # Output directory (created automatically)
├── go.mod
└── go.sum
```

## Requirements

- Go 1.21+
- PDF files placed in the `./pdf` directory

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/Mastermou8/fileParser.git
   cd fileParser
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Place one or more PDF files in the `./pdf` directory.

4. Run the tool:

   ```bash
   go run main.go
   ```

Extracted first-page PDFs will be saved in the `./pdf` directory alongside the originals.

## Dependencies

| Package | Purpose |
|---|---|
| [pdfcpu](https://github.com/pdfcpu/pdfcpu) | PDF processing and page extraction |
