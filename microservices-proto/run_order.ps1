# Configurações
$GITHUB_USERNAME = "ruandg"
$GITHUB_EMAIL = "lucaspdn04@gmail.com"

$SERVICE_NAME = "order"
$RELEASE_VERSION = "v1.2.3"

# Configurar caminhos
$PROTOC_DIR = "$env:TEMP\protoc"
$PROTOC_BIN = "$PROTOC_DIR\bin\protoc.exe"
$GOBIN = "$(go env GOPATH)\bin"
$PROTOC_GEN_GO = "$GOBIN\protoc-gen-go.exe"
$PROTOC_GEN_GO_GRPC = "$GOBIN\protoc-gen-go-grpc.exe"

# Verifica se Go está instalado
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Error "Go não está instalado ou não está no PATH. Instale o Go antes de continuar."
    exit 1
}

# Baixar protoc, se necessário
if (-not (Test-Path $PROTOC_BIN)) {
    Write-Host "Baixando e instalando protoc..."
    $protocZipUrl = "https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protoc-21.12-win64.zip"
    $zipPath = "$env:TEMP\protoc.zip"
    Invoke-WebRequest -Uri $protocZipUrl -OutFile $zipPath
    Expand-Archive -Path $zipPath -DestinationPath $PROTOC_DIR -Force
    Remove-Item $zipPath
}

# Adicionar binários ao PATH da sessão
$env:PATH += ";$PROTOC_DIR\bin;$GOBIN"

# Instalar protoc-gen-go, se necessário
if (-not (Test-Path $PROTOC_GEN_GO)) {
    Write-Host "Instalando protoc-gen-go..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
}

# Instalar protoc-gen-go-grpc, se necessário
if (-not (Test-Path $PROTOC_GEN_GO_GRPC)) {
    Write-Host "Instalando protoc-gen-go-grpc..."
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
}

# Criar pasta de saída
Write-Host "`nGerando código Go a partir dos arquivos .proto"
New-Item -ItemType Directory -Force -Path "golang" | Out-Null

# Gerar os arquivos .pb.go e .grpc.pb.go
protoc --go_out=./golang `
  --go_opt=paths=source_relative `
  --go-grpc_out=./golang `
  --go-grpc_opt=paths=source_relative `
  .\$SERVICE_NAME\*.proto

if ($LASTEXITCODE -ne 0) {
    Write-Error "`n❌ Erro ao executar protoc. Verifique se os arquivos .proto estão corretos."
    exit 1
}

Write-Host "`nArquivos gerados:"
Get-ChildItem -Path ".\golang\$SERVICE_NAME"

# Inicializar módulo Go
Set-Location "golang\$SERVICE_NAME"
go mod init "github.com/$GITHUB_USERNAME/microservices-proto/golang/$SERVICE_NAME" 2>$null
go mod tidy 2>$null

Write-Host "`n✅ Processo finalizado com sucesso!"