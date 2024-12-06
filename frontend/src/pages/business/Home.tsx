import React, { useState } from 'react';
import { motion } from 'framer-motion';
import SignUp from '../../components/business/SignUp';
import Login from '../../components/business/LogIn';

const BusinessHome: React.FC = () => {
    const [isSignUp, setIsSignUp] = useState(true); // To toggle between SignUp and Login

    return (
        <motion.div
            initial={{ opacity: 0, x: -100 }}
            animate={{ opacity: 1, x: 0 }}
            exit={{ opacity: 0, x: 100 }}
            transition={{ duration: 0.5 }}
            className="flex h-screen"
        >
            <div className="flex-1 flex flex-col justify-center items-left p-8">
                <h1 className="text-5xl font-bold mb-4">Join as a Business</h1>
                <p className="text-lg text-gray-700">
                    Boost your business's sales and spread your service by connecting with the top content creators!
                </p>
            </div>

            <div className="flex-1 flex flex-col justify-center items-left bg-white p-8 ml-10">
                <div className="w-full max-w-md">
                    {isSignUp ? <SignUp /> : <Login />}
                    <p>
                        <span style={{ color: 'blue', cursor: 'pointer' }}
                            onClick={() => setIsSignUp(!isSignUp)}
                        >
                            {isSignUp ? "Already a user? Log In!" : "New here? Sign Up"}
                        </span>
                    </p>
                </div>
            </div>
        </motion.div>
    );
};

export default BusinessHome;
