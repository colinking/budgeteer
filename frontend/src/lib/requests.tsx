import { grpc } from 'grpc-web-client'
import { getAccessToken, isAuthenticated } from './auth'

const port = process.env.PORT || 9091

export function getHost(): string {
  return `https://localhost:${port}`
}

export function getMetadata(): grpc.Metadata {
  const headers: {
    [key: string]: string
  } = {}

  if (isAuthenticated()) {
    headers.Authorization = `Bearer ${getAccessToken()}`
  }
  return new grpc.Metadata(headers)
}
