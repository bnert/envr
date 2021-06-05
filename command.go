package envr

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
)

func mapToArray(m map[string]string, f func(string, string) string) []string {
  a := make([]string, len(m))

  i := 0
  for k, v := range m {
    a[i] = f(k, v)
    i++
  }

  return a
}

type Cmd struct {
  env map[string]string
  bin string
  args string
}

func (c *Cmd) SetEnv(k, v string) {
  c.env[k] = v
}

func (c *Cmd) SetEnvFromMap(m map[string]string) *Cmd {
  for k, v := range m {
    c.SetEnv(k, v)
  }
  return c
}

func (c *Cmd) SetBin(b string) *Cmd {
  c.bin = b
  return c
}

func (c *Cmd) SetArgs(a string) *Cmd {
  c.args = a
  return c
}

func (c *Cmd) Env() map[string]string {
  return c.env
}

func (cmd *Cmd) Run() error {
  c := exec.Command(cmd.bin, cmd.args)
  c.Stdin = os.Stdin
  c.Stdout = os.Stdout
  c.Stderr = os.Stderr

  c.Env = mapToArray(cmd.env, func(k, v string) string {
    return strings.Join([]string{k, v}, "=")
  })
  fmt.Printf("Running bin: %s\nArgs: %v\n", cmd.bin, cmd.args)
  return c.Run()
}

func CreateCmd() *Cmd {
  c := &Cmd{
    env: map[string]string{},
  }
  return c
}

