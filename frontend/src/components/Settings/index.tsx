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
    new PlaidClient(getHost()).exchangeToken(
      req,
      (error: ServiceError | null, responseMessage: ExchangeTokenResponse | null) => {
        if (error) {
          throw error
        }
        const resp = responseMessage as ExchangeTokenResponse
        console.log(resp.getAccessToken());
        console.log(resp.getItemId());
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
