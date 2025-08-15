package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	// Check if running on Render (PORT env var is set)
	port := os.Getenv("PORT")
	if port == "" {
		// Local dev mode — HTTPS
		if _, err := os.Stat("cert.pem"); os.IsNotExist(err) {
			log.Println("cert.pem not found — generating self-signed cert...")
			if err := generateSelfSignedCert([]string{"localhost", "127.0.0.1", "::1"}); err != nil {
				log.Fatalf("failed to generate cert: %v", err)
			}
			log.Println("Generated cert.pem and key.pem")
		}
		setupStaticFiles()
		addr := ":8443"
		log.Printf("Starting LOCAL HTTPS server at https://localhost%s\n", addr)
		if err := http.ListenAndServeTLS(addr, "cert.pem", "key.pem", nil); err != nil {
			log.Fatalf("ListenAndServeTLS: %v", err)
		}
	} else {
		// Render/Production mode — HTTP only
		setupStaticFiles()
		addr := ":" + port
		log.Printf("Starting HTTP server at http://0.0.0.0%s\n", addr)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}
}

func setupStaticFiles() {
	if _, err := os.Stat("static"); os.IsNotExist(err) {
		if err := os.Mkdir("static", 0755); err != nil {
			log.Fatalf("mkdir static: %v", err)
		}
		index := `<!doctype html>
<html>
<head><meta charset="utf-8"><title>Go Web</title></head>
<body>
  <h1>Hello from Go!</h1>
  <p>Served from <code>static/index.html</code></p>
</body>
</html>`
		_ = os.WriteFile("static/index.html", []byte(index), 0644)
	}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
}

func generateSelfSignedCert(hosts []string) error {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Local Dev"},
		},
		NotBefore:             time.Now().Add(-1 * time.Hour),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return err
	}

	certOut, err := os.Create("cert.pem")
	if err != nil {
		return err
	}
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		certOut.Close()
		return err
	}
	certOut.Close()

	keyOut, err := os.OpenFile("key.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	privBytes := x509.MarshalPKCS1PrivateKey(priv)
	if err := pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privBytes}); err != nil {
		keyOut.Close()
		return err
	}
	keyOut.Close()
	return nil
}
