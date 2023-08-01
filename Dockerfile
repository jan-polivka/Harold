FROM node:18-alpine
WORKDIR /app
COPY . .
RUN npm run build
RUN npm install --omit=dev
CMD ["npm", "run", "prod"]
