#!/bin/sh

if [ ! -f /usr/local/bin/rustdesk-api-server-pro ]; then
    ln -s /app/rustdesk-api-server-pro /usr/local/bin/rustdesk-api-server-pro
fi

if [ ! -f /app/server.db ]; then
    /app/rustdesk-api-server-pro sync
fi

if [ ! -f /app/.init.lock ] && [ -n "$ADMIN_USER" ] && [ -n "$ADMIN_PASS" ]; then
    /app/rustdesk-api-server-pro user add $ADMIN_USER $ADMIN_PASS --admin
    touch /app/.init.lock
fi

/app/rustdesk-api-server-pro start