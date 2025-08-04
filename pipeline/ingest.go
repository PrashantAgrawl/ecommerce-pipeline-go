
package pipeline

import (
    "encoding/csv"
    "os"
    "bufio"
    "io"
)

func ReadCSVInChunks(path string, chunkSize int, handler func([][]string)) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(bufio.NewReader(file))
    var records [][]string
    _, _ = reader.Read() // skip header
    count := 0
    for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            continue
        }
        records = append(records, line)
        count++
        if count >= chunkSize {
            handler(records)
            records = [][]string{}
            count = 0
        }
    }
    if len(records) > 0 {
        handler(records)
    }
    return nil
}
