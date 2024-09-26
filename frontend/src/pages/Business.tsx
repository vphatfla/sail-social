import React from 'react';
import {motion} from "framer-motion"

const Business: React.FC = () => {
  return (
    <motion.div
      initial={{ opacity: 0, x: -100 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: 100 }}
      transition={{ duration: 0.5 }}
    >
      <h1>Local Businesses</h1>
      <p>Learn more about the businesses who trust us!</p>
    </motion.div>
  );
};

export default Business;