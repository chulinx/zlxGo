package ssh

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"golang.org/x/crypto/ssh"
	"os"
)

// @Des generate ssh pubKey and privateKey
// @Param privateKeyPath string
// @Return err, pubKeyString
func MakeSSHKeyPair(privateKeyPath string) (error, string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return err, ""
	}

	// generate and write private key as PEM
	privateKeyFile, err := os.Create(privateKeyPath)
	defer privateKeyFile.Close()
	if err != nil {
		return err, ""
	}
	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		return err, ""
	}

	// generate and write public key
	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err, ""
	}
	return nil, string(ssh.MarshalAuthorizedKey(pub))
}
