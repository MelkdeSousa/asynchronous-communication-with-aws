#!/bin/bash

# ENVIRONMENTS=("develop" "test" "prod" "local")

# for env in "${ENVIRONMENTS[@]}"; do
awslocal sns create-topic \
    --region us-east-1 \
    --name "user-created.fifo" \
    --attributes FifoTopic=true

awslocal sqs create-queue \
    --region us-east-1 \
    --queue-name points-user-created.fifo \
    --attributes FifoQueue=true

awslocal sns subscribe \
    --region us-east-1 \
    --topic-arn "arn:aws:sns:us-east-1:000000000000:user-created.fifo" \
    --protocol sqs \
    --notification-endpoint "arn:aws:sqs:us-east-1:000000000000:points-user-created.fifo"
# done
