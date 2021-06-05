package envr

import (
  "os/exec"
  "fmt"
  "testing"
  r "github.com/stretchr/testify/require"
)


func Test_Cmd(t *testing.T) {
  m := map[string]string{
    "HOME": "/home/user",
    "PATH": "/bin:/usr/sbin",
  }
  c := CreateCmd()
  c.SetEnvFromMap(m)

  for k, v := range c.Env() {
    vCheck, ok := m[k]

    r.True(t, ok)
    r.Equal(t, v, vCheck, "kv map doesn't match")
  }
}

func Test_CmdRun(t *testing.T) {
  c := CreateCmd()
  env, err := EnvFileToMap("testfiles/.env")
  r.Nil(t, err, "Unexpected error")

  err = c.
    SetEnvFromMap(env).
    SetBin("testfiles/test.sh").
    Run()

  sourceErrIfAvailable := func() string {
    if err == nil {
      return ""
    }
    e := err.(*exec.ExitError)
    return fmt.Sprintf("Unexpected error: %v", e.ProcessState)
  }

  r.Nil(t, err, sourceErrIfAvailable())
}
