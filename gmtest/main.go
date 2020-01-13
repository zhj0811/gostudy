package main

import (
	"bytes"
	"github.com/tjfoc/gmsm/sm2"
	"fmt"
	"os"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("gmtest")

// package-scoped constants
// const packageName = "apiserver"

func init() {
	format := logging.MustStringFormatter("%{shortfile} %{time:2006-01-02 15:04:05.000} [%{module}] %{level:.4s} : %{message}")
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)

	logging.SetBackend(backendFormatter)
}

func main(){
	priv, err := sm2.GenerateKey() // 生成密钥对
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("priv %v", priv)
	msg := []byte("Tongji Fintech Research Institute")
	pub := &priv.PublicKey
	log.Infof("public key is %v", pub)
	ciphertxt, err := pub.Encrypt(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("加密结果:%x\n",ciphertxt)
	plaintxt,err :=  priv.Decrypt(ciphertxt)
	if err != nil {
		log.Fatal(err)
	}

	if !bytes.Equal(msg,plaintxt){
		log.Fatal("原文不匹配")
	}

	r,s,err := sm2.Sign(priv, msg)
	if err != nil {
		log.Fatal(err)
	}
	isok := sm2.Verify(pub,msg,r,s)
	fmt.Printf("Verified: %v\n", isok)
}


