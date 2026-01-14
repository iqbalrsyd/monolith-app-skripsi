# Monolith App - Go Web Application

Aplikasi monolithic Go web service untuk penelitian DevOps comparison dengan arsitektur microservices.

## 📁 Struktur Project

```
monolith-app/
├── main.go              # Entry point aplikasi
├── go.mod               # Go modules
├── go.sum               # Go dependencies
├── Dockerfile           # Container build configuration
├── handlers/            # HTTP handlers
│   ├── user_handler.go  # User management
│   ├── product_handler.go # Product management
│   └── health_handler.go # Health checks
├── models/              # Data models
├── routes/              # Route configuration
├── middleware/          # HTTP middleware
├── utils/               # Utility functions
└── README.md            # This file
```

## 🚀 Fitur

- **User Management**: CRUD operations untuk users
- **Product Management**: CRUD operations untuk products
- **Health Checks**: Monitoring endpoints
- **Database Integration**: Connection ke database
- **API Documentation**: RESTful API endpoints
- **Docker Support**: Containerized deployment
- **Graceful Shutdown**: Clean shutdown handling

## 🔧 Teknologi Stack

- **Backend**: Go 1.22+
- **Web Framework**: Gin HTTP framework
- **Database**: PostgreSQL/MySQL (configurable)
- **Container**: Docker
- **Orchestration**: Kubernetes

## 📋 API Endpoints

### Users
- `GET /api/users` - Get all users
- `GET /api/users/{id}` - Get user by ID
- `POST /api/users` - Create new user
- `PUT /api/users/{id}` - Update user
- `DELETE /api/users/{id}` - Delete user

### Products
- `GET /api/products` - Get all products
- `GET /api/products/{id}` - Get product by ID
- `POST /api/products` - Create new product
- `PUT /api/products/{id}` - Update product
- `DELETE /api/products/{id}` - Delete product

### Health
- `GET /health` - Application health check
- `GET /ready` - Readiness probe
- `GET /metrics` - Application metrics

## 🏗️ Cara Build & Deploy

### Local Development
```bash
# Install dependencies
go mod download

# Run aplikasi
go run main.go

# Build binary
go build -o monolith-app main.go
```

### Docker Build
```bash
# Build Docker image
docker build -t monolith-app .

# Run container
docker run -p 8080:8080 monolith-app
```

### Kubernetes Deployment
```bash
# Apply manifests
kubectl apply -f ../infra-app/k8s/monolith/

# Check deployment
kubectl get pods -l app=monolith-app
```

## 📊 Performance Characteristics

- **Startup Time**: ~2-3 seconds
- **Memory Usage**: ~50-100MB baseline
- **CPU Usage**: Low for moderate loads
- **Response Time**: <100ms for simple operations
- **Throughput**: ~1000 RPS per instance

## 🔍 Monitoring & Observability

### Health Checks
- **Liveness Probe**: `/health`
- **Readiness Probe**: `/ready`
- **Metrics**: `/metrics` (Prometheus format)

### Logging
- **Structured Logging**: JSON format
- **Log Levels**: DEBUG, INFO, WARN, ERROR
- **Request Tracing**: Request ID tracking

## 🧪 Testing

### Unit Tests
```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...
```

### Integration Tests
```bash
# Run integration tests
go test -tags=integration ./...
```

### Load Testing
```bash
# Using k6 (configured in infra-app)
k6 run ../infra-app/k6-tests/monolith-load-test.js
```

## ⚙️ Configuration

### Environment Variables
```bash
# Server Configuration
PORT=8080
HOST=0.0.0.0

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=monolith_db
DB_USER=user
DB_PASSWORD=password

# Application Configuration
LOG_LEVEL=info
METRICS_ENABLED=true
```

### Config Files
- `config.json` - Application configuration
- `docker-compose.yml` - Local development setup

## 🔄 CI/CD Integration

### GitHub Actions Workflows
Repository ini menggunakan GitHub Actions untuk CI/CD automation:

1. **Go CI** (`.github/workflows/go.yml`)
   - Build dan test kode Go
   - Run linting dengan golangci-lint
   - Generate code coverage reports

2. **Docker Build & Push** (`.github/workflows/docker-image.yml`)
   - Build Docker image dengan multi-stage build
   - Push image ke Docker Hub
   - Security scanning dengan Trivy
   - Support multiple tags (latest, version, SHA)

3. **Deploy to Production** (`.github/workflows/deploy.yml`)
   - Auto-deploy setelah Docker build sukses
   - Deploy via SSH ke server
   - Health check verification
   - Support manual deployment dengan workflow_dispatch

📖 **Dokumentasi lengkap**: Lihat [`.github/CICD_SETUP.md`](.github/CICD_SETUP.md) untuk setup dan konfigurasi detail.

### Setup Requirements
- Docker Hub account dengan credentials di GitHub Secrets
- SSH access ke deployment server
- Environment variables configured di GitHub

### Workflow Triggers
- **Push ke main**: Trigger semua workflows
- **Pull Request**: Run tests dan build (no deployment)
- **Tag push (v*)**: Create versioned releases
- **Manual**: Deploy via GitHub Actions UI

### GitOps Workflow
- **Source**: GitHub repository
- **Build**: GitHub Actions
- **Deploy**: SSH ke server atau Kubernetes
- **Monitor**: Prometheus/Grafana

## 🚀 Deployment Scenarios

### Development
- Single instance deployment
- Local database
- Basic monitoring

### Staging
- Multiple instances
- Staging database
- Performance testing

### Production
- Auto-scaling enabled
- Production database
- Full monitoring stack
- Load balancing

## 📈 Research Metrics

Aplikasi ini dirancang untuk penelitian DevOps dengan metrik:
- **Deployment Time**: Build dan deploy duration
- **Reliability**: Uptime dan error rates
- **Performance**: Response time dan throughput
- **Scalability**: Auto-scaling behavior
- **Resource Usage**: CPU dan memory consumption

## 🔗 Related Projects

- **Microservices App**: [GitHub repo link]
- **Infrastructure**: [GitHub repo link]
- **CI/CD Pipeline**: Jenkins configuration in infra-app

## 📝 Development Notes

### Performance Optimizations
- Connection pooling untuk database
- Request caching dimana perlu
- Efficient query patterns
- Memory management

### Security Considerations
- Input validation
- SQL injection prevention
- Rate limiting
- Authentication middleware