package buildin

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Cipher struct {
	cb cipher.Block
}

func NewCipher(key string) (*Cipher, error) {
	if len(key) > aes.BlockSize {
		return nil, errors.New("key len should less or equal to 16")
	}
	keybuf := make([]byte, aes.BlockSize)
	copy(keybuf, []byte(key))
	block, err := aes.NewCipher(keybuf)
	if err != nil {
		return nil, err
	}

	return &Cipher{
		cb: block,
	}, nil
}

func (c *Cipher) Encrypt(text []byte) []byte {
	plainText := make([]byte, len(text)+aes.BlockSize+(-len(text)%aes.BlockSize))
	copy(plainText, text)
	encryptText := make([]byte, aes.BlockSize+len(plainText))
	iv := encryptText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		// TODO do some log
	}
	mode := cipher.NewCBCEncrypter(c.cb, iv)
	mode.CryptBlocks(encryptText[aes.BlockSize:], plainText)

	return encryptText
}

func (c *Cipher) Decrypt(encryptText []byte) ([]byte, error) {
	if len(encryptText) < aes.BlockSize {
		return nil, errors.New("cipher text too short")
	}

	iv := encryptText[:aes.BlockSize]
	encryptText = encryptText[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(c.cb, iv)
	mode.CryptBlocks(encryptText, encryptText)

	return bytes.TrimRight(encryptText, "\x00"), nil
}

func TestAes(t *testing.T) {
	Convey("test cbc", t, func() {
		c, err := NewCipher("password")
		So(err, ShouldBeNil)
		So(c, ShouldNotBeNil)

		encryptText := c.Encrypt([]byte("hello world"))
		text, err := c.Decrypt(encryptText)
		So(text, ShouldResemble, []byte("hello world"))
	})
}

func TestMd5(t *testing.T) {
	Convey("test md5", t, func() {
		h := md5.New()
		h.Write([]byte("hello world"))
		So(h.Size(), ShouldEqual, 16)
		So(hex.EncodeToString(h.Sum(nil)), ShouldEqual, "5eb63bbbe01eeed093cb22bb8f5acdc3")
	})
}

func TestSha256(t *testing.T) {
	Convey("test sha 1", t, func() {
		h := sha1.New()
		h.Write([]byte("hello world"))
		So(h.Size(), ShouldEqual, 20)
		So(hex.EncodeToString(h.Sum(nil)), ShouldEqual, "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
	})

	Convey("test sha 224", t, func() {
		h := sha256.New224()
		h.Write([]byte("hello world"))
		So(h.Size(), ShouldEqual, 28)
		So(hex.EncodeToString(h.Sum(nil)), ShouldEqual, "2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")

		hashed := sha256.Sum224([]byte("hello world"))
		So(hex.EncodeToString(hashed[:]), ShouldEqual, "2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
	})

	Convey("test sha 256", t, func() {
		h := sha256.New()
		h.Write([]byte("hello world"))
		So(h.Size(), ShouldEqual, 32)
		So(hex.EncodeToString(h.Sum(nil)), ShouldEqual, "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")

		hashed := sha256.Sum256([]byte("hello world"))
		So(hex.EncodeToString(hashed[:]), ShouldEqual, "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
	})

	Convey("test sha 256", t, func() {
		h := sha512.New()
		h.Write([]byte("hello world"))
		So(h.Size(), ShouldEqual, 64)
		So(hex.EncodeToString(h.Sum(nil)), ShouldEqual, "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
	})
}

func TestDsa(t *testing.T) {
	Convey("test dsa", t, func() {
		params := &dsa.Parameters{}
		So(dsa.GenerateParameters(params, rand.Reader, dsa.L1024N160), ShouldBeNil)

		privateKey := &dsa.PrivateKey{}
		privateKey.Parameters = *params
		So(dsa.GenerateKey(privateKey, rand.Reader), ShouldBeNil)

		message := []byte("hello world")

		r, s, err := dsa.Sign(rand.Reader, privateKey, message)
		So(err, ShouldBeNil)

		So(dsa.Verify(&privateKey.PublicKey, message, r, s), ShouldBeTrue)
	})
}

func TestEcdsa(t *testing.T) {
	Convey("test ecdsa", t, func() {
		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		So(err, ShouldBeNil)

		message := []byte("hello world")

		r, s, err := ecdsa.Sign(rand.Reader, privateKey, message)
		So(err, ShouldBeNil)

		So(ecdsa.Verify(&privateKey.PublicKey, message, r, s), ShouldBeTrue)
	})
}

func TestEd25519(t *testing.T) {
	Convey("test ed25519", t, func() {
		publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
		So(err, ShouldBeNil)

		message := []byte("hello world")
		buf := ed25519.Sign(privateKey, message)

		So(ed25519.Verify(publicKey, message, buf), ShouldBeTrue)
	})
}

func TestRsa(t *testing.T) {
	Convey("test rsa", t, func() {
		privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
		So(err, ShouldBeNil)

		message := []byte("hello world")

		Convey("test encrypt/decrypt", func() {
			encryptBuf, err := rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, message)
			So(err, ShouldBeNil)

			decryptBuf, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptBuf)
			So(err, ShouldBeNil)
			So(string(decryptBuf), ShouldEqual, "hello world")
		})

		Convey("test sign/verify", func() {
			hashed := sha256.Sum256(message)

			buf, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
			So(err, ShouldBeNil)
			So(rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, hashed[:], buf), ShouldBeNil)
		})
	})
}

func TestHamc(t *testing.T) {
	Convey("test hmac", t, func() {
		mac := hmac.New(sha256.New, []byte("123456"))
		mac.Write([]byte("hello world"))
		So(hex.EncodeToString(mac.Sum(nil)), ShouldEqual, "83b3eb2788457b46a2f17aaa048f795af0d9dabb8e5924dd2fc0ea682d929fe5")
	})
}

func TestX509(t *testing.T) {
	Convey("test x509 rsa", t, func() {
		privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
		So(err, ShouldBeNil)

		{
			buffer := &bytes.Buffer{}
			So(pem.Encode(buffer, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}), ShouldBeNil)
			fmt.Println(buffer.String())

			block, _ := pem.Decode(buffer.Bytes())
			pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
			So(err, ShouldBeNil)
			So(pk, ShouldResemble, privateKey)
		}
		{
			buffer := &bytes.Buffer{}
			So(pem.Encode(buffer, &pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)}), ShouldBeNil)
			fmt.Println(buffer.String())

			block, _ := pem.Decode(buffer.Bytes())
			pk, err := x509.ParsePKCS1PublicKey(block.Bytes)
			So(err, ShouldBeNil)
			So(pk, ShouldResemble, &privateKey.PublicKey)
		}
	})
}

func TestCryptoRand(t *testing.T) {
	Convey("test rand", t, func() {
		token := make([]byte, 16)
		_, _ = rand.Read(token)
		fmt.Println(hex.EncodeToString(token))
	})
}
