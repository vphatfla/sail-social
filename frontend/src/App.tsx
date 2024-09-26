import React, { Suspense } from 'react';
import { Routes, Route, useLocation } from 'react-router-dom';
import { AuthProvider } from './contexts/AuthContext';
import { AnimatePresence } from 'framer-motion';

import NavBar from './components/NavBar';

const Home = React.lazy(() => import('./pages/Home'));
const Learn = React.lazy(() => import('./pages/Learn'));
const Business = React.lazy(() => import('./pages/Business'));
const Creators = React.lazy(() => import('./pages/Creators'));
const Pricing = React.lazy(() => import('./pages/Pricing'));
const Login = React.lazy(() => import('./pages/Login'));
const SignUp = React.lazy(() => import('./pages/SignUp'));

const App: React.FC = () => {
  const location = useLocation();

  return (
    <AuthProvider>
      <NavBar />
      <div className='mt-24 container mx-auto'>
        <AnimatePresence mode='wait'>
          <Suspense 
          // fallback={<div>Loading...</div>}
          >
            <Routes location={location} key={location.pathname}>
              <Route path="/" element={<Home />} />
              <Route path="/learn" element={<Learn />} />
              <Route path="/business" element={<Business />} />
              <Route path="/creators" element={<Creators />} />
              <Route path="/pricing" element={<Pricing />} />
              <Route path="/login" element={<Login />} />
              <Route path="/signup" element={<SignUp />} />
            </Routes>
          </Suspense>
        </AnimatePresence>
      </div>
    </AuthProvider>
  );
};

export default App;
