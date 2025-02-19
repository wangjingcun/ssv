package testing

import (
	"crypto/rsa"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pkg/errors"
	specqbft "github.com/ssvlabs/ssv-spec/qbft"
	spectypes "github.com/ssvlabs/ssv-spec/types"
	"github.com/ssvlabs/ssv-spec/types/testingutils"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"

	qbftstorage "github.com/ssvlabs/ssv/protocol/v2/qbft/storage"
	"github.com/ssvlabs/ssv/utils/rsaencryption"
)

var (
	specModule = "github.com/ssvlabs/ssv-spec"
)

// TODO: add missing tests

// GenerateOperatorSigner generates randomly nodes
func GenerateOperatorSigner(oids ...spectypes.OperatorID) ([]*rsa.PrivateKey, []*spectypes.Operator) {
	nodes := make([]*spectypes.Operator, 0, len(oids))
	sks := make([]*rsa.PrivateKey, 0, len(oids))

	for i := range oids {
		pubKey, privKey, err := rsaencryption.GenerateKeys()
		if err != nil {
			panic(err)
		}
		opKey, err := rsaencryption.PemToPrivateKey(privKey)
		if err != nil {
			panic(err)
		}

		nodes = append(nodes, &spectypes.Operator{
			OperatorID:        oids[i],
			SSVOperatorPubKey: pubKey,
		})

		sks = append(sks, opKey)
	}

	return sks, nodes
}

// MsgGenerator represents a message generator
type MsgGenerator func(height specqbft.Height) ([]spectypes.OperatorID, *specqbft.Message)

// CreateMultipleStoredInstances enables to create multiple stored instances (with decided messages).
func CreateMultipleStoredInstances(
	sks []*rsa.PrivateKey,
	start specqbft.Height,
	end specqbft.Height,
	generator MsgGenerator,
) ([]*qbftstorage.StoredInstance, error) {
	results := make([]*qbftstorage.StoredInstance, 0)
	for i := start; i <= end; i++ {
		signers, msg := generator(i)
		if msg == nil {
			break
		}
		sm := testingutils.MultiSignQBFTMsg(sks, signers, msg)

		var qbftMsg specqbft.Message
		if err := qbftMsg.Decode(sm.SSVMessage.Data); err != nil {
			return nil, err
		}

		results = append(results, &qbftstorage.StoredInstance{
			State: &specqbft.State{
				ID:                   qbftMsg.Identifier,
				Round:                qbftMsg.Round,
				Height:               qbftMsg.Height,
				LastPreparedRound:    qbftMsg.Round,
				LastPreparedValue:    sm.FullData,
				Decided:              true,
				DecidedValue:         sm.FullData,
				ProposeContainer:     specqbft.NewMsgContainer(),
				PrepareContainer:     specqbft.NewMsgContainer(),
				CommitContainer:      specqbft.NewMsgContainer(),
				RoundChangeContainer: specqbft.NewMsgContainer(),
			},
			DecidedMessage: sm,
		})
	}
	return results, nil
}

// SignMsg handle MultiSignMsg error and return just specqbft.SignedMessage
func SignMsg(t *testing.T, sks []*rsa.PrivateKey, signers []spectypes.OperatorID, msg *specqbft.Message) *spectypes.SignedSSVMessage {
	return testingutils.MultiSignQBFTMsg(sks, signers, msg)
}

func GenerateSpecTestJSON(path string, module string) ([]byte, error) {
	// Step 1: Get the spec directory.
	p, err := GetSpecDir(path, module)
	if err != nil {
		return nil, fmt.Errorf("could not get spec test dir: %w", err)
	}

	p = filepath.Join(p, "spectest", "generate")

	// Step 2: Create a temporary directory at /tmp/<module>.
	tmpDir := filepath.Join("/tmp", module)
	if err := os.MkdirAll(tmpDir, 0750); err != nil {
		return nil, fmt.Errorf("failed to create tmp directory: %w", err)
	}
	// Clean up the temp directory when the function exits.
	defer func() {
		err := os.RemoveAll(tmpDir)
		if err != nil {
			fmt.Printf("failed to remove tmp directory: %s", err.Error())
		}
	}()

	// Step 3: Build the Go package, outputting an executable to tmpDir.
	// We'll name the executable after the module.
	binaryPath := filepath.Join(tmpDir, module)
	// nolint: gosec
	cmdBuild := exec.Command("go", "build", "-o", binaryPath, ".")
	cmdBuild.Dir = p
	buildOutput, err := cmdBuild.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("go build failed: %w; output: %s", err, buildOutput)
	}

	// Step 4: Execute the built binary.
	// It is assumed that running the binary generates tests.json in tmpDir.
	// nolint: gosec
	cmdRun := exec.Command(binaryPath)
	cmdRun.Dir = tmpDir
	_, err = cmdRun.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to run binary: %w", err)
	}

	// Step 5: Read the tests.json file generated by the binary.
	testJSONPath := filepath.Join(tmpDir, "tests.json")
	// nolint: gosec
	jsonBytes, err := os.ReadFile(testJSONPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read tests.json: %w", err)
	}

	return jsonBytes, nil
}

// GetSpecDir returns the path to the ssv-spec module.
func GetSpecDir(path, module string) (string, error) {
	if path == "" {
		var err error
		path, err = os.Getwd()
		if err != nil {
			return "", errors.New("could not get current directory")
		}
	}
	goModFile, err := getGoModFile(path)
	if err != nil {
		return "", errors.New("could not get go.mod file")
	}

	// check if there is a replace
	var modPath, modVersion string
	var replace *modfile.Replace
	for _, r := range goModFile.Replace {
		if strings.EqualFold(specModule, r.Old.Path) {
			replace = r
			break
		}
	}

	if replace != nil {
		modPath = replace.New.Path
		modVersion = replace.New.Version
	} else {
		// get from require
		var req *modfile.Require
		for _, r := range goModFile.Require {
			if strings.EqualFold(specModule, r.Mod.Path) {
				req = r
				break
			}
		}
		if req == nil {
			return "", errors.Errorf("could not find %s module", specModule)
		}
		modPath = req.Mod.Path
		modVersion = req.Mod.Version
	}

	// get module path
	p, err := GetModulePath(modPath, modVersion)
	if err != nil {
		return "", errors.Wrap(err, "could not get module path")
	}

	if _, err := os.Stat(p); os.IsNotExist(err) {
		return "", errors.Wrapf(err, "you don't have this module-%s/version-%s installed", modPath, modVersion)
	}

	return filepath.Join(filepath.Clean(p), module), nil
}

func GetModulePath(name, version string) (string, error) {
	// first we need GOMODCACHE
	cache, ok := os.LookupEnv("GOMODCACHE")
	if !ok {
		cache = path.Join(os.Getenv("GOPATH"), "pkg", "mod")
	}

	// then we need to escape path
	escapedPath, err := module.EscapePath(name)
	if err != nil {
		return "", err
	}

	// version also
	escapedVersion, err := module.EscapeVersion(version)
	if err != nil {
		return "", err
	}

	return path.Join(cache, escapedPath+"@"+escapedVersion), nil
}

func getGoModFile(path string) (*modfile.File, error) {
	// find project root path
	for {
		if _, err := os.Stat(filepath.Join(path, "go.mod")); err == nil {
			break
		}
		path = filepath.Dir(path)
		if path == "/" {
			return nil, errors.New("could not find go.mod file")
		}
	}

	// read go.mod
	buf, err := os.ReadFile(filepath.Join(filepath.Clean(path), "go.mod"))
	if err != nil {
		return nil, errors.New("could not read go.mod")
	}

	// parse go.mod
	return modfile.Parse("go.mod", buf, nil)
}
