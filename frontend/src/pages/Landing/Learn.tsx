import React from 'react';
import { motion } from 'framer-motion';

const Learn: React.FC = () => {
  return (
    <motion.div
      initial={{ opacity: 0, x: -100 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: 100 }}
      transition={{ duration: 0.5 }}
    >
      <h1>Learn More</h1>
      <p>
        Access resources, tutorials, and expert guidance to make the most out of Booster Hub's platform.
      </p>
    </motion.div>
  );
};

export default Learn;
