import { searchClient } from '@algolia/client-search';
import { ApiError, EchoResponse } from '@algolia/client-common';
import dotenv from 'dotenv';
import { echoRequester } from '@algolia/requester-node-http';

dotenv.config({ path: '../../.env' });

const appId = process.env.ALGOLIA_APPLICATION_ID || '**** APP_ID *****';
const apiKey = process.env.ALGOLIA_SEARCH_KEY || '**** SEARCH_API_KEY *****';

const searchIndex = process.env.SEARCH_INDEX || 'test_index';
const searchQuery = process.env.SEARCH_QUERY || 'test_query';

// Init client with appId and apiKey
const client = searchClient(appId, apiKey, { requester: echoRequester() });

client.addAlgoliaAgent('Node playground', '0.0.1');

async function testSearch() {
  try {
    const res = (await client.post({
      path: '/test/minimal',
    })) as unknown as EchoResponse;

    console.log(`[OK]`, res);
  } catch (e: any) {
    // Instance of
    if (e instanceof ApiError) {
      return console.log(`[${e.status}] ${e.message}`, e.stackTrace);
    }

    // Other way
    if (e.name === 'RetryError') {
      return console.log(`[${e.name}] ${e.message}`, e.stackTrace);
    }

    console.log('[ERROR]', e);
  }
}

testSearch();
