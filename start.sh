#!/bin/sh
if [ ! -e "database/addresses.sqlite" ]; then
    echo "⏳ First startup, extracting database"
    gunzip database/addresses.sqlite.gz
    echo "⌛ Database extracted, starting server"
    postcode
else
    postcode
fi
