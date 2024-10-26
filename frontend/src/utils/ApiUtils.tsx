interface ApiResponse<T = any> {
  data?: T;
  error?: string;
}
import { useAuth } from '../contexts/AuthContext';

export const useApi = () => {
  const { logout } = useAuth();

  const callApi = async (
    endpoint: string,
    method: 'POST' | 'GET' | 'PUT' | 'DELETE',
    payload: object = {}
  ): Promise<ApiResponse> => {
    try {
      const token = localStorage.getItem('BoosterHubToken');

      const headers: HeadersInit = {
        'Content-Type': 'application/json',
      };

      if (token) {
        headers['Authorization'] = `Bearer ${token}`;
      }

      const url = `http://localhost:3000/api${endpoint}`;

      const response = await fetch(url, {
        method,
        headers,
        body: method !== 'GET' ? JSON.stringify(payload) : null,
      });

      if (response.status === 401) {
        logout();
        return { error: 'Unauthorized access. You have been logged out.' };
      }

      if (response.ok) {
        const result = await response.json();
        return { data: result };
      } else {
        const errorData = await response.json();
        return { error: errorData.error || 'Request failed, please try again.' };
      }
    } catch (error) {
      return { error: 'An error occurred. Please try again.' };
    }
  };

  return { callApi };
};
