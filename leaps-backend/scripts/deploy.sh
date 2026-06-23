#!/bin/bash

# LEAPS Deployment Script
# This script handles deployment of the LEAPS application

set -e

echo "🚀 Starting LEAPS Deployment..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo -e "${RED}❌ Docker is not installed${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Docker is installed${NC}"

# Check if Docker Compose is installed
if ! command -v docker-compose &> /dev/null; then
    echo -e "${RED}❌ Docker Compose is not installed${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Docker Compose is installed${NC}"

# Load environment variables
if [ -f .env ]; then
    export $(cat .env | grep -v '#' | xargs)
    echo -e "${GREEN}✅ Environment variables loaded${NC}"
else
    echo -e "${YELLOW}⚠️  .env file not found, using defaults${NC}"
fi

# Build Docker images
echo -e "${YELLOW}📦 Building Docker images...${NC}"
docker-compose build

# Start services
echo -e "${YELLOW}🔧 Starting services...${NC}"
docker-compose up -d

# Wait for database to be ready
echo -e "${YELLOW}⏳ Waiting for database to be ready...${NC}"
sleep 10

# Check if API is running
echo -e "${YELLOW}🔍 Checking API health...${NC}"
for i in {1..30}; do
    if curl -f http://localhost:8080/health > /dev/null 2>&1; then
        echo -e "${GREEN}✅ API is healthy${NC}"
        break
    fi
    echo "Attempt $i/30..."
    sleep 2
done

# Display service information
echo ""
echo -e "${GREEN}🎉 Deployment Complete!${NC}"
echo ""
echo "Services:"
echo "  API Server:    http://localhost:8080"
echo "  pgAdmin:       http://localhost:5050"
echo "  Database:      localhost:5432"
echo ""
echo "Default Credentials:"
echo "  pgAdmin Email:    admin@leaps.local"
echo "  pgAdmin Password: admin"
echo ""
echo "Useful Commands:"
echo "  View logs:       docker-compose logs -f api"
echo "  Stop services:   docker-compose down"
echo "  Restart services: docker-compose restart"
echo ""
