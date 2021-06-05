package envr

import (
  "fmt"
  "testing"
  r "github.com/stretchr/testify/require"
)

func Test_ReadEnvFile(t *testing.T) {
  m, err := EnvFileToMap("testfiles/.env")
  r.Nil(t, err, "Unexpected error")
  r.Equal(t, "/bin:/usr/bin:/usr/local/sbin", m["PATH"], "")
  r.Equal(t, "\"postgres\"", m["DB_TYPE"], "")
  r.Equal(t, "\"localhost\"", m["DB_HOST"], "")
}

func Test_ReadBadEnvFile(t *testing.T) {
  v, err := EnvFileToMap("testfiles/.env.bad")
  r.NotNil(t, err, fmt.Sprintf("Unexpected nil error: %v", v))
}

func Test_ReadNilEnvFile(t *testing.T) {
  v, err := EnvFileToMap("testfiles/.env.nil")
  r.NotNil(t, err, fmt.Sprintf("Unexpected nil error: %v", v))
}
