package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func gencert() {

	// serialNumber 本来は認証局によって発行される一意の番号だが、
	// テスト的に証明書を生成する目的のためランダムな大きい整数を使用
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)

	// subject 証明書の識別子
	subject := pkix.Name{
		Organization:       []string{"Organization"},
		OrganizationalUnit: []string{"OrganizationalUnit"},
		CommonName:         "gocomm",
	}

	// template 証明書の構成を設定するための構造体
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	// pk RSAの秘密鍵を生成
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	// derBytes 構造体'Certificate'と公開鍵、秘密鍵'pk'を用いてDER形式のバイトデータのスライスを生成
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)

	// 証明書データを符号化して'cert.pem'というファイを保存
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	// 生成した鍵をPEM符号化して'key.pem'というファイルを保存
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()

}
