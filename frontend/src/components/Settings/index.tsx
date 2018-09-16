import React from 'react';

import { getHost } from '../../lib/client';
import { ExchangeTokenRequest, ExchangeTokenResponse } from '../../proto/plaid/plaid_service_pb';
import { PlaidClient, ServiceError } from '../../proto/plaid/plaid_service_pb_service';
import SettingsComponent from './SettingsComponent';

const plaidEnv = process.env.REACT_APP_PLAID_ENV as string;
const plaidPublicKey = process.env.REACT_APP_PLAID_PUBLIC_KEY as string;

interface SettingsProps {}

export default class Settings extends React.Component<SettingsProps> {
  public handleOnLinkExit(err: Error | undefined) {
    console.log('User exited Link.');
    if (err) {
      this.handleLinkError(err);
    }
  }

  public handleLinkError(err: Error | ServiceError) {
    console.error(err)
  }

  public handleOnLinkSuccess(token: string, metadata: any) {
    console.log('Successfully authenticated user:');
    console.log(token);
    console.log(metadata);
    const req = new ExchangeTokenRequest();
    req.setToken(token);
    console.log(req);
    const client = new PlaidClient(getHost());
    client.exchangeToken(
      req,
      (error: ServiceError, responseMessage: ExchangeTokenResponse | null) => {
        console.log('returned from request');
        console.log(error);
        console.log((responseMessage as ExchangeTokenResponse).getAccessToken());
        console.log((responseMessage as ExchangeTokenResponse).getItemId());
      }
    );
  }

  public render() {
    return (
      <SettingsComponent
        handleOnLinkExit={this.handleOnLinkExit}
        handleOnLinkSuccess={this.handleOnLinkSuccess}
        plaidEnv={plaidEnv}
        plaidPublicKey={plaidPublicKey}
      />
    );
  }
}
