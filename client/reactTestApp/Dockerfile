# Step 1: Use an official Node.js runtime as a parent image
FROM node:22-slim AS node
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the package.json and package-lock.json (or yarn.lock) files to install dependencies
COPY package*.json ./

# Step 4: Install dependencies
RUN pnpm install

# Step 5: Copy the entire application source code into the container
COPY . .

# Step 6: Build the Vite app (this will run the build script in package.json)
RUN pnpm run build

# Step 7: Serve the app using Nginx
FROM nginx:alpine

# Step 8: Copy the built files from the previous build stage to the Nginx server's default directory
COPY --from=node /app/dist /usr/share/nginx/html

# Step 9: Copy custom Nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Step 10: Create an img directory and set permissions
RUN mkdir -p /usr/share/nginx/html/img && \
  chmod -R 777 /usr/share/nginx/html/img

# Step 11: Expose the port that Nginx is running on (default 80)
EXPOSE 80

# Step 12: Start Nginx in the background
CMD ["nginx", "-g", "daemon off;"]
