import { Heading4, InlineLink } from '@algolia/ui-library';
import React from 'react';

import '../css/custom.css';

export default function Root({ children }) {
  function WipNotice() {
    return (
      <div id="notice" className="uil-ta-center wip-notice">
        <Heading4 className="uil-fw-bold uil-m-8">
          The generated API clients are a work in progress, you can also find
          our stable clients{' '}
          <InlineLink target="_blank" href="https://www.algolia.com/doc/">
            on the Algolia documentation
          </InlineLink>
          .
        </Heading4>
      </div>
    );
  }

  return (
    <>
      <WipNotice />
      {children}
    </>
  );
}
