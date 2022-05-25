/* eslint-disable @typescript-eslint/no-var-requires */
/* eslint-disable import/no-commonjs */

const fs = require('fs');
const path = require('path');

const darkCodeTheme = require('prism-react-renderer/themes/dracula');
const lightCodeTheme = require('prism-react-renderer/themes/github');

function getSpecFiles() {
  const ROOT_DIR = path.resolve(process.cwd(), '..');
  const bundledSpecsPath = path.resolve(ROOT_DIR, 'specs/bundled');
  const specFiles = [];

  fs.readdirSync(bundledSpecsPath).forEach((file) => {
    if (file.endsWith('.doc.yml')) {
      const fileName = file.replace('.doc.yml', '');

      specFiles.push({
        fileName,
        path: `${bundledSpecsPath}/${file}`,
      });
    }
  });

  if (specFiles.length === 0) {
    throw new Error('Unable to find spec files');
  }

  return specFiles;
}

function getSpecsForPlugin() {
  return getSpecFiles().map((specFile) => {
    return {
      id: specFile.fileName,
      spec: specFile.path,
      route: `/specs/${specFile.fileName}`,
    };
  });
}

function getSpecsForNavBar() {
  return getSpecFiles().map((specFile) => {
    /** @type {import('@docusaurus/theme-common').NavbarItem} */
    return {
      label: specFile.fileName,
      href: `/specs/${specFile.fileName}`,
      className: 'header-restapi',
    };
  });
}

/** @type {import('@docusaurus/types').Config} */
(
  module.exports = {
    title: 'Algolia API',
    tagline: 'Documentation for the Algolia API and Clients.',
    url: 'https://api-clients-automation.netlify.app/',
    baseUrl: '/',
    favicon: 'img/logo-small.png',
    organizationName: 'Algolia',
    projectName: 'Algolia API and Clients',
    onBrokenLinks: 'throw',
    onBrokenMarkdownLinks: 'throw',

    presets: [
      [
        'redocusaurus',
        {
          specs: getSpecsForPlugin(),
          theme: {
            options: { disableSearch: true },
            primaryColor: '#5468ff',
            theme: {
              typography: { fontSize: '14px', lineHeight: '1.2em' },
              spacing: {
                unit: 5,
                sectionHorizontal: 30,
              },
            },
          },
        },
      ],
      [
        '@docusaurus/preset-classic',
        /** @type {import('@docusaurus/preset-classic').Options} */
        ({
          docs: {
            path: 'docs',
            sidebarPath: 'sidebars.js',
            editUrl:
              'https://github.com/algolia/api-clients-automation/edit/main/website/',
            showLastUpdateAuthor: true,
            showLastUpdateTime: true,
          },
          blog: false,
          theme: {
            customCss: require.resolve('./src/css/custom.css'),
          },
        }),
      ],
    ],

    themeConfig:
      /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
      ({
        navbar: {
          title: 'Algolia API',
          logo: {
            alt: 'Algolia',
            src: 'img/logo-small.png',
            srcDark: 'img/logo-small.png',
          },
          items: [
            // left
            {
              label: 'Clients',
              to: 'docs/clients/introduction',
              position: 'left',
            },
            {
              label: 'Contributing',
              to: 'docs/contributing/introduction',
              position: 'right',
            },
            {
              label: 'HTTP API',
              position: 'left',
              type: 'dropdown',
              items: getSpecsForNavBar(),
            },
            // right
            {
              href: 'https://github.com/algolia/api-clients-automation',
              position: 'right',
              className: 'header-github-link',
            },
          ],
        },
        algolia: {
          appId: 'QPBQ67WNIG',
          apiKey: 'b590ae1153bf574215ca1605c19eb1fe',
          indexName: 'api-clients-automation',
        },
        colorMode: {
          defaultMode: 'light',
          disableSwitch: false,
          respectPrefersColorScheme: true,
        },
        // Breaks scrollbar
        // announcementBar: {
        //   content:
        //     '⭐️ If you like our API clients, give them a star on <a target="_blank" rel="noopener noreferrer" href="https://github.com/algolia/api-clients-automation">GitHub</a>! ⭐️',
        // },
        footer: {
          style: 'dark',
          links: [
            {
              label: 'GitHub',
              to: 'https://github.com/algolia/api-clients-automation',
            },
            {
              label: 'Twitter',
              to: 'https://twitter.com/algolia',
            },
            {
              label: 'Algolia Website',
              to: 'https://algolia.com/',
            },
            {
              label: 'Algolia Blog',
              to: 'https://algolia.com/blog',
            },
          ],
          copyright: `Copyright © ${new Date().getFullYear()} Algolia | Built with Docusaurus.`,
        },
        prism: {
          theme: lightCodeTheme,
          darkTheme: darkCodeTheme,
          additionalLanguages: ['java', 'php'],
        },
      }),
  }
);
