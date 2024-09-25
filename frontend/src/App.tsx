import React from 'react';
import { Routes, Route } from 'react-router-dom';
import NavBar from './components/NavBar';
import { AuthProvider } from './contexts/AuthContext';

import Home from './pages/Home';
import Learn from './pages/Learn';
import Business from './pages/Business';
import Creators from './pages/Creators';
import Pricing from './pages/Pricing';
import Login from './pages/Login';
import SignUp from './pages/SignUp';

const App: React.FC = () => {
  return (
    <AuthProvider>
      <div>
        <NavBar />
        <div className='mt-24 container mx-auto'>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/learn" element={<Learn />} />
            <Route path="/business" element={<Business />} />
            <Route path="/creators" element={<Creators />} />
            <Route path="/pricing" element={<Pricing />} />
            <Route path="/login" element={<Login />} />
            <Route path="/signup" element={<SignUp />} />
          </Routes>
        </div>
      </div>
    </AuthProvider>
  );
};

export default App;
