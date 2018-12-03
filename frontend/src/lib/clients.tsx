import { UserServiceClient } from '../gen/userpb/user_service_pb_service'
import { PlaidServiceClient } from '../gen/plaidpb/plaid_service_pb_service'
import { getHost } from './requests'

export default {
  users: new UserServiceClient(getHost()),
  plaid: new PlaidServiceClient(getHost())
}
