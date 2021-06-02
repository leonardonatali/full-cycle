FROM nginx:alpine

# Remove all nginx default configs 
RUN rm /etc/nginx/conf.d/default.conf

COPY nginx.conf /etc/nginx/conf.d

EXPOSE 80