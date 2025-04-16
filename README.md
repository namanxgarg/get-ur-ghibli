# Get-Ur-Ghibli

A microservices-based web application where users can upload a photo and generate “Studio Ghibli–style” images. Includes:

- **Auth Service**: Sign-up, login, JWT token management, tracks free usage.  
- **Upload Service**: Manages image uploads (to local or S3).  
- **Ghibli Service**: Generates Ghibli-style images (mock or AI-based).  
- **Order Service**: Handles orders and payments (mock or real payment gateway).  
- **Gateway**: Single entry point, routes requests, enforces JWT.  
- **Frontend**: Simple React app to showcase the flow (upload → generate → pay/order).

---

## Table of Contents
1. [Features](#features)  
2. [Project Architecture](#project-architecture)  
3. [Tech Stack](#tech-stack)  
4. [Requirements](#requirements)  
5. [Installation & Setup](#installation--setup)  
6. [Usage Flow](#usage-flow)  
7. [Services Overview](#services-overview)  
8. [Frontend](#frontend)  
9. [Docker Compose](#docker-compose)  
10. [Advanced Notes](#advanced-notes)  
11. [Contributing](#contributing)  
12. [License](#license)

---

## Features

1. **User Auth**  
   - Sign up, login with JWT tokens.  
   - Tracks whether the user has used the free Ghibli image.  

2. **Image Upload**  
   - Upload a photo (local disk or S3).  

3. **Ghibli Image Generation**  
   - 1 free Ghibli image if user hasn’t used the free option yet.  
   - 10 paid Ghibli images after a small payment (e.g. Rs.100).  

4. **3D Model Orders**  
   - Users can order a 3D Ghibli model (Rs.2000).  
   - Order tracking with statuses.  

5. **Microservices Architecture**  
   - Auth, Upload, Ghibli, Order, plus a Gateway.  
   - Communication via REST calls.  

6. **React Frontend**  
   - Demonstrates the end-to-end flow (upload → generate → order/payment).  

---


- **Gateway**: Routes requests to each service, verifies JWT.  
- **Auth Service**: User database, sign-up/login, tracks free usage.  
- **Upload Service**: Saves images (local or S3).  
- **Ghibli Service**: Generates or mocks Ghibli-style images.  
- **Order Service**: Stores orders, manages payment flow (mock or real).  
- **Frontend**: React single-page application.

---

## Tech Stack

- **Go** (1.20+) for all microservices.  
- **React** for the frontend.  
- **Docker** + **Docker Compose** for local orchestration.  
- (Optional) **Postgres** for Auth/Order data.  
- (Optional) **AWS S3** for storing images.  

---

## Requirements

- **Go** 1.20+  
- **Node** 16+ (for the React frontend)  
- **Docker** & **Docker Compose** (for local runs)  
- (Optional) **Postgres** if using the default database approach  
- (Optional) **AWS account** for S3 image storage  
