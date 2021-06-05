package main

import (
  "os"
  "fmt"
  "flag"
  "strings"

  "github.com/brent-soles/envr"
)

type envAgg []string

func (e *envAgg) String() string {
  return "lol"
}

func (e *envAgg) Set(value string) error {
  *e = append(*e, value)
  return nil
}

var agg envAgg
var envFile string

func init() {
  flag.Var(&agg, "e", "An environment variable to be aggregated")
  flag.StringVar(&envFile, "f", ".env", ".env file")
  flag.Parse()
}

func main() {
  c := envr.CreateCmd()

  fileMap, err := envr.EnvFileToMap(envFile)
  if err == nil {
    c.SetEnvFromMap(fileMap)
  }

  cliPairs := make([]envr.StringPair, len(agg))
  for i, v := range agg {
    r := strings.NewReader(v)
    p, err := envr.ParseEnvFileToPairArray(r)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    cliPairs[i] = p[0]
  }

  c.SetEnvFromMap(envr.PairsToMap(cliPairs))

  bin := os.Args[len(os.Args) - 1]
  binChunks := strings.Split(bin, " ")

  c.SetBin(binChunks[0])
  if len(binChunks) > 1 {
    c.SetArgs(strings.Join(binChunks[1:], " "))
  }

  c.Run()
}
