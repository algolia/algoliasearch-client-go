import { ingestionClient } from '@algolia/ingestion';
import { ApiError } from '@algolia/client-common';
import dotenv from 'dotenv';

dotenv.config({ path: '../../.env' });

const appId = process.env.ALGOLIA_APPLICATION_ID || '**** APP_ID *****';
const apiKey =
  process.env.ALGOLIA_ADMIN_KEY ||
  '**** ADMIN_KEY *****';

// Init client with appId and apiKey
const client = ingestionClient(appId, apiKey, 'us', {authMode: 'WithinHeaders', hosts: [{url: 'staging-data.us.algolia.com', accept: 'readWrite', protocol: 'https'}]});

async function testIngestion() {
  try {
    const res = await client.getAuthentications();

    console.log(`[OK]`, res);
  } catch (e) {
    if (e instanceof ApiError) {
      return console.log(`[${e.status}] ${e.message}`, e.stackTrace, e);
    }

    console.log('[ERROR]', e);
  }
}

testIngestion();
