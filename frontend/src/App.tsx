import React from 'react';
import { Routes, Route, useLocation } from 'react-router-dom';
import { AuthProvider } from './contexts/AuthContext';
import { AnimatePresence } from 'framer-motion';

import NavBar from './components/NavBar';
import Home from './pages/Home';
import Learn from './pages/Learn';
import Business from './pages/Business';
import Creators from './pages/Creators';
import Pricing from './pages/Pricing';
import Login from './pages/Login';
import SignUp from './pages/SignUp';

const App: React.FC = () => {
  const location = useLocation();

  return (
    <AuthProvider>
      <NavBar />
      <div className='mt-24 container mx-auto'>
        <AnimatePresence mode='wait'>
          <Routes location={location} key={location.pathname}>
            <Route path="/" element={<Home />} />
            <Route path="/learn" element={<Learn />} />
            <Route path="/business" element={<Business />} />
            <Route path="/creators" element={<Creators />} />
            <Route path="/pricing" element={<Pricing />} />
            <Route path="/login" element={<Login />} />
            <Route path="/signup" element={<SignUp />} />
          </Routes>
        </AnimatePresence>
      </div>
    </AuthProvider>
  );
};

export default App;
