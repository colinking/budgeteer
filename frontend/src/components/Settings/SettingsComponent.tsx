import { Card, Heading, majorScale, Pane } from 'evergreen-ui';
import React from 'react';
import PlaidLink from 'react-plaid-link';

export interface SettingsProps {
  plaidEnv: string,
  plaidPublicKey: string,
  handleOnLinkExit: (err: Error | undefined) => void,
  handleOnLinkSuccess: (token: string, metadata: any) => void
}

export default class SettingsComponent extends React.Component<SettingsProps> {
  public render() {
    const props = this.props

    return (
      <Pane display="flex" justifyContent="center">
        <Card
          width={majorScale(80)}
          height="auto"
          elevation={0}
          backgroundColor="white"
          marginY={majorScale(2)}
          padding={majorScale(2)}
        >
          <Pane flexDirection="column">
            <Heading marginBottom={majorScale(2)}>Settings</Heading>
            <PlaidLink
              clientName="MOSS"
              env={props.plaidEnv}
              product={['transactions']}
              publicKey={props.plaidPublicKey}
              onExit={props.handleOnLinkExit}
              onSuccess={props.handleOnLinkSuccess}
            >
              Open Link to connect your bank!
            </PlaidLink>
          </Pane>
        </Card>
      </Pane>
    );
  }
}