#!/bin/bash

# This script generates 10 unique messages and pipes them into the Kafka console producer.

for i in {1..10}
do
    # Create a unique key and value for each message
    if (( i % 2 == 0 )); then
        key="user_A"
    else
        key="user_B"
    fi

    # key="photo_id_$i"
    value="{\"filename\":\"photo_$i.jpg\", \"timestamp\":\"$(date +%s)\"}"

    # Echo the message in the key:value format
    echo "$key:$value"

    sleep 1.0

done | docker exec -i kafka-3 /opt/kafka/bin/kafka-console-producer.sh \
    --topic photo-processing-queue \
    --bootstrap-server kafka-3:29092 \
    --property "parse.key=true" \
    --property "key.separator=:"

echo "10 messages have been sent to the photo-processing-queue topic."
