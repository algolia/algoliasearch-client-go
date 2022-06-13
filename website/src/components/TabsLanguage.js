import Tabs from '@theme/Tabs';
import React from 'react';

export const languagesTabValues = [
  { label: 'JavaScript', value: 'javascript' },
  { label: 'PHP', value: 'php' },
  { label: 'Java', value: 'java' },
];

export function TabsLanguage(props) {
  return (
    <Tabs groupId="language" defaultValue="java" values={props.values}>
      {props.children}
    </Tabs>
  );
}

TabsLanguage.defaultProps = {
  values: languagesTabValues,
};
