# Stage 1: Build Angular app
FROM node:16 AS builder

WORKDIR /app

# Copy and install dependencies
COPY package*.json ./
RUN npm install

# Copy source code and build
COPY . .
RUN npm run build --prod

# Stage 2: Serve with NGINX
FROM nginx:alpine

# Copy built app from builder stage
COPY --from=builder /app/dist/my-app /usr/share/nginx/html

# Copy custom nginx config (optional)
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
