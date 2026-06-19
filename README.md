# 🏍️ Dinz RentBike

Rental motor & mobil — booking cepat, bayar mudah.

**Deploy:** [dinz-rentbike.up.railway.app](https://dinz-rentbike.up.railway.app/)

---

## API Endpoints

### Auth
```
POST /api/v1/auth/register
POST /api/v1/auth/login
```

### User
```
GET    /api/v1/users/me
PUT    /api/v1/users/update
PATCH  /api/v1/users/change-password
```

### Vehicles
```
GET  /api/v1/vehicles
GET  /api/v1/vehicles/:id
```

### Rentals
```
GET   /api/v1/rentals
GET   /api/v1/rentals/:id
POST  /api/v1/rentals/create      → auto-create payment
POST  /api/v1/rentals/cancel
```

### Payments
```
GET   /api/v1/payments
GET   /api/v1/payments/:id
POST  /api/v1/payments/cancel
```

### Reviews
```
GET   /api/v1/reviews
POST  /api/v1/reviews/create
POST  /api/v1/reviews/edit
POST  /api/v1/reviews/delete
```

### Admin
```
GET     /api/v1/admin/vehicles
POST    /api/v1/admin/vehicles
PUT     /api/v1/admin/vehicles/:id
DELETE  /api/v1/admin/vehicles/:id
GET     /api/v1/admin/rentals
GET     /api/v1/admin/rentals/:id
PATCH   /api/v1/admin/rentals/:id/status
GET     /api/v1/admin/payments
GET     /api/v1/admin/payments/:id
```

### Webhook
```
POST /webhook    → Xendit payment session callback
```

---

## Environment

```env
# Server
APP_NAME=DinzRentbike
APP_HOST=localhost
APP_PORT=8080

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=password
DB_NAME=rentbike_db
DB_SCHEMA=public

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# JWT
JWT_SECRET=your-jwt-secret-key-here

# Xendit
XENDIT_BASE_URL=https://api.xendit.co
XENDIT_PUBLIC_KEY=
XENDIT_SECRET_KEY=
XENDIT_WEBHOOK_TOKEN=

# Mailjet
MAILJET_BASE_URL=https://api.mailjet.com/v3.1/send
MAILJET_API_KEY=
MAILJET_SECRET_KEY=
```

---

## Run

```bash
go run cmd/app/main.go
```

## Test

```bash
go test ./tests/ -v
```

## Mock Generate

```bash
mockery
```

---

## Flow

```
Register/Login → Browse Vehicles → Create Rental (auto payment) 
→ Pay via Xendit → Webhook updates status → Email notification → Review
```
