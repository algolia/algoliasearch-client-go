// @ts-check

/** @type {import('@docusaurus/plugin-content-docs').SidebarsConfig} */
const sidebars = {
  // Everything related to the API Clients Automation
  automation: [
    'contributing/introduction',
    {
      type: 'category',
      label: 'Getting Started',
      collapsed: false,
      items: [
        'contributing/setup-repository',
        {
          type: 'category',
          label: 'CLI',
          collapsed: false,
          items: [
            'contributing/CLI/specs-commands',
            'contributing/CLI/clients-commands',
            'contributing/CLI/cts-commands',
          ],
        },
      ],
    },
    {
      type: 'category',
      label: 'Contributing',
      collapsed: false,
      items: [
        'contributing/add-new-api-client',
        'contributing/add-new-language',
        {
          type: 'category',
          label: 'Testing',
          collapsed: false,
          items: [
            'contributing/testing/common-test-suite',
            'contributing/testing/playground',
          ],
        },
        'contributing/commit-and-pull-request',
        'contributing/release-process',
      ],
    },
  ],
  // Everything related to the generated clients usage
  clients: [
    'clients/introduction',
    {
      type: 'category',
      label: 'Getting Started',
      collapsed: false,
      items: ['clients/installation', 'clients/migration-guide'],
    },
    {
      type: 'category',
      label: 'Guides',
      collapsed: false,
      items: [
        'clients/guides/send-data-to-algolia',
        'clients/guides/filtering-your-search',
        'clients/guides/retrieving-facets',
      ],
    },
  ],
};

// eslint-disable-next-line import/no-commonjs
module.exports = sidebars;
