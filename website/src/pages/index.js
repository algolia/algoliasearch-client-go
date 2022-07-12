import {
  Hero as AlgoliaHero,
  Button,
  Heading1,
  InlineLink,
  Heading3,
} from '@algolia/ui-library';
import Layout from '@theme/Layout';
import React from 'react';

function Hero() {
  function Title() {
    return (
      <Heading1 className="hero-title">
        Generated API clients, by{' '}
        <span className="uil-color-nebula-500">Algolia</span>.
      </Heading1>
    );
  }

  return (
    <AlgoliaHero
      id="hero"
      cta={[
        <Button
          key="get-started"
          href="/docs/clients/introduction"
          color="grey"
        >
          Get started
        </Button>,
        <Button
          key="contribute"
          href="/docs/contributing/introduction"
          background="blue"
          color="white"
        >
          Contribute
        </Button>,
        <iframe
          src="https://ghbtns.com/github-btn.html?user=algolia&amp;repo=api-clients-automation&amp;type=star&amp;count=true&amp;size=large"
          width={160}
          height={30}
          title="GitHub Stars"
        />,
      ]}
    >
      <Title />
    </AlgoliaHero>
  );
}

function MigrationNotice() {
  return (
    <div id="notice" className="uil-ta-center">
      <Heading3 className="uil-m-0 uil-fw-bold">
        Coming from{' '}
        <InlineLink target="_blank" href="https://www.algolia.com/doc/">
          the current API clients
        </InlineLink>{' '}
        ? Check out our{' '}
        <InlineLink href="/docs/clients/migration-guides/">
          migration guide
        </InlineLink>
        .
      </Heading3>
    </div>
  );
}

export default function Home() {
  return (
    <Layout description="API Clients Automation by Algolia">
      <Hero />
      <MigrationNotice />
    </Layout>
  );
}
