import { searchClient } from '@experimental-api-clients-automation/client-search';

import './app.css';

const client = searchClient('QPBQ67WNIG', 'b590ae1153bf574215ca1605c19eb1fe');

client.addAlgoliaAgent('Browser playground', '0.0.1');

const searchButton = document.querySelector('#search');

searchButton?.addEventListener('click', async () => {
  const { results } = await client.search({
    requests: [
      {
        indexName: 'docsearch',
        query: 'docsearch',
        hitsPerPage: 50,
      },
    ],
  });

  const parent = document.querySelector('#results');

  results[0].hits?.forEach(({ objectID }) => {
    const children = document.createElement('p');
    children.innerHTML = objectID;

    parent?.appendChild(children);
  });
});
