export class RequestTimeoutError extends Error {}

interface BaseResponseError {
  code: string;
  message: string;
}

export interface BaseResponse<T> {
  status: boolean;
  error?: BaseResponseError;
  data?: T;
}

export interface BaseRequest<T> {
  data: BaseRequestData<T>;
}

interface BaseRequestData<T> {
  attributes: T;
}

/**
 * Thực hiện request http bằng fetch
 */
const makeRequest = async <T>({
  method,
  url,
  data,
  timeout = 30000, // 30s
  baseUrl = window._apis.baseUrl,
}: {
  method: 'GET' | 'POST' | 'PUT' | 'DELETE';
  url: string;
  data: any;
  timeout?: number;
  baseUrl?: string;
}) => {
  const controller = new AbortController();
  const id = setTimeout(() => controller.abort(), timeout);
  try {
    const response = await fetch(`${baseUrl}${url}`, {
      method: method, // *GET, POST, PUT, DELETE, etc.
      mode: 'cors', // no-cors, *cors, same-origin
      cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
      credentials: 'include', // include, *same-origin, omit
      headers: {
        'Content-Type': 'application/json',
        // 'Content-Type': 'application/x-www-form-urlencoded',
      },
      redirect: 'follow', // manual, *follow, error
      referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
      body: JSON.stringify(data), // body data type must match "Content-Type" header
    });
    clearTimeout(id);
    return (await response.json()) as T; // parses JSON response into native JavaScript objects
  } catch (error: any) {
    if (error.native === 'AbortError') {
      throw new RequestTimeoutError('Request timed out');
    }
    throw error;
  }
};

const delay = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));

export default {
  makeRequest,
  delay,
} as const;
