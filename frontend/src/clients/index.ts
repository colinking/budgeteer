import { ServiceError } from './gen/userpb/user_service_pb_service'
import { User, Item, Account } from './gen/userpb/user_service_pb'

import * as users from './users'

export {
  // ServiceError is identical for all services, so we only export one of them.
  ServiceError,

  users
}

export type User = User.AsObject
export type Item = Item.AsObject
export type Account = Account.AsObject
