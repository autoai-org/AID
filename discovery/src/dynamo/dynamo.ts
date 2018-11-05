
/**
 * Following is the definition for mapper
 */

import { DataMapper } from '@aws/dynamodb-data-mapper';
import DynamoDB = require('aws-sdk/clients/dynamodb');

const mapper = new DataMapper({
    client: new DynamoDB({region: 'us-east-1'})
})
