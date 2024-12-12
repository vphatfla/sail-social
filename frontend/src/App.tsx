import React, { Suspense } from 'react';
import { Routes, Route, useLocation } from 'react-router-dom';
import { AnimatePresence } from 'framer-motion';
import { Location } from 'history';

// import { useAuth } from './contexts/AuthContext';
import GeneralNavBar from './components/navbars/GeneralNavBar.tsx';
import CreatorNavBar from './components/creator/CreatorNavBar.tsx';
import BusinessNavBar from './components/business/BusinessNavBar.tsx';
import JobFeedComponent from './components/creator/jobfeed/JobFeed.tsx';
import AppliedJobsComponent from './components/creator/AppliedJob.tsx';
import NotificationComponent from './components/creator/Notifcation.tsx';


const Home = React.lazy(() => import('./pages/Home'));
const Learn = React.lazy(() => import('./pages/Learn'));
const BusinessAuth = React.lazy(() => import('./pages/business/Auth.tsx'));
const CreatorAuth = React.lazy(() => import('./pages/creator/Auth.tsx'));
const Pricing = React.lazy(() => import('./pages/Pricing'));

// for creator pages
const CreatorOnboarding = React.lazy(() => import('./pages/creator/Onboarding.tsx'));
const CreatorHome = React.lazy(() => import('./pages/creator/Home.tsx'));

// for business pages
const BusinessOnboarding = React.lazy(() => import('./pages/business/Onboarding.tsx'));
const BusinessHome = React.lazy(() => import('./pages/business/Home.tsx'));

const MainPage: React.FC<{ location: Location }> = ({ location }) => {
  return (
    <Routes location={location} key={location.pathname}>
      <Route path="/" element={<Home />} />
      <Route path="/learn" element={<Learn />} />

      <Route path="/business/auth" element={<BusinessAuth />} />
      <Route path="/business/onboarding" element={<BusinessOnboarding />} />
      <Route path="/business" element={<BusinessHome />} />

      <Route path="/creator/auth" element={<CreatorAuth />} />
      <Route path="/creator/onboarding" element={<CreatorOnboarding />} />
      <Route path="/creator" element={<CreatorHome />}>
        <Route index element={<JobFeedComponent />} />
        <Route path="job-feed" element={<JobFeedComponent />} />
        <Route path="applied-jobs" element={<AppliedJobsComponent />} />
        <Route path="notifications" element={<NotificationComponent />} />
      </Route>

      <Route path="/pricing" element={<Pricing />} />
    </Routes>
  )
}

const App: React.FC = () => {
  const location = useLocation();

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  let dynamicNavBar: any

  if (location.pathname.includes("/creator") && !location.pathname.includes("auth")) {
    dynamicNavBar = <CreatorNavBar />
  } else if (location.pathname.includes("/business") && !location.pathname.includes("auth")) {
    dynamicNavBar = <BusinessNavBar />
  } else dynamicNavBar = <GeneralNavBar />

  return (
    <div>
      {dynamicNavBar}
      <div className='mt-24 container mx-auto'>
        <AnimatePresence mode='wait'>
          <Suspense>
            <MainPage location={location} />
          </Suspense>
        </AnimatePresence>
      </div>
    </div>
  );
};

export default App;
