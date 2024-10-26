import React, { useState } from 'react';
// import { useNavigate } from 'react-router-dom';
import { motion } from 'framer-motion';

import { useApi } from '../../utils/ApiUtils';
// import { decryptToken } from '../../utils/TokenUtils';
import { useAuth } from '../../contexts/AuthContext';


const Onboarding: React.FC = () => {
    const 
    { 
        // setUserInfo, 
        userInfo 
    } = useAuth();
    const [creatorInfo, setCreatorInfo] = useState({
        id: userInfo["userId"],
        email: '',
        phoneNumber: '',
    });
    const [errorMessage, setErrorMessage] = useState<string | null>(null);
    const { callApi } = useApi();
//   const navigate = useNavigate();

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setCreatorInfo({
      ...creatorInfo,
      [name]: value,
    });
  };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        const response = await callApi('creator/onboarding', 'POST', creatorInfo);

        if (response.error) {
        setErrorMessage(response.error || 'Onboarding failed, please try again.');
        } else {
            console.log(response)
        //   const token = response.data.token;
        //   localStorage.setItem('BoosterHubToken', token);

        //   const decodedInfo = decryptToken(token);
        //   console.log(decodedInfo);
        //   setUserInfo(decodedInfo);

        //   navigate('/');
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
            <h2 className="text-2xl font-bold text-center mb-6">Welcome to Onboarding</h2>

            <form onSubmit={handleSubmit} className="onboarding-form">
            {errorMessage && <p className="text-red-600 text-sm text-center mt-2">{errorMessage}</p>}
            <div className="mb-4">
                <label htmlFor="email" className="block text-sm font-medium text-gray-700">
                Email
                </label>
                <input
                type="email"
                id="email"
                name="email"
                value={creatorInfo.email}
                onChange={handleChange}
                required
                className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
                />
            </div>
            <div className="mb-4">
                <label htmlFor="phoneNumber" className="block text-sm font-medium text-gray-700">
                Phone Number
                </label>
                <input
                type="text"
                id="phoneNumber"
                name="phoneNumber"
                value={creatorInfo.phoneNumber}
                onChange={handleChange}
                required
                className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
                />
            </div>
            <button
                type="submit"
                className="w-full bg-blue-500 text-white py-2 rounded-full hover:bg-blue-600 transition-colors mt-4"
            >
                Submit
            </button>
            </form>
        </motion.div>
    );
};

export default Onboarding;