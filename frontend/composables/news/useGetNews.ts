export function useGetNews() {
  const { data, status, error, refresh, clear } = useLazyAsyncData(
    'useGetNews',
    () => $fetch('/api/news')
  );

  return {
    data,
    status,
    error,
    refresh,
    clear
  };
}
