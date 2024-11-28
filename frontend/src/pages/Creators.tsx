import React from 'react';
import {motion} from "framer-motion"

const Creators: React.FC = () => {
  return (
    <motion.div
      initial={{ opacity: 0, x: -100 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: 100 }}
      transition={{ duration: 0.5 }}
    >
      <h1>Creators</h1>
      <p>Join Sail Social creator community and collaborate with the world’s biggest brands.</p>
    </motion.div>
  );
};

export default Creators;
