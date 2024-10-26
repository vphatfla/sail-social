import React, { useState } from 'react';
import { motion } from 'framer-motion';
import { useAuth } from '../contexts/AuthContext';
import { decryptToken } from '../utils/TokenUtils';
import { useNavigate } from 'react-router-dom';
import { useApi } from '../utils/ApiUtils';

const SignUp: React.FC = () => {
  const [userType, setUserType] = useState<string>('creator');
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [successMessage, setSuccessMessage] = useState<string | null>(null);
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
    setSuccessMessage(null);

    const formData = new FormData(event.target as HTMLFormElement);
    const email = formData.get('email') as string;
    const password = formData.get('password') as string;
    const confirmPassword = formData.get('confirmPassword') as string;
    const phoneNumber = formData.get('phoneNumber') as string;
    const businessName = formData.get('businessName') as string | null;

    if (password !== confirmPassword) {
      setErrorMessage('Passwords do not match!');
      return;
    }

    const data = {
      email,
      password,
      phoneNumber,
      ...(userType === 'business' && { businessName }),
    };

    const endpoint = userType === 'creator' ? '/creator/sign-up' : '/business/sign-up';

    const response = await callApi(endpoint, 'POST', data);

    if (response.error) {
      setErrorMessage(response.error || 'Sign up failed, please try again.');
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
      <h2 className="text-2xl font-bold text-center mb-6">Sign Up</h2>

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

      <form onSubmit={handleSubmit} className="signup-form">
        <div className="mb-4">
          <label htmlFor="email" className="block text-sm font-medium text-gray-700">
            Email address
          </label>
          <input
            id="email"
            type="email"
            name="email"
            placeholder="Enter email"
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
            name="password"
            placeholder="Password"
            required
            className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div className="mb-4">
          <label htmlFor="confirmPassword" className="block text-sm font-medium text-gray-700">
            Confirm Password
          </label>
          <input
            id="confirmPassword"
            type="password"
            name="confirmPassword"
            placeholder="Confirm Password"
            required
            className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div className="mb-4">
          <label htmlFor="phoneNumber" className="block text-sm font-medium text-gray-700">
            Phone Number
          </label>
          <input
            id="phoneNumber"
            type="tel"
            required={true}
            name="phoneNumber"
            placeholder="Phone Number"
            className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        {errorMessage && <p className="text-red-600 text-sm text-center mt-2">{errorMessage}</p>}
        {successMessage && <p className="text-green-600 text-sm text-center mt-2">{successMessage}</p>}

        <button
          type="submit"
          className="w-full bg-blue-500 text-white py-2 rounded-full hover:bg-blue-600 transition-colors mt-4"
        >
          Sign Up as {userType === 'creator' ? 'Creator' : 'Local Business'}
        </button>
      </form>
    </motion.div>
  );
};

export default SignUp;
