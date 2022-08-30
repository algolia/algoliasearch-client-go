import { predictClient } from '@algolia/predict';
import { ApiError } from '@algolia/client-common';
import dotenv from 'dotenv';

dotenv.config({ path: '../../.env' });

const appId = process.env.ALGOLIA_PREDICT_APP_ID || '**** APP_ID *****';
const apiKey =
  process.env.ALGOLIA_PREDICT_API_KEY || '**** PREDICT_API_KEY *****';
const userId = process.env.ALGOLIA_PREDICT_USER_ID || 'user1';

// Init client with appId and apiKey
const client = predictClient(appId, apiKey, 'ew');

async function testPredict() {
  try {
    const userProfile = await client.fetchUserProfile({
      userID: userId,
      params: {
        modelsToRetrieve: ['funnel_stage', 'order_value', 'affinities'],
        typesToRetrieve: ['properties', 'segments'],
      },
    });

    console.log(`[OK]`, userProfile);
  } catch (e) {
    if (e instanceof ApiError) {
      return console.log(`[${e.status}] ${e.message}`, e.stackTrace);
    }

    console.log('[ERROR]', e);
  }
}

testPredict();
