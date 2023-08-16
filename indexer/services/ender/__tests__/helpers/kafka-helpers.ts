import { createKafkaMessage } from '@dydxprotocol-indexer/kafka';
import {
  IndexerTendermintBlock,
  IndexerTendermintEvent,
  Timestamp,
  MarketEventV1,
  StatefulOrderEventV1,
} from '@dydxprotocol-indexer/v4-protos';
import { KafkaMessage } from 'kafkajs';

import { DydxIndexerSubtypes } from '../../src/lib/types';
import { defaultHeight, defaultTime, defaultTxHash } from './constants';
import { binaryToBase64String, createIndexerTendermintBlock, createIndexerTendermintEvent } from './indexer-proto-helpers';

export function createKafkaMessageFromMarketEvent({
  marketEvent,
  transactionIndex,
  height,
  time,
  txHash,
}: {
  marketEvent: MarketEventV1 | undefined,
  transactionIndex: number,
  height: number,
  time: Timestamp,
  txHash: string,
}): KafkaMessage {
  const events: IndexerTendermintEvent[] = [];
  if (marketEvent !== undefined) {
    events.push(
      createIndexerTendermintEvent(
        DydxIndexerSubtypes.MARKET,
        binaryToBase64String(
          MarketEventV1.encode(marketEvent).finish(),
        ),
        transactionIndex,
        0,
      ),
    );
  }

  const block: IndexerTendermintBlock = createIndexerTendermintBlock(
    height,
    time,
    events,
    [txHash],
  );

  const binaryBlock: Uint8Array = IndexerTendermintBlock.encode(block).finish();
  return createKafkaMessage(Buffer.from(binaryBlock));
}

export function createKafkaMessageFromStatefulOrderEvent(
  event: StatefulOrderEventV1,
  transactionIndex: number = 0,
  height: number = defaultHeight,
  time: Timestamp = defaultTime,
  txHash: string = defaultTxHash,
): KafkaMessage {
  const events: IndexerTendermintEvent[] = [];
  events.push(
    createIndexerTendermintEvent(
      DydxIndexerSubtypes.STATEFUL_ORDER,
      binaryToBase64String(
        StatefulOrderEventV1.encode(event).finish(),
      ),
      transactionIndex,
      0,
    ),
  );

  const block: IndexerTendermintBlock = createIndexerTendermintBlock(
    height,
    time,
    events,
    [txHash],
  );

  const binaryBlock: Uint8Array = IndexerTendermintBlock.encode(block).finish();
  return createKafkaMessage(Buffer.from(binaryBlock));
}