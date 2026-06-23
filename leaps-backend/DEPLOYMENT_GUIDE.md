# LEAPS Deployment Guide

## Prerequisites

- Docker 20.10+
- Docker Compose 2.0+
- Git
- 2GB RAM minimum
- 10GB disk space

## Quick Start

### 1. Clone the Repository

```bash
git clone <repository-url>
cd leaps-backend
```

### 2. Configure Environment

```bash
cp .env.example .env
```

Edit `.env` with your configuration:

```env
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your-secure-password
DB_NAME=leaps_db
PORT=8080
```

### 3. Deploy with Docker Compose

```bash
docker-compose up -d
```

This will:
- Build the API Docker image
- Start PostgreSQL database
- Start the API server
- Start pgAdmin for database management

### 4. Verify Deployment

```bash
# Check if API is running
curl http://localhost:8080/health

# View logs
docker-compose logs -f api
```

## Services

### API Server
- **URL:** http://localhost:8080
- **Health Check:** http://localhost:8080/health
- **Container:** leaps-api

### PostgreSQL Database
- **Host:** localhost
- **Port:** 5432
- **Container:** leaps-postgres

### pgAdmin
- **URL:** http://localhost:5050
- **Email:** admin@leaps.local
- **Password:** admin
- **Container:** leaps-pgadmin

## Database Management

### Connect to Database

```bash
# Using psql
psql -h localhost -U postgres -d leaps_db

# Using Docker
docker-compose exec postgres psql -U postgres -d leaps_db
```

### Backup Database

```bash
docker-compose exec postgres pg_dump -U postgres leaps_db > backup.sql
```

### Restore Database

```bash
docker-compose exec -T postgres psql -U postgres leaps_db < backup.sql
```

## Common Commands

### Start Services

```bash
docker-compose up -d
```

### Stop Services

```bash
docker-compose down
```

### View Logs

```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f api
docker-compose logs -f postgres
```

### Restart Services

```bash
docker-compose restart
```

### Rebuild Images

```bash
docker-compose build --no-cache
```

### Remove All Data

```bash
docker-compose down -v
```

## Troubleshooting

### API Won't Start

1. Check logs:
```bash
docker-compose logs api
```

2. Verify database connection:
```bash
docker-compose logs postgres
```

3. Check environment variables:
```bash
docker-compose config
```

### Database Connection Error

1. Ensure PostgreSQL is running:
```bash
docker-compose ps
```

2. Check database logs:
```bash
docker-compose logs postgres
```

3. Verify credentials in .env file

### Port Already in Use

Change ports in docker-compose.yml:

```yaml
ports:
  - "8081:8080"  # API on 8081
  - "5433:5432"  # Database on 5433
```

## Production Deployment

### 1. Use Environment Variables

Never commit .env file. Use environment variables:

```bash
export DB_PASSWORD=secure-password
docker-compose up -d
```

### 2. Use Secrets Management

For production, use Docker Secrets or external secret management:

```bash
echo "secure-password" | docker secret create db_password -
```

### 3. Enable HTTPS

Use a reverse proxy (Nginx, Traefik) with SSL certificates:

```yaml
services:
  nginx:
    image: nginx:latest
    ports:
      - "443:443"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf
      - ./certs:/etc/nginx/certs
```

### 4. Set Resource Limits

```yaml
services:
  api:
    deploy:
      resources:
        limits:
          cpus: '1'
          memory: 512M
        reservations:
          cpus: '0.5'
          memory: 256M
```

### 5. Enable Logging

```yaml
services:
  api:
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "3"
```

## Monitoring

### Health Checks

The API includes health checks:

```bash
curl http://localhost:8080/health
```

Response:
```json
{
  "status": "ok",
  "message": "LEAPS API is running"
}
```

### Database Monitoring

Use pgAdmin:
1. Open http://localhost:5050
2. Login with admin@leaps.local / admin
3. Add server connection to postgres:5432

### Log Monitoring

```bash
# Real-time logs
docker-compose logs -f

# Last 100 lines
docker-compose logs --tail=100

# Specific time range
docker-compose logs --since 2024-01-01 --until 2024-01-02
```

## Scaling

### Horizontal Scaling

Run multiple API instances:

```yaml
services:
  api:
    deploy:
      replicas: 3
```

### Load Balancing

Use Nginx or HAProxy:

```nginx
upstream api {
    server api:8080;
    server api:8081;
    server api:8082;
}

server {
    listen 80;
    location / {
        proxy_pass http://api;
    }
}
```

## Backup and Recovery

### Automated Backups

Create a backup script:

```bash
#!/bin/bash
BACKUP_DIR="/backups"
DATE=$(date +%Y%m%d_%H%M%S)

docker-compose exec -T postgres pg_dump -U postgres leaps_db > \
  $BACKUP_DIR/leaps_db_$DATE.sql

# Keep only last 7 days
find $BACKUP_DIR -name "leaps_db_*.sql" -mtime +7 -delete
```

Schedule with cron:

```bash
0 2 * * * /path/to/backup.sh
```

## Security Checklist

- [ ] Change default passwords
- [ ] Use strong database password
- [ ] Enable HTTPS/SSL
- [ ] Set up firewall rules
- [ ] Enable database backups
- [ ] Monitor logs regularly
- [ ] Keep Docker images updated
- [ ] Use secrets management
- [ ] Enable authentication
- [ ] Set up rate limiting

## Support

For issues or questions:
1. Check logs: `docker-compose logs`
2. Review this guide
3. Check API documentation
4. Contact support team

