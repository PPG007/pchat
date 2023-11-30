export const getMessage = (e: any, fallback: string) => {
  if (typeof e === 'string') {
    return e
  }
  return fallback;
}