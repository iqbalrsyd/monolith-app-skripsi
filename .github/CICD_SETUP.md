# GitHub Actions CI/CD Setup

## Overview
Project ini menggunakan 3 workflow GitHub Actions untuk CI/CD:

1. **Go CI** - Build, test, dan lint kode Go
2. **Docker Build and Push** - Build Docker image dan push ke Docker Hub
3. **Deploy to Production** - Deploy aplikasi ke server production

## Workflows

### 1. Go CI ([.github/workflows/go.yml](.github/workflows/go.yml))
Workflow ini otomatis berjalan saat:
- Push ke branch `main`
- Pull request ke branch `main`

**Fitur:**
- ✅ Setup Go 1.21
- ✅ Caching dependencies untuk build lebih cepat
- ✅ Verify dan download dependencies
- ✅ Run `go vet` untuk static analysis
- ✅ Run `golangci-lint` untuk code quality
- ✅ Build aplikasi
- ✅ Run tests dengan race detector dan code coverage
- ✅ Upload coverage ke Codecov (opsional)

### 2. Docker Build and Push ([.github/workflows/docker-image.yml](.github/workflows/docker-image.yml))
Workflow ini otomatis berjalan saat:
- Push ke branch `main`
- Push tag dengan format `v*` (e.g., v1.0.0)
- Pull request ke branch `main` (hanya build, tidak push)

**Fitur:**
- ✅ Setup Docker Buildx untuk build multi-platform
- ✅ Login ke Docker Hub (hanya untuk push, bukan PR)
- ✅ Generate metadata (tags, labels) otomatis
- ✅ Build dan push Docker image dengan multiple tags
- ✅ Cache layer Docker untuk build lebih cepat
- ✅ Scan vulnerability dengan Trivy
- ✅ Upload hasil scan ke GitHub Security tab

**Tags yang di-generate:**
- `latest` - untuk push ke main branch
- `main` - untuk push ke main branch
- `v1.0.0` - untuk tag semver
- `1.0` - untuk tag major.minor
- `main-sha-abc123` - untuk setiap commit

### 3. Deploy to Production ([.github/workflows/deploy.yml](.github/workflows/deploy.yml))
Workflow ini otomatis berjalan saat:
- Workflow "Docker Build and Push" selesai sukses di branch `main`
- Manual trigger via workflow_dispatch

**Fitur:**
- ✅ Deploy via SSH ke server
- ✅ Pull Docker image terbaru
- ✅ Stop dan remove container lama
- ✅ Run container baru dengan auto-restart
- ✅ Clean up old images
- ✅ Health check setelah deployment
- ✅ Support multiple environments (production/staging)

## Setup Secrets

Anda perlu menambahkan secrets berikut di GitHub repository:
**Settings → Secrets and variables → Actions → New repository secret**

### Required Secrets untuk Docker:
```
DOCKER_USERNAME     : Username Docker Hub Anda
DOCKER_PASSWORD     : Password atau Access Token Docker Hub
```

### Required Secrets untuk Deployment:
```
DEPLOY_HOST        : IP atau hostname server (e.g., 192.168.1.100)
DEPLOY_USER        : Username SSH untuk login ke server (e.g., ubuntu)
DEPLOY_KEY         : Private SSH key untuk authentication
DEPLOY_PORT        : Port SSH (default: 22, opsional)
DEPLOY_PATH        : Path direktori di server (default: /opt/monolith-app, opsional)
```

### Optional Secrets:
```
CODECOV_TOKEN      : Token untuk upload coverage ke Codecov (opsional)
```

## Cara Setup Secrets

### 1. Docker Hub Credentials

1. Login ke [Docker Hub](https://hub.docker.com)
2. Go to **Account Settings → Security → Access Tokens**
3. Generate new token dengan nama "GitHub Actions"
4. Copy token tersebut
5. Di GitHub repo: **Settings → Secrets → New repository secret**
   - Name: `DOCKER_USERNAME` → Value: username Docker Hub
   - Name: `DOCKER_PASSWORD` → Value: access token yang di-copy

### 2. SSH Key untuk Deployment

**Di server deployment:**
```bash
# Jika belum punya SSH key di GitHub Actions
ssh-keygen -t ed25519 -C "github-actions" -f ~/.ssh/github_actions
cat ~/.ssh/github_actions.pub >> ~/.ssh/authorized_keys
cat ~/.ssh/github_actions  # Copy private key ini
```

**Di GitHub repository:**
- Name: `DEPLOY_KEY` → Value: paste private key lengkap (termasuk BEGIN dan END)
- Name: `DEPLOY_HOST` → Value: IP server (e.g., 192.168.1.100)
- Name: `DEPLOY_USER` → Value: username SSH (e.g., ubuntu)

### 3. Environment Setup (Optional tapi Recommended)

Untuk protection rules dan approval:
1. **Settings → Environments → New environment** → Nama: `production`
2. Configure environment:
   - ✅ Required reviewers (pilih tim/orang yang harus approve)
   - ✅ Wait timer (delay sebelum deploy)
   - ⚠️ Deployment branches → Selected branches → `main`

## Testing Workflows Locally

### Test Go Build & Test:
```bash
go mod verify
go mod download
go vet ./...
go build -v ./...
go test -v -race -coverprofile=coverage.out ./...
```

### Test Docker Build:
```bash
docker build -t monolith-app:test .
docker run -p 8080:8080 monolith-app:test
```

### Test Golangci-lint:
```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run lint
golangci-lint run --timeout=5m
```

## Workflow Triggers

| Event | Go CI | Docker | Deploy |
|-------|-------|--------|--------|
| Push to main | ✅ | ✅ | ✅ (after Docker) |
| Pull Request | ✅ | ✅ (build only) | ❌ |
| Tag push (v*) | ❌ | ✅ | ❌ |
| Manual | ❌ | ❌ | ✅ |

## Monitoring & Debugging

### View Workflow Runs:
1. Go to repository → **Actions** tab
2. Click workflow name di sidebar
3. Click specific run untuk detail logs

### Common Issues:

**1. Docker push failed - authentication required**
- ✅ Check `DOCKER_USERNAME` dan `DOCKER_PASSWORD` secrets
- ✅ Pastikan Docker Hub token masih valid

**2. Deployment failed - SSH connection**
- ✅ Check `DEPLOY_KEY` format (harus include BEGIN/END)
- ✅ Check `DEPLOY_HOST` accessible dari GitHub runners
- ✅ Check firewall rules di server (allow SSH dari GitHub IPs)

**3. Golangci-lint timeout**
- ✅ Adjust timeout di `.golangci.yml`
- ✅ Disable linters yang tidak diperlukan

**4. Health check failed after deployment**
- ✅ Check container logs: `docker logs monolith-app`
- ✅ Check port 8080 terbuka di server
- ✅ Pastikan `/health` endpoint exists

## Best Practices

1. **Branch Protection**: Enable di `main` branch
   - Require PR reviews
   - Require status checks (Go CI, Docker Build)
   - No direct push to main

2. **Semantic Versioning**: Use tags untuk releases
   ```bash
   git tag -a v1.0.0 -m "Release version 1.0.0"
   git push origin v1.0.0
   ```

3. **Manual Deployment**: Untuk deploy ke staging
   - Go to Actions → Deploy to Production
   - Click "Run workflow"
   - Select branch dan environment

4. **Rollback**: Jika deployment bermasalah
   ```bash
   # SSH ke server
   ssh user@server
   
   # Pull versi sebelumnya
   docker pull username/monolith-app:v1.0.0
   
   # Restart dengan versi lama
   docker stop monolith-app && docker rm monolith-app
   docker run -d --name monolith-app -p 8080:8080 username/monolith-app:v1.0.0
   ```

## Next Steps

- [ ] Setup Codecov untuk code coverage tracking
- [ ] Add integration tests di CI
- [ ] Setup staging environment
- [ ] Add database migrations di deployment
- [ ] Setup monitoring (Prometheus/Grafana)
- [ ] Add Slack/Discord notifications untuk deployment
- [ ] Setup blue-green deployment atau canary releases
