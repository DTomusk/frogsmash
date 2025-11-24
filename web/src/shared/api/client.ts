export const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

export async function apiFetch<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
  const isFormData = options.body instanceof FormData;
  const token = localStorage.getItem('token');

  // Don't include Content-Type header for FormData requests
  const headers = { 
      ...(options.headers || {}),
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...(!isFormData ? { 'Content-Type': 'application/json' } : {}),
    };

  // Serialize body to JSON if it's a plain object
  const body =
    !isFormData && options.body && typeof options.body === 'object'
      ? JSON.stringify(options.body)
      : options.body;

  // Make the fetch request
  let response = await fetch(`${API_BASE_URL}${endpoint}`, {
    ...options,
    headers,
    body,
    credentials: 'include',
  });

  // If unauthorized, try to refresh the token and retry the request
  if (response.status === 401) {
    const refreshResponse = await fetch(`${API_BASE_URL}/auth/refresh-token`, {
      method: 'POST',
      credentials: 'include',
    });

    if (!refreshResponse.ok) {
      localStorage.removeItem('token');
      throw new Error('Session expired. Please log in again.');
    }

    const refreshData = await refreshResponse.json();
    localStorage.setItem('token', refreshData.jwt);

    const retryHeaders = { 
      ...(options.headers || {}),
      Authorization: `Bearer ${refreshData.jwt}`,
      ...(!isFormData ? { 'Content-Type': 'application/json' } : {}),
    };

    response = await fetch(`${API_BASE_URL}${endpoint}`, {
      ...options,
      headers: retryHeaders,
      body,
      credentials: 'include',
    });
  }

  if (!response.ok) {
    const errorText = await response.text();
    let errorData: any = {};

    try {
      errorData = JSON.parse(errorText);
    } catch {
      // TODO: consider what to do if response is not JSON
    }

    const message =
      errorData.error ||
      errorData.message ||
      `API request failed with status ${response.status}`;

    // Create rich error object
    const err: any = new Error(message);
    err.status = response.status;
    err.data = errorData;

    throw err;
  }

  // If the response has text, parse it as JSON
  const text = await response.text();
  return text ? JSON.parse(text) : ({} as T);
}
