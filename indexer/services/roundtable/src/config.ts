/**
 * Environment variables required by Roundtable.
 */

import {
  parseSchema,
  baseConfigSchema,
  parseBoolean,
  parseInteger,
  parseNumber,
  parseString,
} from '@dydxprotocol-indexer/base';
import {
  kafkaConfigSchema,
} from '@dydxprotocol-indexer/kafka';
import {
  postgresConfigSchema,
} from '@dydxprotocol-indexer/postgres';
import {
  redisConfigSchema,
} from '@dydxprotocol-indexer/redis';

export const configSchema = {
  ...baseConfigSchema,
  ...postgresConfigSchema,
  ...kafkaConfigSchema,
  ...redisConfigSchema,

  // Loop Enablement
  LOOPS_ENABLED_MARKET_UPDATER: parseBoolean({ default: true }),
  LOOPS_ENABLED_DELETE_ZERO_PRICE_LEVELS: parseBoolean({ default: true }),
  LOOPS_ENABLED_PNL_TICKS: parseBoolean({ default: true }),
  LOOPS_ENABLED_REMOVE_EXPIRED_ORDERS: parseBoolean({ default: true }),
  LOOPS_ORDERBOOK_INSTRUMENTATION: parseBoolean({ default: true }),
  LOOPS_CANCEL_STALE_ORDERS: parseBoolean({ default: true }),
  LOOPS_ENABLED_UPDATE_RESEARCH_ENVIRONMENT: parseBoolean({ default: true }),

  // Loop Timing
  LOOPS_INTERVAL_MS_MARKET_UPDATER: parseInteger({
    // 10 seconds
    default: 10000,
  }),
  LOOPS_INTERVAL_MS_DELETE_ZERO_PRICE_LEVELS: parseInteger({
    // 2 minutes
    default: 120_000,
  }),
  LOOPS_INTERVAL_MS_PNL_TICKS: parseInteger({
    // 30 seconds
    default: 30_000,
  }),
  LOOPS_INTERVAL_MS_REMOVE_EXPIRED_ORDERS: parseInteger({
    // 2 minutes
    default: 120_000,
  }),
  LOOPS_INTERVAL_MS_ORDERBOOK_INSTRUMENTATION: parseInteger({
    // 5 seconds
    default: 5_000,
  }),
  LOOPS_INTERVAL_MS_CANCEL_STALE_ORDERS: parseInteger({
    // 30 seconds
    default: 30_000,
  }),
  LOOPS_INTERVAL_MS_UPDATE_RESEARCH_ENVIRONMENT: parseInteger({
    // 1 hour
    default: 3_600_000,
  }),

  // Start delay
  START_DELAY_ENABLED: parseBoolean({ default: true }),
  MAX_START_DELAY_MS: parseInteger({ default: 180000 }), // 3 minutes
  MAX_START_DELAY_FRACTION_OF_INTERVAL: parseNumber({ default: 0.1 }),
  JITTER_FRACTION_OF_DELAY: parseNumber({ default: 0.01 }),

  // Lock multipliers
  MARKET_UPDATER_LOCK_MULTIPLIER: parseInteger({ default: 10 }),
  DELETE_ZERO_PRICE_LEVELS_LOCK_MULTIPLIER: parseInteger({ default: 1 }),

  // Maximum number of running tasks - set this equal to PG_POOL_MIN in .env, default is 2
  MAX_CONCURRENT_RUNNING_TASKS: parseInteger({ default: 2 }),
  EXCEEDED_MAX_CONCURRENT_RUNNING_TASKS_DELAY_MS: parseInteger({ default: 1000 }),

  // PNL ticks
  PNL_TICK_UPDATE_INTERVAL_MS: parseInteger({ default: 3_600_000 }), // 1 hour
  PNL_TICK_MAX_ROWS_PER_UPSERT: parseInteger({ default: 1000 }),

  // Remove expired orders
  BLOCKS_TO_DELAY_EXPIRY_BEFORE_SENDING_REMOVES: parseInteger({ default: 20 }),

  // Cancel stale orders
  CANCEL_STALE_ORDERS_QUERY_BATCH_SIZE: parseInteger({ default: 10000 }),

  // Update research environment
  AWS_ACCOUNT_ID: parseString(),
  AWS_REGION: parseString(),
  S3_BUCKET_ARN: parseString(),
  ECS_TASK_ROLE_ARN: parseString(),
  KMS_KEY_ARN: parseString(),
  RDS_INSTANCE_NAME: parseString(),
  ATHENA_CATALOG_NAME: parseString({ default: 'AwsDataCatalog' }),
  ATHENA_DATABASE_NAME: parseString({ default: 'default' }),
  ATHENA_WORKING_GROUP: parseString({ default: 'primary' }),
  SKIP_TO_ATHENA_TABLE_WRITING: parseBoolean({ default: false }),
};

export default parseSchema(configSchema);