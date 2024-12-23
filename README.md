<a href="https://golang.org/" target="_blank" rel="noreferrer"> 
        <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="golang" width="40" height="40"/> 
    </a> <br>
Go dilinde RSA tabanlı SSH anahtarları oluşturan bir programdır. Aynı zamanda, Ctrl + C gibi bir sistem sinyaline (SIGINT) yanıt vererek güvenli bir şekilde sonlanmayı sağlar. İşte kodun genel bir özeti:

Ana Özellikler:
RSA Anahtar Çifti Oluşturma:


![Ekran görüntüsü 2024-09-29 190424](https://github.com/user-attachments/assets/9d4c3b1e-1148-4d52-becc-7ec65b83d941)
![Ekran görüntüsü 2024-09-29 190327](https://github.com/user-attachments/assets/2b6f3eab-2a71-4626-b294-6492b4f53f92)



Program, 2048 bitlik bir RSA özel anahtar (private key) oluşturur.
Bu özel anahtardan bir SSH uyumlu genel anahtar (public key) üretir.
Anahtarları Dosyaya Yazma:


Özel anahtar, id_rsa adında bir dosyaya PEM formatında kaydedilir.
Genel anahtar ise, id_rsa.pub dosyasına SSH formatında kaydedilir.
Sinyal Yakalama:

Program, Ctrl + C (SIGINT) veya SIGTERM sinyallerini dinler.
Bu sinyal yakalandığında, program güvenli bir şekilde kapanır.
Hata Yönetimi:

Özel veya genel anahtarların oluşturulması veya dosyaya yazılması sırasında herhangi bir hata meydana gelirse, bu hatalar terminale yazdırılır ve program sonlandırılır.
Adımlar:
Sinyal Yakalama (SıgnInt):



Program başlar başlamaz bir goroutine içinde Ctrl + C sinyallerini yakalamaya başlar. Kullanıcı bu sinyali gönderdiğinde, program kapanacaktır.
RSA Anahtarı Oluşturma:

Program, 2048 bitlik bir RSA özel anahtarı oluşturur ve bir dosyaya PEM formatında yazar.
SSH Public Anahtar Oluşturma:

RSA özel anahtarından SSH formatında bir genel anahtar türetilir ve bu anahtar .pub dosyasına kaydedilir.
Program Sonlandırma:

Anahtar oluşturma işlemi bittikten sonra, program "Çıkmak için Ctrl + C'ye basın..." mesajını gösterir ve kullanıcıdan bir sinyal bekler.
Hata Mesajları:

Eğer anahtar oluşturulurken bir hata olursa (örneğin, dosya yazılamazsa), program hata mesajlarını terminale yazar.
Çalışma Akışı:
Program başlar ve sinyalleri dinlemeye başlar.
RSA özel anahtarı ve genel anahtarı oluşturulur ve dosyalara kaydedilir.
Program sonlandırılmak için kullanıcının Ctrl + C tuşuna basmasını bekler. Bu işlem yapılınca program güvenli bir şekilde kapanır.
Bu programın amacı, SSH anahtar çiftlerini oluşturmak ve kullanıcının manuel olarak programı durdurmasını beklemektir.
