import { UserServiceClient } from '../gen/userpb/user_service_pb_service'
import { getHost } from './requests'

export default {
  users: new UserServiceClient(getHost())
}
