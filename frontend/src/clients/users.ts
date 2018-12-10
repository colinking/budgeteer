import { UserServiceClient, ServiceError } from './gen/userpb/user_service_pb_service'
import { getHost, getMetadata } from './lib/requests'
import { LoginRequest, LoginResponse, GetRequest, GetResponse, AddItemRequest, AddItemResponse } from './gen/userpb/user_service_pb'

const client = new UserServiceClient(getHost())

export async function login(loadRequest?: ((req: LoginRequest) => void)): Promise<LoginResponse> {
  return new Promise<LoginResponse>((resolve, reject) => {
    const req = new LoginRequest()
    if (loadRequest) {
      loadRequest(req)
    }

    client.login(req, getMetadata(), (err: ServiceError | null, res: LoginResponse | null) => {
      if (err || !res) {
        return reject(err)
      }
      resolve(res)
    })
   }
 )
}

export async function get(loadRequest?: ((req: GetRequest) => void)): Promise<GetResponse> {
  return new Promise<GetResponse>((resolve, reject) => {
    const req = new GetRequest()
    if (loadRequest) {
      loadRequest(req)
    }

    client.get(req, getMetadata(), (err: ServiceError | null, res: GetResponse | null) => {
      if (err || !res) {
        return reject(err)
      }
      resolve(res)
    })
   }
 )
}

export async function addItem(loadRequest?: ((req: AddItemRequest) => void)): Promise<AddItemResponse> {
  return new Promise<AddItemResponse>((resolve, reject) => {
    const req = new AddItemRequest()
    if (loadRequest) {
      loadRequest(req)
    }

    client.addItem(req, getMetadata(), (err: ServiceError | null, res: AddItemResponse | null) => {
      if (err || !res) {
        return reject(err)
      }
      resolve(res)
    })
   }
 )
}
