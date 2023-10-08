package create

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

type nonProtocolTypeArgs struct {
	typ    string
	outdir string
}

type nonProtocolTypeTestCase struct {
	name    string
	args    nonProtocolTypeArgs
	wantErr bool
}

func TestCreateCmdByNonProtocolType(t *testing.T) {
	require.Nil(t, setup(nil))

	pwd, _ := os.Getwd()
	defer func() {
		os.Chdir(pwd)
	}()

	tests := setNonProtocolTypeTestCases()

	out := filepath.Join(os.TempDir(), "create/generated-non-protocol-type")
	os.RemoveAll(out)

	for _, tt := range tests {

		opts := []string{}

		opts = append(opts, "--non-protocol-type", tt.args.typ)
		outdir := filepath.Join(out, tt.args.typ)
		opts = append(opts, "-o", outdir)
		opts = append(opts, "--mock", "false") // No protocol types are needed for mock operations.

		resetFlags(createCmd)
		runCreateCmd(t, tt.name, opts, outdir, tt.wantErr)
	}
}

func setNonProtocolTypeTestCases() []nonProtocolTypeTestCase {
	tests := []nonProtocolTypeTestCase{
		{
			"non-protocol-type-http",
			nonProtocolTypeArgs{"http", "http"},
			false,
		}, {
			"non-protocol-type-kafka",
			nonProtocolTypeArgs{"kafka", "kafka"},
			false,
		}, {
			"non-protocol-type-hippo",
			nonProtocolTypeArgs{"hippo", "hippo"},
			false,
		},
	}
	return tests
}
