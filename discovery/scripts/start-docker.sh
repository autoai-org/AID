#!bin/bash
docker run  \
            --env AWS_ACCESS_KEY=$AWS_ACCESS_KEY \
            --env AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
            --env LOGSENE_TOKEN=$LOGSENE_TOKEN \
            --env COGNITO_POOL_ID=$COGNITO_POOL_ID \
            --env COGNITO_CLIENT_ID=$COGNITO_CLIENT_ID \
            --env COGNITO_REGION=$COGNITO_REGION \
            --env PARSE_ID=$PARSE_ID \
            --env PARSE_TOKEN=$PARSE_TOKEN \
            --env PARSE_URL=$PARSE_URL \
            --env PARSE_MASTER_KEY=$PARSE_MASTER_KEY \
            -d -p 3000:3000 docker.io/autoai/cvpm-discovery