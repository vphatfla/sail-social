import React, { Suspense } from 'react';
import { Routes, Route, useLocation } from 'react-router-dom';
import { AnimatePresence } from 'framer-motion';
import { Location } from 'history';

// import { useAuth } from './contexts/AuthContext';
import NavBar from './components/NavBar';


const Home = React.lazy(() => import('./pages/Home'));
const Learn = React.lazy(() => import('./pages/Learn'));
const Business = React.lazy(() => import('./pages/business/Home'));
const Creator = React.lazy(() => import('./pages/creator/Home'));
const Pricing = React.lazy(() => import('./pages/Pricing'));

// for creator pages
const CreatorOnboarding = React.lazy(() => import('./pages/creator/Onboarding.tsx'));
const CreatorFeed = React.lazy(() => import('./pages/creator/Feed.tsx'));

// for business pages
const BusinessOnboarding = React.lazy(() => import('./pages/business/Onboarding.tsx'));


const MainPage: React.FC<{ location: Location }> = ({ location }) => {
  return (
    <Routes location={location} key={location.pathname}>
      <Route path="/" element={<Home />} />
      <Route path="/learn" element={<Learn />} />

      <Route path="/business" element={<Business />} />
      <Route path="/business/onboarding" element={<BusinessOnboarding />} />

      <Route path="/creator" element={<Creator />} />
      <Route path="/creator/onboarding" element={<CreatorOnboarding />} />
      <Route path="/creator/feed" element={<CreatorFeed />} />

      <Route path="/pricing" element={<Pricing />} />
    </Routes>
  )
}

const App: React.FC = () => {
  const location = useLocation();
  // const { isOnboarded, isAuthenticated } = useAuth();

  return (
    <div>
      <NavBar />
      <div className='mt-24 container mx-auto'>
        <AnimatePresence mode='wait'>
          <Suspense
          // fallback={<div>Loading...</div>}
          >
            <MainPage location={location} />
          </Suspense>
        </AnimatePresence>
      </div>
    </div>
  );
};

export default App;
