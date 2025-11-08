# Stage 1: Build binaries
FROM node:25-alpine3.21 AS builder

WORKDIR /app

# Install dependencies
COPY package.json package-lock.json ./

RUN npm ci

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Stage 2: Create the final image
FROM nginx:1.28-alpine

# Remove default nginx website
RUN rm -rf /usr/share/nginx/html/*

# Copy built assets from builder stage
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD [ "nginx", "-g", "daemon off;" ]