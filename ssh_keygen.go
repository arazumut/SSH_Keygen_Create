package main

import (

	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"golang.org/x/crypto/ssh"

)


//Produced By K.Umut Araz

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		<-sigs
		fmt.Println("\nSinyal yakalandı, program sonlandırılıyor...")
		done <- true
	}()

	fmt.Println("SSH anahtarları oluşturuluyor...")

	keySize := 2048

	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		fmt.Println("Özel anahtar oluşturulurken hata:", err)
		return
	}

	privateKeyFile, err := os.Create("id_rsa")
	if err != nil {
		fmt.Println("Özel anahtar dosyası oluşturulamadı:", err)
		return
		
	}
	defer privateKeyFile.Close()

	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)
	_, err = privateKeyFile.Write(privateKeyPEM)
	if err != nil {
		fmt.Println("Özel anahtar dosyasına yazılamadı:", err)
		return
	}

	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		fmt.Println("Genel anahtar oluşturulurken hata:", err)
		return
	}

	publicKeyFile, err := os.Create("id_rsa.pub")
	if err != nil {
		fmt.Println("Genel anahtar dosyası oluşturulamadı:", err)
		return
	}
	defer publicKeyFile.Close()

	publicKeyBytes := ssh.MarshalAuthorizedKey(publicKey)
	_, err = publicKeyFile.Write(publicKeyBytes)
	if err != nil {
		fmt.Println("Genel anahtar dosyasına yazılamadı:", err)
		return
	}

	fmt.Println("SSH anahtarları başarıyla oluşturuldu.")

	fmt.Println("Çıkmak için Ctrl + C'ye basın...")
	<-done
	fmt.Println("Program kapatıldı.")
}
