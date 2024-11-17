import React, { Suspense } from 'react';
import { Routes, Route, useLocation } from 'react-router-dom';
import { AnimatePresence } from 'framer-motion';
import { Location } from 'history';

import { useAuth } from './contexts/AuthContext';
import NavBar from './components/NavBar';
import Onboarding from './components/onboarding/Onboarding';

const Home = React.lazy(() => import('./pages/Home'));
const Learn = React.lazy(() => import('./pages/Learn'));
const Business = React.lazy(() => import('./pages/Business'));
const Creator = React.lazy(() => import('./pages/creator/Home'));
const Pricing = React.lazy(() => import('./pages/Pricing'));
const Login = React.lazy(() => import('./pages/Login'));
const SignUp = React.lazy(() => import('./pages/SignUp'));

const MainPage: React.FC<{ location: Location }> = ({ location }) => {
  return (
    <Routes location={location} key={location.pathname}>
      <Route path="/" element={<Home />} />
      <Route path="/learn" element={<Learn />} />
      <Route path="/business" element={<Business />} />
      <Route path="/creator" element={<Creator />} />
      <Route path="/pricing" element={<Pricing />} />
      <Route path="/login" element={<Login />} />
      <Route path="/signup" element={<SignUp />} />
    </Routes>
  )
}

const App: React.FC = () => {
  const location = useLocation();
  const { isOnboarded, isAuthenticated } = useAuth();

  return (
    <div>
      <NavBar />
      <div className='mt-24 container mx-auto'>
        <AnimatePresence mode='wait'>
          <Suspense
          // fallback={<div>Loading...</div>}
          >
            {
              isAuthenticated() ?
                (isOnboarded() ? <MainPage location={location} /> : <Onboarding />)
                : <MainPage location={location} />
            }
          </Suspense>
        </AnimatePresence>
      </div>
    </div>
  );
};

export default App;
