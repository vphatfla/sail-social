import React from 'react';

import { useAuth } from '../../contexts/AuthContext';

const OnboardBusiness = React.lazy(() => import('./OnboardBusiness'));
const OnboardCreator = React.lazy(() => import('./OnboardCreator'));

const Onboarding: React.FC = () => {
    const {isCreator, isBusiness, logout} = useAuth();
    if (isCreator()){
        return (
            <OnboardCreator/>
        )
    }
    else if (isBusiness()){
        return (
            <OnboardBusiness/>
        )
    }
    else{
        logout()
    }
}

export default Onboarding;