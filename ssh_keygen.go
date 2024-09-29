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

func main() {
	// SIGINT sinyallerini yakalamak için kanal oluşturma
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Anahtar oluşturma işlemi tamamlandığında programı sonlandırmak için kanal
	done := make(chan bool, 1)

	go func() {
		<-sigs
		fmt.Println("\nSinyal yakalandı, program sonlandırılıyor...")
		done <- true
	}()

	// Anahtar oluşturma işlemi
	fmt.Println("SSH anahtarları oluşturuluyor...")

	keySize := 2048

	// Özel anahtar oluşturma
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		fmt.Println("Özel anahtar oluşturulurken hata:", err)
		return
	}

	// Özel anahtarı dosyaya kaydetme
	privateKeyFile, err := os.Create("id_rsa")
	if err != nil {
		fmt.Println("Özel anahtar dosyası oluşturulamadı:", err)
		return
	}
	defer privateKeyFile.Close()

	// PEM formatında kodlayarak dosyaya yazma
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)
	privateKeyFile.Write(privateKeyPEM)

	// Public anahtar oluşturma (SSH formatında)
	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		fmt.Println("Genel anahtar oluşturulurken hata:", err)
		return
	}

	// Public anahtarı dosyaya kaydetme
	publicKeyFile, err := os.Create("id_rsa.pub")
	if err != nil {
		fmt.Println("Genel anahtar dosyası oluşturulamadı:", err)
		return
	}
	defer publicKeyFile.Close()

	// SSH formatındaki public key'i yazma
	publicKeyBytes := ssh.MarshalAuthorizedKey(publicKey)
	publicKeyFile.Write(publicKeyBytes)

	fmt.Println("SSH anahtarları başarıyla oluşturuldu.")

	// Programı sonlandırmak için sinyal bekliyoruz
	fmt.Println("Çıkmak için Ctrl + C'ye basın...")
	<-done // Kullanıcı Ctrl + C'ye basana kadar burada bekler
	fmt.Println("Program kapatıldı.")

	// Output:
	// SSH anahtarları oluşturuluyor...
	// Özel anahtar oluşturulurken hata: crypto/rsa: invalid key size for private key
	// Özel anahtar dosyası oluşturulamadı: open id_rsa: no such file or directory
	// Genel anahtar oluşturulurken hata: ssh: unsupported key type
	// Genel anahtar dosyası oluşturulamadı: open id_rsa.pub: no such file or directory
	// SSH anahtarları başarıyla oluşturuldu.
	// ��ıkmak için Ctrl + C'ye basın...
	// Program kapatıldı.
	// Program kapatıldı.
	// $
	// $

}
