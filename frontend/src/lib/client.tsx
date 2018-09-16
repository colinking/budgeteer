
const port = process.env.PORT || 9091

export function getHost(): string {
  return `https://localhost:${port}`
}