import auth0, {
  Auth0DecodedHash,
  Auth0UserProfile,
  ParseHashOptions
} from 'auth0-js'
import * as tp from 'typed-promisify'

export declare interface User {
  auth_id: string
  name: string
  email: string
  picture: string
}

export const ACCESS_TOKEN_KEY = 'access_token'
export const ID_TOKEN_KEY = 'id_token'
export const EXPIRATION_KEY = 'expires_at'

const auth = new auth0.WebAuth({
  clientID: process.env.REACT_APP_AUTH0_CLIENT_ID as string,
  domain: process.env.REACT_APP_AUTH0_DOMAIN as string,
  redirectUri: process.env.REACT_APP_AUTH0_CALLBACK_URL as string,
  audience: process.env.REACT_APP_BACKEND_AUDIENCE_ID as string,
  responseType: 'token id_token',
  scope: 'openid profile email'
})

const parseHash = tp.promisify<ParseHashOptions, Auth0DecodedHash | null>(
  auth.parseHash,
  auth
)
const userInfo = tp.promisify<string, Auth0UserProfile>(
  auth.client.userInfo,
  auth.client
)

/**
 * Parses an Auth0 callback response and stores the newly logged-in user, if any.
 */
export async function handleAuthenticationCallback(hash: string) {
  const authResult = await parseHash({ hash })

  if (!authResult) {
    return Promise.reject("Invalid authentication metadata. Unable to parse")
  }

  // Set the time that the access token will expire at.
  const expiresAt = JSON.stringify(
    authResult.expiresIn! * 1000 + new Date().getTime()
  )
  localStorage.setItem(ACCESS_TOKEN_KEY, authResult.accessToken!)
  localStorage.setItem(ID_TOKEN_KEY, authResult.idToken!)
  localStorage.setItem(EXPIRATION_KEY, expiresAt)
  window.dispatchEvent(new Event('storage'))
}

/**
 * Check whether there is a currently logged-in and non-expired user.
 */
export function isAuthenticated(): boolean {
  // Check whether the access token has expired.
  const expiration = localStorage.getItem(EXPIRATION_KEY)
  return !!expiration && new Date().getTime() < JSON.parse(expiration)
}

/**
 * Redirect user to Auth0 login page.
 * User will be redirected to our callback handler after completing, exiting, or erroring on that flow.
 */
export function login(): void {
  auth.authorize()
}

/**
 * Logs out any currently logged-in user, if one exists.
 */
export function logout(): void {
  // Clear access token and ID token from local storage.
  localStorage.removeItem(ACCESS_TOKEN_KEY)
  localStorage.removeItem(ID_TOKEN_KEY)
  localStorage.removeItem(EXPIRATION_KEY)
  window.dispatchEvent(new Event('storage'))
}

export function getAccessToken(): string {
  const accessToken = localStorage.getItem(ACCESS_TOKEN_KEY)
  if (!accessToken) {
    throw new Error('No access token found')
  }
  return accessToken
}

/**
 * Pulls authentication data from local storage to query Auth for user information.
 */
export async function getLoggedInUser(): Promise<User> {
  if (!isAuthenticated()) {
    throw new Error('No logged-in user')
  }

  const token = await getAccessToken()
  const authUser = await userInfo(token)

  return {
    auth_id: authUser.sub!,
    email: authUser.email!,
    name: authUser.name,
    picture: authUser.picture
  }
}
