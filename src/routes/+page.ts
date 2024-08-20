import type { PageLoad } from "./$types";
export const prerender = true;

export const load: PageLoad = async ({ fetch, setHeaders }) => {
  // Call the API to check if it's up
  const url = `https://gorelease-1-v5611824.deta.app/backend/api/v1/health`;
  const response = await fetch(url);
  if (response.status !== 200) {
    throw new Error(`API health check failed`);
  }
  const data = await response.json();
  return {
    props: {
      status: data.status,
    },
  };
};
