# Pull official base image
FROM node:20-alpine

# Set the working directory in the container
WORKDIR /app

# Add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH
# Copy package.json and package-lock.json

# Copy local code to the container image.
COPY . ./

# Install application dependencies
RUN npm install

EXPOSE 3000
# Run the web service on container startup.
CMD ["npm", "run", "dev"]