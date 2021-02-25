package test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

//	viper 从环境变量或配置文件中读取信息
func TestViper(t *testing.T) {
	//从环境变量中读取CORE_**_**
	viper.SetEnvPrefix("core")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	file := filepath.Join("./" + "core.yaml")
	viper.SetConfigFile(file)
	//从文件./config.* 读取信息
	//viper.SetConfigName("core")
	//viper.AddConfigPath("./")

	gopath := os.Getenv("GOPATH")

	for _, p := range filepath.SplitList(gopath) {
		peerpath := filepath.Join(p, "src/vip")
		viper.AddConfigPath(peerpath)
	}

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		t.Errorf("Fatal error when reading config file: %s\n", err)
	}

	//export CORE_SECURITY_ENABLED=true
	environment := viper.GetBool("security.enabled")
	t.Log("environment:", environment) //environment:true

	fullstate := viper.GetString("statetransfer.timeout.fullstate")
	t.Logf("fullstate:%s", fullstate)

	abcdValuea := viper.GetString("peer.abcd")
	t.Log("abcdValuea is:", abcdValuea)
}

func TestFlag(t *testing.T) {
	file := filepath.Join("./core.yaml")
	t.Log(file)
	// t.Errorf(file)
}

func TestENVConfig(t *testing.T) {
	file := filepath.Join("./core.yaml")
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		t.Errorf("viper.ReadInConfig error: %s", err.Error())
		return
	}
	envpath := viper.GetString("test")
	envpath = Subst(envpath)
	t.Logf("value: %s", envpath)
	return
}

// Subst replaces instances of '${VARNAME}' (eg ${GOPATH}) with the variable.
// Variables names that are not set by the SDK are replaced with the environment variable.
func Subst(path string) string {
	const (
		sepPrefix = "${"
		sepSuffix = "}"
	)

	splits := strings.Split(path, sepPrefix)

	var buffer bytes.Buffer

	// first split precedes the first sepPrefix so should always be written
	buffer.WriteString(splits[0]) // nolint: gas

	for _, s := range splits[1:] {
		subst, rest := substVar(s, sepPrefix, sepSuffix)
		buffer.WriteString(subst) // nolint: gas
		buffer.WriteString(rest)  // nolint: gas
	}

	return buffer.String()
}

// substVar searches for an instance of a variables name and replaces them with their value.
// The first return value is substituted portion of the string or noMatch if no replacement occurred.
// The second return value is the unconsumed portion of s.
func substVar(s string, noMatch string, sep string) (string, string) {
	endPos := strings.Index(s, sep)
	if endPos == -1 {
		return noMatch, s
	}

	v, ok := lookupVar(s[:endPos])
	if !ok {
		return noMatch, s
	}

	return v, s[endPos+1:]
}

// lookupVar returns the value of the variable.
// The local variable table is consulted first, followed by environment variables.
// Returns false if the variable doesn't exist.
func lookupVar(v string) (string, bool) {
	// TODO: optimize if the number of variable names grows
	//switch v {
	////case "FABRIC_SDK_GO_PROJECT_PATH":
	////	return metadata.GetProjectPath(), true
	////case "GOPATH":
	////	return goPath(), true
	////case "CRYPTOCONFIG_FIXTURES_PATH":
	////	return metadata.CryptoConfigPath, true
	//case "GOHFC_PATH":
	//	return os.Getenv("GOHFC_PATH"), true
	//}
	return os.LookupEnv(v)
}

/***
//core.yaml
statetransfer:
    recoverdamage: true
    blocksperrequest: 20
    maxdeltas: 200
    timeout:
        singleblock: 2s
        singlestatedelta: 2s
        fullstate: 60s
peer:
    abcd:   3322d
*/

// #go build main exec.out
// #$CORE_SECURITY_ENABLED=true ./exec.out == $CORE_SECURITY_ENABLED=true go run main.go

// #$CORE_SECURITY_ENABLED=true
// #$./exec.out

/**
package main
import "fmt"
var (
    VERSION    string
    BUILD_TIME string
    GO_VERSION string
)
func main() {
    fmt.Printf("%s\n%s\n%s\n", VERSION, BUILD_TIME, GO_VERSION)
}
go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'"
编译后运行显示：	//date go version 为命令，从环境中读取信息
1.0.0
Thu Nov 29 11:48:46 CST 2018
go version go1.10.5 linux/amd64
*/
