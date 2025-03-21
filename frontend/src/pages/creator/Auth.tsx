import React, { useState } from 'react';
import { motion } from 'framer-motion';
import SignUp from '../../components/creator/SignUp'; // Assuming SignUp is a child component
import Login from '../../components/creator/LogIn'; // Assuming Login is a child component

const CreatorAuth: React.FC = () => {
    const [isSignUp, setIsSignUp] = useState(true); // To toggle between SignUp and Login

    return (
        <motion.div
            initial={{ opacity: 0, x: -100 }}
            animate={{ opacity: 1, x: 0 }}
            exit={{ opacity: 0, x: 100 }}
            transition={{ duration: 0.5 }}
            className="flex h-screen"
        >
            {/* Left Section */}
            <div className="flex-1 flex flex-col justify-center items-left p-8">
                <h1 className="text-5xl font-bold mb-4">Join as a Creator</h1>
                <p className="text-lg text-gray-700">
                    Discover new opportunities, grow your brand, and collaborate with businesses on Sail Social.
                    Join today and unlock your potential!
                </p>
            </div>

            {/* Right Section */}
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

export default CreatorAuth;
