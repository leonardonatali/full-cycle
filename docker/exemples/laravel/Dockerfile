FROM php:7.4-cli

WORKDIR /var/www

# Install and update system dependencies and PHP dependencies
RUN \
    apt update && \
    apt install libzip-dev -y && \
    docker-php-ext-install zip

# Install composer
RUN \
    php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" && \
    php composer-setup.php && \
    php -r "unlink('composer-setup.php');"
    
# Install Laravel
RUN \
    php composer.phar create-project laravel/laravel laravel

ENTRYPOINT [ "php", "laravel/artisan", "serve" ]

EXPOSE 80

CMD [ "--host=0.0.0.0", "--port=80"]