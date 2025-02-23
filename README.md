# Point of Sale (POS) System

## Overview
This is a simple Point of Sale (POS) system built with Go. The system allows users to manage inventory, add products through a web interface, generate QR codes, and facilitate sales transactions. Vendors can scan QR codes from their mobile devices to add products to a customer's cart without needing a dedicated mobile app.

## Intended Features
- **User Authentication**: Powered by Ory Kratos for secure user management (tentative).
- **Inventory Management**: Add, update, and delete products from an inventory.
- **QR Code Generation**: Products have QR codes for easy addition to carts.
- **Mobile-Friendly Interface**: Vendors can scan QR codes and process sales from their mobile browsers.
- **Session Management**: Maintain active sessions for vendors to keep track of sales.
- **Pricing & Checkout**: Calculate total prices and facilitate transactions.

## Technologies Used
- **Backend**: Go (Golang)
- **Authentication**: Ory Kratos (tentative)
- **Database**: PostgreSQL 
- **Frontend**: HTML, CSS, JavaScript
- **Containerization**: Docker && docker compose
- **Deployment**: Kubernetes (tentative)

## Installation
### Prerequisites
- Go installed on your machine
- docker

### Steps
1. **Clone the repository**
   ```sh
   git clone https://github.com/OscarMoya/tpv.git
   cd tpv
   ```

2. **Install dependencies**
   ```sh

   cd backend && go mod tidy
   ```

## Usage
- **Add Products**: Use the web form to add new products.
- **Scan & Checkout**: Vendors can scan product QRs and process sales.

## License
This project is licensed under the MIT License.

