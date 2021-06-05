package envr

import (
  "errors"
  "bufio"
  "fmt"
  "io"
  "os"
)

const (
  TokenAssign = byte('=')
  TokenNewLine = byte('\n')
)

type StringPair struct {
  First string
  Second string
}

func ParseEnvFileToPairArray(r io.Reader) ([]StringPair, error) {
  b := bufio.NewReader(r)

  stack := make([][]byte, 1)
  for {
    token, e := b.ReadByte()
    if e != nil && errors.Is(e, io.EOF) {
      break
    }

    if token == TokenNewLine {
      _, eofErr := b.Peek(1)
      if errors.Is(eofErr, io.EOF) {
        break
      }
    }

    if token == TokenAssign || token == TokenNewLine {
      stack = append(stack, []byte{})
      continue
    }

    stackTop := len(stack) - 1
    stack[stackTop] = append(stack[stackTop], token)
  }

  if len(stack) % 2 != 0 {
    return nil, fmt.Errorf(
      "Invalid env file: missing assignment.\nCurrent stack %v",
      string(stack[len(stack) - 1]),
    )
  }

  pairs := make([]StringPair, len(stack) / 2)
  for i := 0; i < len(pairs); i++ {
    pair := StringPair{
      First: string(stack[2 * i]),
      Second: string(stack[(2 * i) + 1]),
    }

    if len(pair.First) == 0 || len(pair.Second) == 0 {
      return nil, fmt.Errorf(
        "Invalid env file: missing assignment for -> %s = %s",
        pair.First,
        pair.Second,
      )
    }

    pairs[i] = pair
  }

  return pairs, nil
}

func PairsToMap(pairs []StringPair) map[string]string {
  m := map[string]string{}
  for _, p := range pairs {
    m[p.First] = p.Second
  }
  return m
}

func EnvFileToMap(fp string) (map[string]string, error) {
  fh, err := os.Open(fp)
  if err != nil {
    return nil, err
  }
  defer fh.Close()

  fh.Seek(0, 0)

  pairs, err := ParseEnvFileToPairArray(fh)
  if err != nil {
    return nil, err
  }

  return PairsToMap(pairs), nil
}

