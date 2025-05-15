# CO3109 Multidisciplinary Project - Angular

## Overview

This is the Angular-based web client for the Face Recognition Attendance System. The system allows businesses to manage employee attendance through facial recognition technology, providing a modern and secure way to track employee presence.

## Features

- **Face Registration**: Employees can register their faces through the system
- **Real-time Face Detection**: Integration with camera for real-time face detection
- **Multi-business Support**: Support for multiple businesses/organizations
- **Role-based Access**: Different access levels for managers and employees
- **Attendance Dashboard**: Visual representation of attendance data
- **Authentication System**: Secure login with email verification

## Technology Stack

- **Frontend Framework**: Angular 15.2.0
- **UI Components**: Angular Material 15.2.9
- **Authentication**: JWT-based authentication
- **Face Detection**: Integration with camera API
- **Styling**: Custom CSS with responsive design
- **HTTP Client**: Angular HttpClient for API communication

## Project Structure

```
src/
├── app/
│   ├── features/
│   │   ├── auth/              # Authentication components
│   │   ├── dashboard/         # Dashboard components
│   │   ├── face-register/     # Face registration module
│   │   └── select-business/   # Business selection components
│   ├── services/             # Application services
│   ├── params/              # API parameters/models
│   └── environments/        # Environment configurations
```

## Prerequisites

- Node.js (LTS version)
- npm package manager
- Angular CLI (version 15.2.11)
- Webcam/Camera access
- Modern web browser with WebRTC support

## Installation

1. Clone the repository:
   ```bash
   git clone [repository-url]
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Configure environment variables:
   - Create `src/environments/environment.ts` for development
   - Set API endpoints and other configuration

## Development Server

Run the development server:
```bash
npm start
```

The application will be available at `http://localhost:4200/`

## Building for Production

Build the project:
```bash
npm run build
```

The build artifacts will be stored in the `dist/` directory.

## Docker Support

Build and run with Docker:
```bash
docker build -t face-recognition-client .
docker run -p 80:80 face-recognition-client
```

## Key Components

### Face Registration
- Camera integration for face capture
- Face detection and validation
- Secure storage of face data

### Dashboard
- Real-time attendance tracking
- Monthly attendance reports
- Employee management interface

### Authentication
- Email-based registration
- OTP verification
- Secure login system
- Session management

## Security Features

- JWT-based authentication
- Secure face data transmission
- HTTPS enforcement
- Cross-Origin Resource Sharing (CORS) protection

## Browser Compatibility

- Chrome (latest)
- Firefox (latest)
- Edge (latest)
- Safari (latest)
