import React, { useState } from 'react';
import { motion } from 'framer-motion';
import { useAuth } from '../contexts/AuthContext';
import { decryptToken } from '../utils/TokenUtils';
import { useNavigate } from 'react-router-dom';
import { useApi } from '../utils/ApiUtils';

const Login: React.FC = () => {
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [userType, setUserType] = useState<string>('creator');
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const { setUserInfo, isAuthenticated } = useAuth();
  const { callApi } = useApi();
  const navigate = useNavigate();

  if (isAuthenticated()) {
    navigate('/');
  }

  const handleToggleChange = (val: string) => {
    setUserType(val);
    setErrorMessage(null);
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    setErrorMessage(null);

    const data = {
      email,
      password,
    };

    const endpoint = userType === 'creator' ? '/creator/log-in' : '/business/log-in';

    const response = await callApi(endpoint, 'POST', data);

    if (response.error) {
      setErrorMessage(response.error || 'Login failed, please try again.');
    } else {
      const token = response.data.token;
      localStorage.setItem('BoosterHubToken', token);

      const decodedInfo = decryptToken(token);
      setUserInfo(decodedInfo);

      navigate('/');
    }
  };

  return (
    <motion.div
      initial={{ opacity: 0, x: -100 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: 100 }}
      transition={{ duration: 0.5 }}
      className="max-w-md mx-auto p-6 bg-white rounded-lg mt-10"
    >
      <h2 className="text-2xl font-bold text-center mb-6">Login</h2>

      <div className="flex justify-center mb-6">
        <button
          className={`px-6 py-2 rounded-l-full border transition-colors ${userType === 'creator' ? 'bg-blue-500 text-white' : 'bg-white text-gray-600 hover:bg-blue-100'}`}
          onClick={() => handleToggleChange('creator')}
        >
          Creator
        </button>
        <button
          className={`px-6 py-2 rounded-r-full border transition-colors ${userType === 'business' ? 'bg-blue-500 text-white' : 'bg-white text-gray-600 hover:bg-blue-100'}`}
          onClick={() => handleToggleChange('business')}
        >
          Local Business
        </button>
      </div>

      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label htmlFor="email" className="block text-sm font-medium text-gray-700">
            Email address
          </label>
          <input
            id="email"
            type="email"
            placeholder="Enter email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div className="mb-4">
          <label htmlFor="password" className="block text-sm font-medium text-gray-700">
            Password
          </label>
          <input
            id="password"
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        {errorMessage && <p className="text-red-600 text-sm text-center mt-2">{errorMessage}</p>}

        <button
          type="submit"
          className="w-full bg-blue-500 text-white py-2 rounded-full hover:bg-blue-600 transition-colors mt-4"
        >
          Login as {userType === 'creator' ? 'Creator' : 'Local Business'}
        </button>
      </form>
    </motion.div>
  );
};

export default Login;
