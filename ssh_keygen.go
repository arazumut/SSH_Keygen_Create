package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Kanal oluşturma ve sinyal yakalama
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		<-sigs
		fmt.Println("\nSinyal yakalandı, program sonlandırılıyor...")
		done <- true
	}()

	fmt.Println("SSH anahtarları oluşturuluyor...")

	// Anahtar boyutu
	keySize := 2048

	// Özel anahtar oluşturma
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		log.Fatalf("Özel anahtar oluşturulurken hata: %v", err)
	}

	// Özel anahtarı PEM formatında kodlama
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	// Özel anahtarı dosyaya yazma
	err = os.WriteFile("id_rsa", privateKeyPEM, 0600)
	if err != nil {
		log.Fatalf("Özel anahtar dosyaya yazılırken hata: %v", err)
	}

	// Genel anahtar oluşturma
	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatalf("Genel anahtar oluşturulurken hata: %v", err)
	}

	// Genel anahtarı dosyaya yazma
	err = os.WriteFile("id_rsa.pub", ssh.MarshalAuthorizedKey(publicKey), 0644)
	if err != nil {
		log.Fatalf("Genel anahtar dosyaya yazılırken hata: %v", err)
	}

	fmt.Println("SSH anahtarları başarıyla oluşturuldu ve kaydedildi.")
	fmt.Println("Çıkmak için CTRL+C.")

	<-done
}
