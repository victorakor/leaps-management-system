# LEAPS Backend — Deploy to Railway

## Prerequisites
- [Railway account](https://railway.app) (free tier available)
- Your code pushed to GitHub

## Step 1: Push to GitHub

```bash
cd leaps-backend
git init
git add .
git commit -m "Initial commit: LEAPS Management System"
git remote add origin https://github.com/YOUR_USERNAME/leaps-backend.git
git push -u origin main
```

## Step 2: Deploy on Railway

1. Go to [railway.app](https://railway.app) and sign in
2. Click **New Project** → **Deploy from GitHub repo**
3. Select your `leaps-backend` repo
4. Railway auto-detects the Dockerfile and builds it

## Step 3: Add PostgreSQL Database

1. In your Railway project, click **New** → **Database** → **PostgreSQL**
2. Railway provisions a free Postgres DB and provides connection details

## Step 4: Set Environment Variables

In Railway → your service → **Variables** tab, add:

```
DB_HOST         = (from Railway Postgres → Connect tab)
DB_PORT         = 5432
DB_USER         = (from Railway Postgres)
DB_PASSWORD     = (from Railway Postgres)
DB_NAME         = railway
JWT_SECRET      = (generate a strong random string, e.g. openssl rand -hex 32)
PORT            = 8080
```

Railway also provides a `DATABASE_URL` variable — you can use that instead by updating `config/db.go` to parse it.

## Step 5: Run Migrations

In Railway → your Postgres DB → **Query** tab, paste and run the contents of:
`migrations/001_init_schema.sql`

Or connect via psql:
```bash
psql $DATABASE_URL -f migrations/001_init_schema.sql
```

## Step 6: Test Your Live API

Railway gives you a URL like `https://leaps-backend-production.up.railway.app`

```bash
# Health check
curl https://YOUR_APP.railway.app/health

# Register a user
curl -X POST https://YOUR_APP.railway.app/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"full_name":"Admin User","email":"admin@school.ng","password":"secure123","role_id":"ROLE_UUID_HERE"}'

# Login
curl -X POST https://YOUR_APP.railway.app/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@school.ng","password":"secure123"}'
```

## Alternative: Deploy to Render.com

1. Go to [render.com](https://render.com)
2. **New** → **Web Service** → Connect GitHub repo
3. Build Command: `go build -o leaps-api ./main.go`
4. Start Command: `./leaps-api`
5. Add a **PostgreSQL** database from Render dashboard
6. Set environment variables (same as above)
7. Run migrations via Render's database shell

## Connecting the Frontend

Update `leaps-frontend/assets/js/api.js`:
```javascript
const API_BASE_URL = 'https://YOUR_APP.railway.app/api';
```

Then deploy the frontend separately (Netlify, Vercel, or Render static site).
