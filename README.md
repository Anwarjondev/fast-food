# Fast Food API

A RESTful API for a fast food ordering system built with Go and Gin framework.

## Features

- User Authentication (Register, Login, Logout)
- Email Verification
- Password Reset
- Food Categories
- Food Items by Category
- Order Management
- Order Status Tracking
- **Swagger API Documentation**

## Prerequisites

- Go 1.24 or higher
- PostgreSQL
- SMTP Server (for email functionality)

## Environment Variables

Create a `.env` file in the root directory with the following variables (or set them in your deployment environment, e.g., Railway):

```env
DB_DNS=postgres://username:password@localhost:5432/dbname
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
EMAIL_SENDER=your-email@gmail.com
EMAIL_PASSWORD=your-app-specific-password
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/fast-food.git
cd fast-food
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

## API Documentation (Swagger)

After running the application, access the Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

This provides interactive documentation for all endpoints.

## API Endpoints

### Authentication
- `POST /register` - Register a new user
- `POST /login` - Login user
- `POST /logout` - Logout user
- `POST /confirm` - Confirm email with code
- `POST /resend-code` - Resend confirmation code
- `POST /forgot-password` - Request password reset
- `POST /reset-password` - Reset password

### Categories
- `GET /categories` - Get all categories
- `GET /categories/:id` - Get category by ID
- `GET /categories/:id/foods` - Get foods by category

### Orders
- `POST /orders` - Create new order
- `GET /orders/active` - Get active orders
- `GET /orders/completed` - Get completed orders
- `GET /orders/all` - Get all orders
- `PUT /orders/:order_id` - Cancel order

## Database Schema

The application uses PostgreSQL with the following main tables:
- users
- confirm
- category
- food
- orders
- order_detail

## Deployment

You can deploy this app to [Railway](https://railway.app/) or any platform that supports Go. Set environment variables in the platform's dashboard as needed.

## License

MIT License