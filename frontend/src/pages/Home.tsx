import React from 'react';
import { motion } from 'framer-motion';
import { useNavigate } from 'react-router-dom';

const Home: React.FC = () => {
  const navigate = useNavigate();
  return (
    <motion.div
      initial={{ opacity: 0, x: -100 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: 100 }}
      transition={{ duration: 0.5 }}
      className="flex h-screen"
    >
      {/* Left Section */}
      <div className="flex-1 flex items-center justify-center">
        <h1 className="text-5xl font-bold">Booster Hub</h1>
      </div>

      {/* Right Section */}
      <div className="flex-1 flex flex-row items-center justify-center bg-white space-x-10">
        <p className="text-xl">I am a</p>
        <div className="space-x-4">
          <button className="btn btn-outline-primary w-48" onClick={() => navigate('/creator')}>Creator</button>
          <button className="btn btn-outline-primary w-48" onClick={() => navigate('/business')}>Business</button>
        </div>
      </div>
    </motion.div>
  );
};

export default Home;
