package internal

//go:generate openssl req -x509 -nodes -days 24763 -newkey rsa:1024 -keyout key.pem -out cert.pem -config openssl.cnf

// LocalhostCert is a PEM-encoded TLS certificate with SAN IPs "127.0.0.1" and
// "[::1]", expiring Jan 1 17:37:04 2089 GMT.
//
// This has been generated using OpenSSL. The contents of the generated
// certificate can be read using:
//
// openssl x509 -in cert.pem -text -noout
var LocalhostCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDBzCCAnCgAwIBAgIJAIglicE/84ejMA0GCSqGSIb3DQEBCwUAMIGwMQswCQYD
VQQGEwJYWDEMMAoGA1UECAwDTi9BMQwwCgYDVQQHDANOL0ExIDAeBgNVBAoMF1Nl
bGYtc2lnbmVkIGNlcnRpZmljYXRlMSAwHgYDVQQLDBdTZWxmLXNpZ25lZCBjZXJ0
aWZpY2F0ZTEgMB4GA1UEAwwXU2VsZi1zaWduZWQgY2VydGlmaWNhdGUxHzAdBgkq
hkiG9w0BCQEWEHRlc3RAZXhhbXBsZS5jb20wIBcNMjEwMzE2MTczNzA0WhgPMjA4
OTAxMDExNzM3MDRaMIGwMQswCQYDVQQGEwJYWDEMMAoGA1UECAwDTi9BMQwwCgYD
VQQHDANOL0ExIDAeBgNVBAoMF1NlbGYtc2lnbmVkIGNlcnRpZmljYXRlMSAwHgYD
VQQLDBdTZWxmLXNpZ25lZCBjZXJ0aWZpY2F0ZTEgMB4GA1UEAwwXU2VsZi1zaWdu
ZWQgY2VydGlmaWNhdGUxHzAdBgkqhkiG9w0BCQEWEHRlc3RAZXhhbXBsZS5jb20w
gZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBANOP4kVt9rz4GcwIUCwTTLVYJKUb
yyBmTMmgH+RnrDh69TWa937Mt3N0YH3rkN8gDWQlUqv6lYMtFgzR5rakchHHsv9P
doO5jb3S+OOt6yBa70P+SdL4s1z6/xhdo7JyDbgZGrX5Z5UMb6bI4liy/swRZD4E
MUycEFaWmG9zsFIpAgMBAAGjJTAjMCEGA1UdEQQaMBiHBH8AAAGHEAAAAAAAAAAA
AAAAAAAAAAEwDQYJKoZIhvcNAQELBQADgYEAOq9WsqqDaszwBVPstAPlB8ld5aLS
9u9mlZrmUzP5dVpqlVSqnqBBi8WAhdDccYZR8Bx21bGqPlO84xxA4nl4zfm0EgDi
sqjm801oprpdhXY6hKiNngc7EGwEIhw4oAjF21W2dtth6rrO15o6ZTKIeifElHit
0f0QOHfqh0tB04U=
-----END CERTIFICATE-----`)

// LocalhostKey is a PEM-encoded TLS private key.
var LocalhostKey = []byte(`-----BEGIN PRIVATE KEY-----
MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBANOP4kVt9rz4GcwI
UCwTTLVYJKUbyyBmTMmgH+RnrDh69TWa937Mt3N0YH3rkN8gDWQlUqv6lYMtFgzR
5rakchHHsv9PdoO5jb3S+OOt6yBa70P+SdL4s1z6/xhdo7JyDbgZGrX5Z5UMb6bI
4liy/swRZD4EMUycEFaWmG9zsFIpAgMBAAECgYEAmK/4i0lg0WNyItpBGn5XV24l
DHCxulF6y+3f9pWIKz86qBSO2RngsNfmKBsSdKaKY6O4NdzleUXZ6yi2jUzD8Zcy
4Yhm1U/uIdOrK8gNVVFaRX6scA3TQKKoBg1uI0f9G2RXnKea8s2ow2/cHpI6nRnh
wJ5lTcWfHJPqX3wPAYECQQDqmjW+ejWDwirdJ+1F3n5h0UVI5kFFiUkBJU0PK203
pPJfIqt7f6JIWM0FywBddCfNSP3BJGGmPdKomB0U5iNxAkEA5tuzBxZJVbnP3IIM
Mgni+jqbCNZbNkhu/bIFKvPoKsVTzecVgp9UB0wNxCeH+NT7+fECl63F0ifH5e6B
6QZOOQJANdXDZX1n0F97NJrX8QOYntvF+W+VJN7XGOM1ZrjBbFZ2o+wxy7pDRDTU
f3LRF6DzDUGlsW+m1N40/ClD6yrQkQJBAKfJZgTwnbxAoGKT4mk75kwX7DYwFXsO
ihy5FihxvXvYj2BHY8rcIiZrkFXQpXKk2b5+/HNhSTXE0/S5tJG4k4ECQQCmQ3vN
SopEnF/JSNv2R4jsxELPnXECISNa6OJPKD5OqFKpMMFhaAzWk2uhy8ESJSTKKxOq
U5VH3aQgoKNkYZ/q
-----END PRIVATE KEY-----`)
