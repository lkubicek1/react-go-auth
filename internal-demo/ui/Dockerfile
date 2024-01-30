# Use the official Node.js image as a base
FROM node:latest

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json (or yarn.lock) files
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of your app's source code
COPY . .

# Vite uses port 3000 by default, but you can specify a different one
EXPOSE 3000

# Start the Vite development server
CMD ["npm", "run", "dev"]
