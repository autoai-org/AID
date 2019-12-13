
/**
 * Following is the definition for mapper
 */

import { DataMapper } from '@aws/dynamodb-data-mapper';
import AWS = require('aws-sdk');
import DynamoDB = require('aws-sdk/clients/dynamodb');
import { AWSConfig } from './config';

AWS.config.update(AWSConfig);

const mapper = new DataMapper({
    client: new DynamoDB({region: 'us-east-1'})
});

export default mapper;