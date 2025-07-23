### Generate Swagger
```bash
swag init -g cmd/main.go
```

### Run Project
```bash
go run ./cmd/main.go
```

### Image Path
```bash
\\thdcweb06\BTLInterfaceProd\AX\ax_cust_qr_code
```

### Data Generate QR String
```bash
{
  "account_num": "142984",
  "amount": "0",
  "biller_id": "010550400104700",
  "company_bank": "88401"
}
```
### Generate QR Code
https://scanqr.org/
```bash
{
  "path": "./test-ja.png",
  "qr": "|010550400104700\n14298488401\n\n0"
}
```

### Check Exist File
```bash
./143024.png
```

### generate-qr-job
```bash
{
  "account_num": "142984",
  "amount": "0",
  "biller_id": "010550400104700",
  "company_bank": "88401",
  "file_path": "./142984.png"
}
```

### Job-sender
```bash
go run ./cmd/bulkqr/main.go
```

### Worker
```bash
go run ./cmd/worker/main.go
```