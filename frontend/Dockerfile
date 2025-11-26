FROM node:22-alpine AS build

WORKDIR /app

# Copy only package files first for caching
COPY package*.json ./

# Install dependencies (cached if package.json didn't change)
RUN npm install

# Copy the rest of the app
COPY . .

EXPOSE 9090
CMD ["npm", "run", "dev"]